package consts

// Declare Types
type TOKEN string
type GO_ENV = string

const (
	// JWT
	ACC_TOKEN TOKEN = "access_token"
	REF_TOKEN TOKEN = "refresh_token"
	// Env
	DEVELOPMENT GO_ENV = "development"
	PRODUCTION  GO_ENV = "production"
)
