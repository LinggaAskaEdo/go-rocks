package entity

import "github.com/palantir/stacktrace"

var (
	ErrCode      = stacktrace.GetCode
	New          = stacktrace.NewError
	NewWithCode  = stacktrace.NewErrorWithCode
	RootCause    = stacktrace.RootCause
	Wrap         = stacktrace.Propagate
	WrapWithCode = stacktrace.PropagateWithCode
)

type (
	Code         = stacktrace.ErrorCode
	ErrorMessage map[Code]Message

	Message struct {
		StatusCode    int    `json:"status_code"`
		EN            string `json:"en"`
		ID            string `json:"id"`
		HasAnnotation bool
	}
)
