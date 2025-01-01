package httpErrors

import (
	"github.com/santduv/gyma-api/internal/modules/shared/app/constants"
	"github.com/santduv/gyma-api/internal/modules/shared/app/types"
)

func NewUnauthorizedError(message string, details *types.JsonMap) *HttpError {
	return NewHttpError(&types.HttpErrorArgs{
		StatusCode: constants.HTTP_STATUS_UNAUTHORIZED,
		Message:    message,
		Details:    details,
	})
}
