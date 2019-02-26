package configuration

type CONFIGURATION interface {
	AppEndpointIp() *string
	SetAppEndpointIp(appEndpointIp *string)
	AppEndpointPort() *string
	SetAppEndpointPort(appEndpointPort *string)
	AppAuthToken() string
	SetAppAuthToken(AppAuthToken string)
}

/*
 * Factory for the CONFIGURATION interface returning CONFIGURATION_IMPL
 */
func NewCONFIGURATION(AppAuthToken string, AppEndpointIp string, AppEndpointPort string) CONFIGURATION {
	configuration := new(CONFIGURATION_IMPL)
	configuration.SetAppEndpointIp(&AppEndpointIp)
	configuration.SetAppEndpointPort(&AppEndpointPort)
	configuration.SetAppAuthToken(AppAuthToken)
	return configuration
}
