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
 * Client structure as interface implementation
 */
type COHESITYAPPSDK_IMPL struct {
     protectedsourcevolumeinfo protectedsourcevolumeinfo.PROTECTEDSOURCEVOLUMEINFO
     volume volume.VOLUME
     tokenmanagement tokenmanagement.TOKENMANAGEMENT
     settings settings.SETTINGS
     mount mount.MOUNT
     config  configuration.CONFIGURATION
}

/**
     * Access to Configuration
     * @return Returns the Configuration instance
*/
func (me *COHESITYAPPSDK_IMPL) Configuration() configuration.CONFIGURATION {
    return me.config
}
/**
     * Access to ProtectedSourceVolumeInfo controller
     * @return Returns the ProtectedSourceVolumeInfo() instance
*/
func (me *COHESITYAPPSDK_IMPL) ProtectedSourceVolumeInfo() protectedsourcevolumeinfo.PROTECTEDSOURCEVOLUMEINFO {
    if(me.protectedsourcevolumeinfo) == nil {
        me.protectedsourcevolumeinfo = protectedsourcevolumeinfo.NewPROTECTEDSOURCEVOLUMEINFO(me.config)
    }
    return me.protectedsourcevolumeinfo
}
/**
     * Access to Volume controller
     * @return Returns the Volume() instance
*/
func (me *COHESITYAPPSDK_IMPL) Volume() volume.VOLUME {
    if(me.volume) == nil {
        me.volume = volume.NewVOLUME(me.config)
    }
    return me.volume
}
/**
     * Access to TokenManagement controller
     * @return Returns the TokenManagement() instance
*/
func (me *COHESITYAPPSDK_IMPL) TokenManagement() tokenmanagement.TOKENMANAGEMENT {
    if(me.tokenmanagement) == nil {
        me.tokenmanagement = tokenmanagement.NewTOKENMANAGEMENT(me.config)
    }
    return me.tokenmanagement
}
/**
     * Access to Settings controller
     * @return Returns the Settings() instance
*/
func (me *COHESITYAPPSDK_IMPL) Settings() settings.SETTINGS {
    if(me.settings) == nil {
        me.settings = settings.NewSETTINGS(me.config)
    }
    return me.settings
}
/**
     * Access to Mount controller
     * @return Returns the Mount() instance
*/
func (me *COHESITYAPPSDK_IMPL) Mount() mount.MOUNT {
    if(me.mount) == nil {
        me.mount = mount.NewMOUNT(me.config)
    }
    return me.mount
}

