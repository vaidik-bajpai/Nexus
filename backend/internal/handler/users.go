package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/vaidik-bajpai/Nexus/backend/internal/helper"
	"github.com/vaidik-bajpai/Nexus/backend/internal/mailer"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

func (h *handler) handleUserRegistration(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("handling user registration")

	var usr types.CreateCredentialsUser
	if err := helper.ReadJSON(r, &usr); err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "failed to read request body",
			Data:    nil,
		})
		return
	}

	h.logger.Debug("request body", zap.Any("request", usr))

	// validate
	if err := h.validator.Struct(usr); err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "failed to validate request body",
			Data:    nil,
		})
		return
	}

	h.logger.Debug("request body validated")

	hashedPassword, err := helper.HashPassword(usr.Password)
	if err != nil {
		helper.WriteJSON(w, http.StatusInternalServerError, &types.Response{
			Status:  http.StatusInternalServerError,
			Message: "failed to hash password",
			Data:    nil,
		})
		return
	}

	h.logger.Debug("password hashed")

	usr.PasswordHash = hashedPassword

	// create token
	usr.Token = strconv.Itoa(helper.CreateRandomToken())
	usr.TokenTTL = time.Now().Add(24 * time.Hour)
	usr.TokenScope = "email_verification"

	// create user
	err = h.store.CreateCredentialsUser(r.Context(), &usr)
	if err != nil {
		helper.WriteJSON(w, http.StatusInternalServerError, &types.Response{
			Status:  http.StatusInternalServerError,
			Message: "failed to create user",
			Data:    nil,
		})
		return
	}

	if err := mailer.SendEmailVerificationEmail(
		[]string{usr.Email},
		"Email Verification",
		fmt.Sprintf("http://localhost:3000/verify-email?token=%s", usr.Token),
	); err != nil {
		helper.WriteJSON(w, http.StatusInternalServerError, &types.Response{
			Status:  http.StatusInternalServerError,
			Message: "failed to send email verification email",
			Data:    nil,
		})
		return
	}

	helper.WriteJSON(w, http.StatusCreated, &types.Response{
		Status:  http.StatusCreated,
		Message: "user created successfully",
		Data:    nil,
	})
}

func (h *handler) handleUserLogin(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("handling user login")

	var usr types.LoginCredentialsUser
	if err := helper.ReadJSON(r, &usr); err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "failed to read request body",
			Data:    nil,
		})
		return
	}

	h.logger.Debug("request body", zap.Any("request", usr))

	// validate
	if err := h.validator.Struct(usr); err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "failed to validate request body",
			Data:    nil,
		})
		return
	}

	h.logger.Debug("request body validated")

	// check if user exists
	user, err := h.store.GetUserByEmail(r.Context(), usr.Email)
	if err != nil {
		helper.WriteJSON(w, http.StatusInternalServerError, &types.Response{
			Status:  http.StatusInternalServerError,
			Message: "failed to get user by email",
			Data:    nil,
		})
		return
	}

	h.logger.Debug("user", zap.Any("user", user))

	// check if password is correct
	if err := helper.ComparePassword(user.Password, usr.Password); err != nil {
		helper.WriteJSON(w, http.StatusUnauthorized, &types.Response{
			Status:  http.StatusUnauthorized,
			Message: "invalid user credentials",
			Data:    nil,
		})
		return
	}

	h.logger.Debug("password is correct")

	// generate token
	accessToken, refreshToken, err := helper.GenerateAccessAndRefreshTokens(&types.User{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	})
	if err != nil {
		helper.WriteJSON(w, http.StatusInternalServerError, &types.Response{
			Status:  http.StatusInternalServerError,
			Message: "failed to generate token",
			Data:    nil,
		})
		return
	}

	h.logger.Debug("access token", zap.String("accessToken", accessToken))
	h.logger.Debug("refresh token", zap.String("refreshToken", refreshToken))

	err = h.store.UpdateUserRefreshToken(r.Context(), user.ID, refreshToken)
	if err != nil {
		helper.WriteJSON(w, http.StatusInternalServerError, &types.Response{
			Status:  http.StatusInternalServerError,
			Message: "failed to update user refresh token",
			Data:    nil,
		})
		return
	}

	h.logger.Debug("user refresh token updated")

	helper.WriteJSON(w, http.StatusOK, &types.Response{
		Status:  http.StatusOK,
		Message: "user logged in successfully",
		Data: map[string]string{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		},
	})
}

func (h *handler) handleUserOAuthFlow(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("handling user oauth flow")

	provider := r.PathValue("provider")
	if provider != "google" {
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "provider not supported",
			Data:    nil,
		})
		return
	}

	h.logger.Debug("provider", zap.String("provider", provider))

	oAuthCfg := h.oauth2[provider]
	if oAuthCfg == nil {
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "provider not supported",
		})
		return
	}

	state, err := helper.GenerateOAuthState()
	if err != nil {
		helper.WriteJSON(w, http.StatusInternalServerError, &types.Response{
			Status:  http.StatusInternalServerError,
			Message: "failed to generate state token",
		})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "oauth_state",
		Value:    state,
		Path:     "/",
		MaxAge:   600,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	h.logger.Debug("oAuthCfg", zap.Any("oAuthCfg", oAuthCfg))

	authURL := oAuthCfg.AuthCodeURL(state, oauth2.AccessTypeOffline, oauth2.ApprovalForce)

	h.logger.Debug("authURL", zap.String("authURL", authURL))
	h.logger.Info("redirecting to authURL", zap.String("authURL", authURL))

	http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
}

func (h *handler) handleUserOAuthCallback(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("handling user oauth callback")

	provider := r.PathValue("provider")
	if provider != "google" {
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "provider not supported",
		})
		return
	}

	// Verify state token
	stateCookie, err := r.Cookie("oauth_state")
	if err != nil || stateCookie == nil {
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "state cookie not found",
		})
		return
	}

	returnedState := r.URL.Query().Get("state")
	if returnedState == "" || returnedState != stateCookie.Value {
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "invalid state parameter",
		})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "oauth_state",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})

	code := r.URL.Query().Get("code")
	if code == "" {
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "code is required",
		})
		return
	}

	oAuthCfg := h.oauth2[provider]
	if oAuthCfg == nil {
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "provider not supported",
		})
		return
	}

	token, err := oAuthCfg.Exchange(context.Background(), code)
	if err != nil {
		helper.WriteJSON(w, http.StatusInternalServerError, &types.Response{
			Status:  http.StatusInternalServerError,
			Message: "failed to exchange code for token",
		})
		return
	}

	h.logger.Debug("token exchanged successfully", zap.Any("token", token))

	client := h.oauth2[provider].Client(context.Background(), token)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		helper.WriteJSON(w, http.StatusInternalServerError, &types.Response{
			Status:  http.StatusInternalServerError,
			Message: "failed to get user info",
		})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		helper.WriteJSON(w, http.StatusInternalServerError, &types.Response{
			Status:  http.StatusInternalServerError,
			Message: "failed to read user info",
			Data:    nil,
		})
		return
	}

	var userInfo types.GoogleOAuthUserInfo
	if err := json.Unmarshal(body, &userInfo); err != nil {
		helper.WriteJSON(w, http.StatusInternalServerError, &types.Response{
			Status:  http.StatusInternalServerError,
			Message: "failed to unmarshal user info",
			Data:    nil,
		})
		return
	}

	h.logger.Debug("user info", zap.Any("userInfo", userInfo))

	createOAuthUser := &types.CreateOAuthUser{
		Email:             userInfo.Email,
		Provider:          provider,
		ProviderAccountID: userInfo.ID,
	}

	if err := h.store.CreateOAuthUser(r.Context(), createOAuthUser); err != nil {
		h.logger.Error("failed to create oauth user", zap.Error(err))
		helper.WriteJSON(w, http.StatusInternalServerError, &types.Response{
			Status:  http.StatusInternalServerError,
			Message: "failed to create oauth user",
		})
		return
	}

	h.logger.Debug("oauth user created", zap.Any("oauthUser", createOAuthUser))

	accessToken, refreshToken, err := helper.GenerateAccessAndRefreshTokens(&types.User{
		ID:    createOAuthUser.ID,
		Email: createOAuthUser.Email,
	})
	if err != nil {
		helper.WriteJSON(w, http.StatusInternalServerError, &types.Response{
			Status:  http.StatusInternalServerError,
			Message: "failed to generate access and refresh tokens",
		})
		return
	}

	h.logger.Debug("access token", zap.String("accessToken", accessToken))
	h.logger.Debug("refresh token", zap.String("refreshToken", refreshToken))

	err = h.store.UpdateUserRefreshToken(r.Context(), createOAuthUser.ID, refreshToken)
	if err != nil {
		h.logger.Error("failed to update user refresh token", zap.Error(err))
		helper.WriteJSON(w, http.StatusInternalServerError, &types.Response{
			Status:  http.StatusInternalServerError,
			Message: "failed to update user refresh token",
			Data:    nil,
		})
		return
	}

	helper.WriteJSON(w, http.StatusOK, &types.Response{
		Status:  http.StatusOK,
		Message: "user oauth callback successfully",
		Data: map[string]string{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		},
	})
}

func (h *handler) handleUserLogout(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("handling user logout")

	user := r.Context().Value(types.UserCtxKey).(*types.User)
	h.logger.Debug("user", zap.Any("user", user))

	err := h.store.UpdateUserRefreshToken(r.Context(), user.ID, "")
	if err != nil {
		h.logger.Error("failed to update user refresh token", zap.Error(err))
		helper.WriteJSON(w, http.StatusInternalServerError, &types.Response{
			Status:  http.StatusInternalServerError,
			Message: "failed to update user refresh token",
		})
	}

	h.logger.Debug("user refresh token updated")

	helper.WriteJSON(w, http.StatusOK, &types.Response{
		Status:  http.StatusOK,
		Message: "user logged out successfully",
	})
}

func (h *handler) handlePasswordResetFlow(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("handling password reset flow")

	var passwordResetRequest types.PasswordResetRequest
	if err := helper.ReadJSON(r, &passwordResetRequest); err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "failed to read request body",
			Data:    nil,
		})
		return
	}

	if err := h.validator.Struct(passwordResetRequest); err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "failed to validate request body",
			Data:    nil,
		})
		return
	}

	h.logger.Debug("password reset request", zap.Any("passwordResetRequest", passwordResetRequest))

	user, err := h.store.GetUserByEmail(r.Context(), passwordResetRequest.Email)
	if err != nil {
		helper.WriteJSON(w, http.StatusInternalServerError, &types.Response{
			Status:  http.StatusInternalServerError,
			Message: "failed to get user by email",
		})
		return
	}

	if user.Password == "" {
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "invalid user credentials",
			Data:    nil,
		})
		return
	}

	token := &types.Token{
		UserID: user.ID,
		Token:  strconv.Itoa(helper.CreateRandomToken()),
		TTL:    time.Now().Add(1 * time.Hour),
		Scope:  "reset_password",
	}

	err = h.store.CreateToken(r.Context(), token)
	if err != nil {
		h.logger.Error("failed to create password reset token", zap.Error(err))
		helper.WriteJSON(w, http.StatusInternalServerError, &types.Response{
			Status:  http.StatusInternalServerError,
			Message: "failed to create password reset token",
			Data:    nil,
		})
		return
	}

	if err := mailer.SendPasswordResetEmail(
		[]string{passwordResetRequest.Email},
		"Password Reset",
		fmt.Sprintf("http://localhost:3000/reset-password?token=%s", token.Token),
	); err != nil {
		h.logger.Error("failed to send password reset email", zap.Error(err))
		helper.WriteJSON(w, http.StatusInternalServerError, &types.Response{
			Status:  http.StatusInternalServerError,
			Message: "failed to send password reset email",
			Data:    nil,
		})
		return
	}

	h.logger.Debug("password reset email sent", zap.Any("passwordResetEmail", passwordResetRequest.Email))

	helper.WriteJSON(w, http.StatusOK, &types.Response{
		Status:  http.StatusOK,
		Message: "password reset flow successfully",
	})
}

func (h *handler) handlePasswordReset(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("handling password reset")

	token := r.URL.Query().Get("token")
	if token == "" {
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "token is required",
		})
		return
	}

	usr, err := h.store.GetUserByToken(r.Context(), token)
	if err != nil {
		helper.WriteJSON(w, http.StatusInternalServerError, &types.Response{
			Status:  http.StatusInternalServerError,
			Message: "failed to get user by token",
		})
		return
	}

	if usr.Password == "" || usr.Token.Scope != "reset_password" || time.Now().After(usr.Token.TTL) {
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "invalid or expired token",
			Data:    nil,
		})
		return
	}

	var passwordReset types.PasswordReset
	if err := helper.ReadJSON(r, &passwordReset); err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "failed to read request body",
			Data:    nil,
		})
		return
	}

	if err := h.validator.Struct(passwordReset); err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "failed to validate request body",
			Data:    nil,
		})
		return
	}

	h.logger.Debug("user", zap.Any("user", usr.User))

	hashedPassword, err := helper.HashPassword(passwordReset.Password)
	if err != nil {
		helper.WriteJSON(w, http.StatusInternalServerError, &types.Response{
			Status:  http.StatusInternalServerError,
			Message: "failed to hash password",
			Data:    nil,
		})
		return
	}

	err = h.store.UpdateUserPassword(r.Context(), usr.ID, hashedPassword)
	if err != nil {
		helper.WriteJSON(w, http.StatusInternalServerError, &types.Response{
			Status:  http.StatusInternalServerError,
			Message: "failed to update user password",
		})
		return
	}

	h.logger.Debug("user password updated")

	helper.WriteJSON(w, http.StatusOK, &types.Response{
		Status:  http.StatusOK,
		Message: "password reset successfully",
	})
}

func (h *handler) handleRefreshToken(w http.ResponseWriter, r *http.Request) {
	var refreshToken types.RefreshToken
	if err := helper.ReadJSON(r, refreshToken); err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "invalid refresh token credentials",
			Data:    nil,
		})
		return
	}

	if err := h.validator.Struct(refreshToken); err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "invalid refresh token credentials",
			Data:    nil,
		})
		return
	}

	usr, err := helper.VerifyToken(refreshToken.RefreshToken, helper.GetStrEnvOrPanic("REFRESH_TOKEN_SECRET"))
	if err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "invalid refresh token credentials",
			Data:    nil,
		})
		return
	}

	dbUsr, err := h.store.GetUserByEmail(context.Background(), usr.Email)
	if err != nil {
		helper.WriteJSON(w, http.StatusInternalServerError, &types.Response{
			Status:  http.StatusInternalServerError,
			Message: "something went wrong with our server",
			Data:    nil,
		})
		return
	}

	if dbUsr.RefereshToken != refreshToken.RefreshToken {
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "invalid refresh token credentials",
			Data:    nil,
		})
		return
	}

	accessToken, _, err := helper.GenerateAccessAndRefreshTokens(usr)
	if err != nil {
		helper.WriteJSON(w, http.StatusInternalServerError, &types.Response{
			Status:  http.StatusInternalServerError,
			Message: "could not issue new access token",
			Data:    nil,
		})
		return
	}

	helper.WriteJSON(w, http.StatusOK, &types.Response{
		Status:  http.StatusOK,
		Message: "access token refresh successfully",
		Data: map[string]any{
			"access_token": accessToken,
		},
	})
}
