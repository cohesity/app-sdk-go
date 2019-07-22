// Copyright 2019 Cohesity Inc.
package configuration

import(
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
    /** Replace the value of app_endpoint_ip with SetAppEndpointIp function */
    app_endpoint_ip *string
    /** Replace the value of app_endpoint_port with SetAppEndpointPort function */
    app_endpoint_port *string
    /** OAuth 2.0 Access Token */
    /** Replace the value of oauthaccesstoken with SetOAuthAccessToken function */
    oauthaccesstoken string
    /**Configures verification of SSL certificates**/
    skip_ssl bool
}

/*
 * Getter function returning app_endpoint_ip
 */
func (me *CONFIGURATION_IMPL) AppEndpointIp() *string{
    return me.app_endpoint_ip
}

/*
 * Setter function setting up app_endpoint_ip
 */
func (me *CONFIGURATION_IMPL) SetAppEndpointIp(appEndpointIp *string) {
    me.app_endpoint_ip = appEndpointIp
}

/*
 * Getter function returning app_endpoint_port
 */
func (me *CONFIGURATION_IMPL) AppEndpointPort() *string{
    return me.app_endpoint_port
}

/*
 * Setter function setting up app_endpoint_port
 */
func (me *CONFIGURATION_IMPL) SetAppEndpointPort(appEndpointPort *string) {
    me.app_endpoint_port = appEndpointPort
}

/*
 * Getter function returning oauthaccesstoken
 */
func (me *CONFIGURATION_IMPL) OAuthAccessToken() string{
    return me.oauthaccesstoken
}

/*
 * Setter function setting up oauthaccesstoken
 */
func (me *CONFIGURATION_IMPL) SetOAuthAccessToken(oAuthAccessToken string) {
    me.oauthaccesstoken = oAuthAccessToken
}

/*
 * Getter function returning skip_ssl
 */
func (me *CONFIGURATION_IMPL) SkipSSL() bool{
    return me.skip_ssl
}

/*
 * Setter function setting up skip_ssl
 */
func (me *CONFIGURATION_IMPL) SetSkipSSL(skipSSL bool) {
    me.skip_ssl = skipSSL
}
// Setting up Default Environment
var Environment = PRODUCTION

// A map of environments and their corresponding servers/baseurls
var EnvironmentsMap = map[Environments](map[Servers]string){

    PRODUCTION : map[Servers]string{
        ENUM_DEFAULT:"http://{app_endpoint_ip}:{app_endpoint_port}/athenaservices/api/v1/public",
    },
}

// Make and return the map of parameters
func GetBaseURIParameters(config CONFIGURATION) map[string]interface{} {
     kvpMap := map[string]interface{}{
         "app_endpoint_ip": config.AppEndpointIp(),
         "app_endpoint_port": config.AppEndpointPort(),
    }
    return kvpMap;
}

// Gets the URL for a particular alias in the current environment and appends it with template parameters
// return the baseurl
func GetBaseURI(server Servers, config CONFIGURATION) string {
    url := EnvironmentsMap[Environment][server]
    appendedURL, _ := apihelper.AppendUrlWithTemplateParameters(url, GetBaseURIParameters(config))
    return appendedURL

}
