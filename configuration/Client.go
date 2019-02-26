package configuration

import (
	"github.com/cohesity/app-sdk-go/apihelper"
)

/* Setting up enums for Environments and Servers
 */
type Environments int

type Servers int

// Environment Enums
const (
	PRODUCTION Environments = 1 + iota
)

// Servers Enums
const (
	ENUM_DEFAULT Servers = 1 + iota
)

type CONFIGURATION_IMPL struct {
	/** Replace the value of AppEndpointIp with SetAppEndpointIp function */
	appEndpointIp *string
	/** Replace the value of AppEndpointPort with SetAppEndpointPort function */
	appEndpointPort *string
	/** OAuth 2.0 Access Token */
	/** Replace the value of appAuthToken with SetAppAuthToken function */
	appAuthToken string
}

/*
 * Getter function returning AppEndpointIp
 */
func (me *CONFIGURATION_IMPL) AppEndpointIp() *string {
	return me.appEndpointIp
}

/*
 * Setter function setting up AppEndpointIp
 */
func (me *CONFIGURATION_IMPL) SetAppEndpointIp(appEndpointIp *string) {
	me.appEndpointIp = appEndpointIp
}

/*
 * Getter function returning AppEndpointPort
 */
func (me *CONFIGURATION_IMPL) AppEndpointPort() *string {
	return me.appEndpointPort
}

/*
 * Setter function setting up AppEndpointPort
 */
func (me *CONFIGURATION_IMPL) SetAppEndpointPort(appEndpointPort *string) {
	me.appEndpointPort = appEndpointPort
}

/*
 * Getter function returning appAuthToken
 */
func (me *CONFIGURATION_IMPL) AppAuthToken() string {
	return me.appAuthToken
}

/*
 * Setter function setting up appAuthToken
 */
func (me *CONFIGURATION_IMPL) SetAppAuthToken(AppAuthToken string) {
	me.appAuthToken = AppAuthToken
}

// Setting up Default Environment
var Environment = PRODUCTION

// A map of environments and their corresponding servers/baseurls
var EnvironmentsMap = map[Environments](map[Servers]string){

	PRODUCTION: map[Servers]string{
		ENUM_DEFAULT: "http://{AppEndpointIp}:{AppEndpointPort}/athenaservices/api/v1/public",
	},
}

// Make and return the map of parameters
func GetBaseURIParameters(config CONFIGURATION) map[string]interface{} {
	kvpMap := map[string]interface{}{
		"AppEndpointIp":   config.AppEndpointIp(),
		"AppEndpointPort": config.AppEndpointPort(),
	}
	return kvpMap
}

// Gets the URL for a particular alias in the current environment and appends it with template parameters
// return the baseurl
func GetBaseURI(server Servers, config CONFIGURATION) string {
	url := EnvironmentsMap[Environment][server]
	appendedURL, _ := apihelper.AppendUrlWithTemplateParameters(url, GetBaseURIParameters(config))
	return appendedURL

}
