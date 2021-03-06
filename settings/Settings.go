// Copyright 2019 Cohesity Inc.
package settings

import "github.com/cohesity/app-sdk-go/configuration"
import "github.com/cohesity/app-sdk-go/models"

/*
 * Interface for the SETTINGS_IMPL
 */
type SETTINGS interface {
    GetAppSettings () (*models.AppSettings, error)
}

/*
 * Factory for the SETTINGS interaface returning SETTINGS_IMPL
 */
func NewSETTINGS(config configuration.CONFIGURATION) *SETTINGS_IMPL {
    client := new(SETTINGS_IMPL)
    client.config = config
    return client
}