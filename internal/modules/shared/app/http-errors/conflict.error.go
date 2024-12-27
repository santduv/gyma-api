package httpErrors

import (
	"github.com/santduv/gyma-api/internal/modules/shared/app/constants"
	"github.com/santduv/gyma-api/internal/modules/shared/app/types"
)

func NewConflictError(message string, details *types.JsonMap) *HttpError {
	return NewHttpError(&types.HttpErrorArgs{
		StatusCode: constants.HTTP_STATUS_CONFLICT,
		Message:    message,
		Details:    details,
	})
}
