package httpErrors

import (
	"github.com/santduv/gyma-api/internal/modules/shared/app/constants"
	"github.com/santduv/gyma-api/internal/modules/shared/app/types"
)

func NewInternalServerError(message string, details *types.JsonMap) *HttpError {
	return NewHttpError(&types.HttpErrorArgs{
		StatusCode: constants.HTTP_STATUS_INTERNAL_SERVER_ERROR,
		Message:    message,
		Details:    details,
	})
}
