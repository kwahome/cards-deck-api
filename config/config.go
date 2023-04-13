package config

// HTTP represents configuration properties required for starting http server.
type HTTP struct {
	Host                      string
	Port                      string
	ReadTimeoutSeconds        int
	WriteTimeoutSeconds       int
	DefaultHttpRequestTimeout int
	Debug                     bool
}

// Config represents all the application configurations grouped as per their category.
type Config struct {
	HTTP      HTTP
	AuthToken string
}
