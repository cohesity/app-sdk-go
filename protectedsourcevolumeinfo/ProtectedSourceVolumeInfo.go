// Copyright 2019 Cohesity Inc.
package protectedsourcevolumeinfo

import "github.com/cohesity/app-sdk-go/configuration"
import "github.com/cohesity/app-sdk-go/models"

/*
 * Interface for the PROTECTEDSOURCEVOLUMEINFO_IMPL
 */
type PROTECTEDSOURCEVOLUMEINFO interface {
	GetProtectedSourceVolumeInfo(*models.GetSourceVolumeInfoParams) (*models.ProtectedSourceVolumeInfo, error)
}

/*
 * Factory for the PROTECTEDSOURCEVOLUMEINFO interaface returning PROTECTEDSOURCEVOLUMEINFO_IMPL
 */
func NewPROTECTEDSOURCEVOLUMEINFO(config configuration.CONFIGURATION) *PROTECTEDSOURCEVOLUMEINFO_IMPL {
	client := new(PROTECTEDSOURCEVOLUMEINFO_IMPL)
	client.config = config
	return client
}
