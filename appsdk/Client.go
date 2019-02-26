package CohesityAppSdk

import (
	"github.com/cohesity/app-sdk-go/configuration"
	"github.com/cohesity/app-sdk-go/mount"
	"github.com/cohesity/app-sdk-go/settings"
	"github.com/cohesity/app-sdk-go/tokenmanagement"
)

/*
 * Client structure as interface implementation
 */
type COHESITYAPPSDK_IMPL struct {
	mount           mount.MOUNT
	settings        settings.SETTINGS
	tokenmanagement tokenmanagement.TOKENMANAGEMENT
	config          configuration.CONFIGURATION
}

/**
 * Access to Configuration
 * @return Returns the Configuration instance
 */
func (me *COHESITYAPPSDK_IMPL) Configuration() configuration.CONFIGURATION {
	return me.config
}

/**
 * Access to Mount controller
 * @return Returns the Mount() instance
 */
func (me *COHESITYAPPSDK_IMPL) Mount() mount.MOUNT {
	if (me.mount) == nil {
		me.mount = mount.NewMOUNT(me.config)
	}
	return me.mount
}

/**
 * Access to Settings controller
 * @return Returns the Settings() instance
 */
func (me *COHESITYAPPSDK_IMPL) Settings() settings.SETTINGS {
	if (me.settings) == nil {
		me.settings = settings.NewSETTINGS(me.config)
	}
	return me.settings
}

/**
 * Access to TokenManagement controller
 * @return Returns the TokenManagement() instance
 */
func (me *COHESITYAPPSDK_IMPL) TokenManagement() tokenmanagement.TOKENMANAGEMENT {
	if (me.tokenmanagement) == nil {
		me.tokenmanagement = tokenmanagement.NewTOKENMANAGEMENT(me.config)
	}
	return me.tokenmanagement
}
