// Copyright 2019 Cohesity Inc.
package CohesityAppSdk

import(
	"github.com/cohesity/app-sdk-go/configuration"
	"github.com/cohesity/app-sdk-go/protectedsourcevolumeinfo"
	"github.com/cohesity/app-sdk-go/volume"
	"github.com/cohesity/app-sdk-go/tokenmanagement"
	"github.com/cohesity/app-sdk-go/settings"
	"github.com/cohesity/app-sdk-go/mount"
)


/*
 * Interface for the COHESITYAPPSDK_IMPL
 */
type COHESITYAPPSDK interface {
    ProtectedSourceVolumeInfo()       protectedsourcevolumeinfo.PROTECTEDSOURCEVOLUMEINFO
    Volume()                volume.VOLUME
    TokenManagement()       tokenmanagement.TOKENMANAGEMENT
    Settings()              settings.SETTINGS
    Mount()                 mount.MOUNT
    Configuration()         configuration.CONFIGURATION
}

/*
 * Factory for the COHESITYAPPSDK interface returning COHESITYAPPSDK_IMPL
 */
func NewAppSdkClient(AppAuthToken string, AppEndpointIp string, AppEndpointPort string) COHESITYAPPSDK {
    cohesityAppSdkClient := new(COHESITYAPPSDK_IMPL)
    cohesityAppSdkClient.config = configuration.NewCONFIGURATION(AppEndpointIp, AppEndpointPort)

    cohesityAppSdkClient.config.SetOAuthAccessToken(AppAuthToken)

    return cohesityAppSdkClient
}
