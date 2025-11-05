package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/vaidik-bajpai/Nexus/backend/internal/helper"
	"github.com/vaidik-bajpai/Nexus/backend/internal/mailer"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
	"go.uber.org/zap"
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
			Message: "invalid password",
			Data:    nil,
		})
	}

	h.logger.Debug("password is correct")

	// generate token
	accessToken, refreshToken, err := helper.GenerateAccessAndRefreshTokens(user.ID)
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
