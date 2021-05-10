package signhost

type Config struct {
	testing bool
	auth    string
	appKey  string
}

func NewConfig(t bool, auth string, appKey string) *Config {
	return &Config{
		testing: t,
		auth:    auth,
		appKey: appKey,
	}
}
