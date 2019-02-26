package CohesityAppSdk

import (
	"github.com/cohesity/app-sdk-go/configuration"
	"github.com/cohesity/app-sdk-go/mount"
	"github.com/cohesity/app-sdk-go/settings"
	"github.com/cohesity/app-sdk-go/tokenmanagement"
)

/*
 * Interface for the COHESITYAPPSDK_IMPL
 */
type COHESITYAPPSDK interface {
	Mount() mount.MOUNT
	Settings() settings.SETTINGS
	TokenManagement() tokenmanagement.TOKENMANAGEMENT
	Configuration() configuration.CONFIGURATION
}

/*
 * Factory for the COHESITYAPPSDK interface returning COHESITYAPPSDK_IMPL
 */
func NewAppSdkClient(AppAuthToken string, AppEndpointIp string, AppEndpointPort string) COHESITYAPPSDK {
	cohesityAppSdkClient := new(COHESITYAPPSDK_IMPL)
	cohesityAppSdkClient.config = configuration.NewCONFIGURATION(AppAuthToken, AppEndpointIp, AppEndpointPort)
	return cohesityAppSdkClient
}
