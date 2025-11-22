package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	mailerMock "github.com/vaidik-bajpai/Nexus/backend/internal/mailer/mock"
	m "github.com/vaidik-bajpai/Nexus/backend/internal/store/mock"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

func createTestHandler(mockStore *m.MockStore, mockMailer *mailerMock.MockMailer) *handler {
	logger, _ := zap.NewDevelopment()
	validator := validator.New()

	return &handler{
		logger:    logger,
		validator: validator,
		store:     mockStore,
		mailer:    mockMailer,
		oauth2:    make(map[string]*oauth2.Config),
	}
}

func TestHandleUserRegistration(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    types.CreateCredentialsUser
		setupMock      func(*m.MockStore, *mailerMock.MockMailer)
		expectedStatus int
		expectedMsg    string
	}{
		{
			name: "successful registration",
			requestBody: types.CreateCredentialsUser{
				Email:    "test@example.com",
				Username: "testuser",
				Password: "password123",
			},
			setupMock: func(ms *m.MockStore, mm *mailerMock.MockMailer) {
				ms.On("CreateCredentialsUser", mock.Anything, mock.AnythingOfType("*types.CreateCredentialsUser")).
					Return(nil)
				mm.On("SendEmailVerificationEmail", []string{"test@example.com"}, "Email Verification", mock.AnythingOfType("string")).
					Return(nil)
			},
			expectedStatus: http.StatusCreated,
			expectedMsg:    "user created successfully",
		},
		{
			name: "invalid email format",
			requestBody: types.CreateCredentialsUser{
				Email:    "invalid-email",
				Username: "testuser",
				Password: "password123",
			},
			setupMock:      func(ms *m.MockStore, mm *mailerMock.MockMailer) {},
			expectedStatus: http.StatusBadRequest,
			expectedMsg:    "failed to validate request body",
		},
		{
			name: "password too short",
			requestBody: types.CreateCredentialsUser{
				Email:    "test@example.com",
				Username: "testuser",
				Password: "short",
			},
			setupMock:      func(ms *m.MockStore, mm *mailerMock.MockMailer) {},
			expectedStatus: http.StatusBadRequest,
			expectedMsg:    "failed to validate request body",
		},
		{
			name: "username too short",
			requestBody: types.CreateCredentialsUser{
				Email:    "test@example.com",
				Username: "ab",
				Password: "password123",
			},
			setupMock:      func(ms *m.MockStore, mm *mailerMock.MockMailer) {},
			expectedStatus: http.StatusBadRequest,
			expectedMsg:    "failed to validate request body",
		},
		{
			name: "store error on user creation",
			requestBody: types.CreateCredentialsUser{
				Email:    "test@example.com",
				Username: "testuser",
				Password: "password123",
			},
			setupMock: func(ms *m.MockStore, mm *mailerMock.MockMailer) {
				ms.On("CreateCredentialsUser", mock.Anything, mock.AnythingOfType("*types.CreateCredentialsUser")).
					Return(errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedMsg:    "failed to create user",
		},
		{
			name:           "empty request body",
			requestBody:    types.CreateCredentialsUser{},
			setupMock:      func(ms *m.MockStore, mm *mailerMock.MockMailer) {},
			expectedStatus: http.StatusBadRequest,
			expectedMsg:    "failed to validate request body",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			mockStore := new(m.MockStore)
			mockMailer := new(mailerMock.MockMailer)
			tt.setupMock(mockStore, mockMailer)
			handler := createTestHandler(mockStore, mockMailer)

			// Create request body
			bodyBytes, err := json.Marshal(tt.requestBody)
			assert.NoError(t, err)

			req := httptest.NewRequest(http.MethodPost, "/api/v1/users/register", bytes.NewReader(bodyBytes))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			// Execute
			handler.handleUserRegistration(rr, req)

			// Assert
			assert.Equal(t, tt.expectedStatus, rr.Code)

			var response types.Response
			err = json.Unmarshal(rr.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, int64(tt.expectedStatus), response.Status)
			assert.Equal(t, tt.expectedMsg, response.Message)

			// Verify mock expectations
			mockStore.AssertExpectations(t)
			mockMailer.AssertExpectations(t)
		})
	}
}
