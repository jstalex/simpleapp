package config

type AppConfig struct {
	ListenAddr string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		ListenAddr: ":8080",
	}
}
