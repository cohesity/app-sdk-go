// Copyright 2019 Cohesity Inc.
package mount

import "github.com/cohesity/app-sdk-go/configuration"
import "github.com/cohesity/app-sdk-go/models"

/*
 * Interface for the MOUNT_IMPL
 */
type MOUNT interface {
	DeleteUnmount(*models.DeleteUnmountParams) error

	CreateMount(*models.CreateMountParams) error
}

/*
 * Factory for the MOUNT interaface returning MOUNT_IMPL
 */
func NewMOUNT(config configuration.CONFIGURATION) *MOUNT_IMPL {
	client := new(MOUNT_IMPL)
	client.config = config
	return client
}
