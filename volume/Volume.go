// Copyright 2019 Cohesity Inc.
package volume

import "github.com/cohesity/app-sdk-go/configuration"
import "github.com/cohesity/app-sdk-go/models"

/*
 * Interface for the VOLUME_IMPL
 */
type VOLUME interface {
	GetVolume(*models.GetVolumeParams) (*models.VolumeInfo, error)

	DeleteVolume(*models.DeleteVolumeParams) error

	CreateVolume(*models.CreateVolumeParams) error
}

/*
 * Factory for the VOLUME interaface returning VOLUME_IMPL
 */
func NewVOLUME(config configuration.CONFIGURATION) *VOLUME_IMPL {
	client := new(VOLUME_IMPL)
	client.config = config
	return client
}
