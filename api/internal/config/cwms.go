package config

type CwmsConfig struct {
	CwmsApiUrl string `envconfig:"CWMS_API_URL"`
}
