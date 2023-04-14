package config

// Config represents all the application configurations
type Config struct {
	Host                      string `mapstructure:"APP_HOST"`
	Port                      string `mapstructure:"APP_PORT"`
	ReadTimeoutSeconds        int    `mapstructure:"READ_TIMEOUT"`
	WriteTimeoutSeconds       int    `mapstructure:"WRITE_TIMEOUT"`
	DefaultHttpRequestTimeout int    `mapstructure:"DEFAULT_HTTP_REQUEST_TIMEOUT"`
	Debug                     bool   `mapstructure:"DEBUG"`
	AuthToken                 string `mapstructure:"AUTH_TOKEN"`
}
