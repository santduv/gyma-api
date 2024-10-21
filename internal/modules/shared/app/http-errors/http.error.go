package httpErrors

import (
	"time"

	"github.com/santduv/gyma-api/internal/modules/shared/app/types"
)

type HttpError struct {
	Message    string         `json:"message"`
	StatusCode int            `json:"status"`
	Details    *types.JsonMap `json:"details,omitempty"`
	Date       string         `json:"date"`
}

func (e *HttpError) Error() string {
	return e.Message
}

func NewHttpError(args *types.HttpErrorArgs) *HttpError {
	return &HttpError{
		StatusCode: args.StatusCode,
		Date:       time.Now().Format(time.RFC3339),
		Message:    args.Message,
		Details:    args.Details,
	}
}
