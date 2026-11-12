package constants

const (
	PORT = "PORT"

	// PostgreSQL (primary database since the migration off Mongo)
	POSTGRES_HOST           = "POSTGRES_HOST"
	POSTGRES_PORT           = "POSTGRES_PORT"
	POSTGRES_USER           = "POSTGRES_USER"
	POSTGRES_PASSWORD       = "POSTGRES_PASSWORD"
	POSTGRES_DB             = "POSTGRES_DB"
	POSTGRES_SSLMODE        = "POSTGRES_SSLMODE"
	POSTGRES_TIMEZONE       = "POSTGRES_TIMEZONE"
	POSTGRES_MAX_CONNS      = "POSTGRES_MAX_CONNS"
	POSTGRES_MAX_IDLE_CONNS = "POSTGRES_MAX_IDLE_CONNS"

	// MongoDB (legacy, config/mongo is kept for reference/rollback)
	MONGO_HOST          = "MONGO_HOST"
	MONGO_PORT          = "MONGO_PORT"
	MONGO_TYPE          = "MONGO_TYPE"
	MONGO_OPTION        = "MONGO_OPTION"
	MONGO_USERNAME      = "MONGO_USERNAME"
	MONGO_PASSWORD      = "MONGO_PASSWORD"
	MONGO_DATABASE_NAME = "MONGO_DATABASE_NAME"

	CLOUD_NAME       = "CLOUD_NAME"
	CLOUD_API_KEY    = "CLOUD_API_KEY"
	CLOUD_API_SECRET = "CLOUD_API_SECRET"

	REDIS_URL      = "REDIS_URL"
	REDIS_USERNAME = "REDIS_USERNAME"
	REDIS_PASSWORD = "REDIS_PASSWORD"
	REDIS_HOST     = "REDIS_HOST"
	REDIS_PORT     = "REDIS_PORT"
	REDIS_DB       = "REDIS_DB"

	SECRET = "SECRET"

	// SESSION_EXPIRATION is read by both the session store and the JWT
	// middleware; it accepts a duration ("72h") or seconds ("259200").
	SESSION_EXPIRATION = "SESSION_EXPIRATION"

	STATUS = "status"
	DATA   = "data"
	ERROR  = "error"

	SUCCESS = "Success"
	FAILED  = "Failed"

	FOLDER = "folder"

	ERROR_MESSAGE = "error_message"

	LOCALE_ENGLISH   = "en"
	LOCALE_INDONESIA = "id"
)
