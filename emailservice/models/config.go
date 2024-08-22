package models

type ProviderConfig struct {
	Name   string
	Config map[string]string
}

type ServiceConfig struct {
	Providers []ProviderConfig
}
