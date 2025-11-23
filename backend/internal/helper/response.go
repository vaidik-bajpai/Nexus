// helper/response.go
package helper

import (
	"net/http"

	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
	"go.uber.org/zap"
)

// SendErrorResponse sends an error response and logs it
func SendErrorResponse(logger *zap.Logger, w http.ResponseWriter, statusCode int, message string, data any, err error) {
	response := &types.Response{
		Status:  int64(statusCode),
		Message: message,
		Data:    data,
	}

	WriteJSON(w, int64(statusCode), response)

	fields := []zap.Field{
		zap.Int("status_code", statusCode),
		zap.String("message", message),
	}

	if err != nil {
		fields = append(fields, zap.Error(err))
	}

	if statusCode >= 500 {
		logger.Error("internal server error", fields...)
	} else if statusCode >= 400 {
		logger.Warn("client error", fields...)
	}
}

// SendSuccessResponse sends a success response and logs it
func SendSuccessResponse(logger *zap.Logger, w http.ResponseWriter, statusCode int, message string, data any) {
	response := &types.Response{
		Status:  int64(statusCode),
		Message: message,
		Data:    data,
	}

	WriteJSON(w, int64(statusCode), response)

	logger.Info("success response",
		zap.Int("status_code", statusCode),
		zap.String("message", message),
	)
}

func BadRequest(logger *zap.Logger, w http.ResponseWriter, message string, data any) {
	SendErrorResponse(logger, w, http.StatusBadRequest, message, data, nil)
}

func UnprocessableEntity(logger *zap.Logger, w http.ResponseWriter, message string, data any) {
	SendErrorResponse(logger, w, http.StatusUnprocessableEntity, message, data, nil)
}

func InternalServerError(logger *zap.Logger, w http.ResponseWriter, data any, err error) {
	SendErrorResponse(logger, w, http.StatusInternalServerError, "something went wrong with our servers", data, err)
}

func Created(logger *zap.Logger, w http.ResponseWriter, message string, data any) {
	SendSuccessResponse(logger, w, http.StatusCreated, message, data)
}

func OK(logger *zap.Logger, w http.ResponseWriter, message string, data any) {
	SendSuccessResponse(logger, w, http.StatusOK, message, data)
}
