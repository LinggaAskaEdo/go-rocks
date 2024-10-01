package preference

const (
	// Redis Type
	REDIS_APPS    string = "APPS"
	REDIS_LIMITER string = "LIMITER"
	REDIS_AUTH    string = "AUTH"

	// Correlation ID
	ContextKeyRequestID     string = "requestID"
	ContextKeyCorrelationID string = "correlation_id"

	// Limiter Error Message
	LimitError   string = "Limit should > 0"
	CommandError string = "The command of first number should > 0"
	FormatError  string = "Please check the format with your input."
	MethodError  string = "Please check the method is one of http method."
	ServerError  string = "StatusInternalServerError, please wait a minute."

	// Database Type
	MYSQL    string = `mysql`
	POSTGRES string = `postgres`

	// UserAgent Header
	ContentType string = `content-type`
	ContentJSON string = `application/json`

	// Lang Header
	LangEN string = `en`
	LangID string = `id`

	// Custom HTTP Header
	AppLang string = `x-app-lang`

	// Cache Control Header
	CacheControl          string = `cache-control`
	CacheNoCache          string = `no-cache`
	CacheNoStore          string = `no-store`
	CacheMustRevalidate   string = `must-revalidate`
	CacheMustDBRevalidate string = `must-db-revalidate`
)
