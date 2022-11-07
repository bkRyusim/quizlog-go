package config

var Config = RuntimeConfig{}

type RuntimeConfig struct {
	TistoryConfig TistoryConfig `mapstructure:"tistory"`
}

type TistoryConfig struct {
	ClientId     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
	RedirectUri  string `mapstructure:"redirect_uri"`
}
