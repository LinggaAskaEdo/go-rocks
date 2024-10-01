package common

import "github.com/linggaaskaedo/go-rocks/stdlib/errors/entity"

var ErrorMessages = entity.ErrorMessage{
	CodeHTTPBadRequest:          entity.ErrMsgBadRequest,
	CodeHTTPNotFound:            entity.ErrMsgNotFound,
	CodeHTTPUnauthorized:        entity.ErrMsgUnauthorized,
	CodeHTTPInternalServerError: entity.ErrMsgISE,
	CodeHTTPUnmarshal:           entity.ErrMsgBadRequest,
	CodeHTTPMarshal:             entity.ErrMsgISE,
	CodeHTTPConflict:            entity.ErrMsgConflict,
	CodeHTTPForbidden:           entity.ErrMsgForbidden,
	CodeHTTPUnprocessableEntity: entity.ErrMsgUnprocessable,
	CodeHTTPTooManyRequest:      entity.ErrMsgTooManyRequest,
	CodeHTTPServiceUnavailable:  entity.ErrMsgServiceUnavailable,
	CodeHTTPParamDecode:         entity.ErrMsgBadRequest,
	CodeHTTPErrorOnReadBody:     entity.ErrMsgISE,

	CodeSQLBuilder:                    entity.ErrMsgISE,
	CodeSQLRead:                       entity.ErrMsgISE,
	CodeSQLRowScan:                    entity.ErrMsgISE,
	CodeSQLCreate:                     entity.ErrMsgISE,
	CodeSQLUpdate:                     entity.ErrMsgISE,
	CodeSQLDelete:                     entity.ErrMsgISE,
	CodeSQLUnlink:                     entity.ErrMsgISE,
	CodeSQLTxBegin:                    entity.ErrMsgISE,
	CodeSQLTxCommit:                   entity.ErrMsgISE,
	CodeSQLPrepareStmt:                entity.ErrMsgISE,
	CodeSQLRecordMustExist:            entity.ErrMsgNotFound,
	CodeSQLCannotRetrieveLastInsertID: entity.ErrMsgISE,
	CodeSQLCannotRetrieveAffectedRows: entity.ErrMsgISE,
	CodeSQLUniqueConstraint:           entity.ErrMsgUniqueConst,
	CodeSQLRecordDoesNotMatch:         entity.ErrMsgBadRequest,
	CodeSQLRecordIsExpired:            entity.ErrMsgBadRequest,
	CodeSQLRecordDoesNotExist:         entity.ErrMsgNotFound,
	CodeSQLForeignKeyMissing:          entity.ErrMsgISE,

	CodeTokenStillValid:        entity.ErrMsgTokenStillValid,
	CodeTokenRefreshStillValid: entity.ErrMsgRefreshStillValid,
}
