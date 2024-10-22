package helpers

import (
	"github.com/santduv/gyma-api/internal/modules/shared/app/constants"
	"github.com/santduv/gyma-api/internal/modules/shared/app/types"
)

func OkResponse(message string, data interface{}) *types.ApiResponse {
	return &types.ApiResponse{
		Message: message,
		Status:  constants.HTTP_STATUS_OK,
		Data:    data,
	}
}

func CreatedResponse(message string, data interface{}) *types.ApiResponse {
	return &types.ApiResponse{
		Message: message,
		Status:  constants.HTTP_STATUS_CREATED,
		Data:    data,
	}
}
