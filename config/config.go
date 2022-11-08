package config

var Config = RuntimeConfig{}

type RuntimeConfig struct {
	AuthConfig    AuthConfig    `mapstructure:"auth"`
	TistoryConfig TistoryConfig `mapstructure:"tistory"`
}

type AuthConfig struct {
	SecretKey string `mapstructure:"secret_key"`
}

type TistoryConfig struct {
	ClientId     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
	RedirectUri  string `mapstructure:"redirect_uri"`
}
