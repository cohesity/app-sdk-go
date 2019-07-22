// Copyright 2019 Cohesity Inc.

package configuration

type CONFIGURATION interface {
	AppEndpointIp() *string
	SetAppEndpointIp(appEndpointIp *string)
	AppEndpointPort() *string
	SetAppEndpointPort(appEndpointPort *string)
	OAuthAccessToken() string
	SetOAuthAccessToken(AppAuthToken string)
	SkipSSL() bool
	SetSkipSSL(skipSSL bool)
}

/*
 * Factory for the CONFIGURATION interface returning CONFIGURATION_IMPL
 */
 func NewCONFIGURATION(AppEndpointIp string, AppEndpointPort string) CONFIGURATION {
	configuration := new(CONFIGURATION_IMPL)
	configuration.SetAppEndpointIp(&AppEndpointIp)
	configuration.SetAppEndpointPort(&AppEndpointPort)
	return configuration
}
