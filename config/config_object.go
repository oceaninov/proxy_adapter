/*
Package config
*/
package config

type ConfigObject struct {
	AppHost     string `mapstructure:"APP_HOST"`
	AppPort     string `mapstructure:"APP_PORT"`
	AppName     string `mapstructure:"APP_NAME"`
	AppLogLevel string `mapstructure:"APP_LOG_LEVEL"`

	AzureADHost         string `mapstructure:"AZURE_AD_HOST" envconfig:"AZURE_AD_HOST"`
	AzureADTenantID     string `mapstructure:"AZURE_AD_TENANT_ID" envconfig:"AZURE_AD_TENANT_ID"`
	AzureADClientID     string `mapstructure:"AZURE_AD_CLIENT_ID" envconfig:"AZURE_AD_CLIENT_ID"`
	AzureADClientSecret string `mapstructure:"AZURE_AD_CLIENT_SECRET" envconfig:"AZURE_AD_CLIENT_SECRET"`
	AzureADRedirectURI  string `mapstructure:"AZURE_AD_REDIRECT_URI" envconfig:"AZURE_AD_REDIRECT_URI"`
	AzureADScope        string `mapstructure:"AZURE_AD_SCOPE" envconfig:"AZURE_AD_SCOPE"`
}
