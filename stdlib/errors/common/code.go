package common

import "github.com/linggaaskaedo/go-rocks/stdlib/errors/entity"

const (
	// Code HTTP Handler
	CodeHTTPBadRequest = entity.Code(iota + 100)
	CodeHTTPNotFound
	CodeHTTPUnauthorized
	CodeHTTPInternalServerError
	CodeHTTPUnmarshal
	CodeHTTPMarshal
	CodeHTTPConflict
	CodeHTTPForbidden
	CodeHTTPUnprocessableEntity
	CodeHTTPTooManyRequest
	CodeHTTPValidatorError
	CodeHTTPServiceUnavailable
	CodeHTTPParamDecode
	CodeHTTPErrorOnReadBody
)

const (
	// Error on SQL
	CodeSQLBuilder = entity.Code(iota + 200)
	CodeSQLRead
	CodeSQLRowScan
	CodeSQLCreate
	CodeSQLUpdate
	CodeSQLDelete
	CodeSQLUnlink
	CodeSQLTxBegin
	CodeSQLTxCommit
	CodeSQLPrepareStmt
	CodeSQLRecordMustExist
	CodeSQLCannotRetrieveLastInsertID
	CodeSQLCannotRetrieveAffectedRows
	CodeSQLUniqueConstraint
	CodeSQLRecordDoesNotMatch
	CodeSQLRecordIsExpired
	CodeSQLRecordDoesNotExist
	CodeSQLForeignKeyMissing
	CodeSQLTxRollback
	CodeRequestIDIsNotMatch
	CodeSQLConflict
	CodeSQLEmptyRow
	CodeSQLTableNotExist
)

const (
	// Error on Token
	CodeTokenStillValid = entity.Code(iota + 300)
	CodeTokenRefreshStillValid
)

const (
	// Error On Cache
	CodeCacheMarshal = entity.Code(iota + 400)
	CodeCacheUnmarshal
	CodeCacheGetSimpleKey
	CodeCacheSetSimpleKey
	CodeCacheDeleteSimpleKey
	CodeCacheGetHashKey
	CodeCacheSetHashKey
	CodeCacheDeleteHashKey
	CodeCacheSetExpiration
	CodeCacheDecode
	CodeCacheLockNotAcquired
	CodeCacheLockFailed
	CodeCacheInvalidCastType
	CodeCacheNotFound
)
