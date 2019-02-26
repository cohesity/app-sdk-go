package mount

import "github.com/cohesity/app-sdk-go/models"
import "github.com/cohesity/app-sdk-go/configuration"

/*
 * Interface for the MOUNT_IMPL
 */
type MOUNT interface {
    CreateMount (*models.MountOptions) (error)

    DeleteUnmount (string) (error)
}

/*
 * Factory for the MOUNT interaface returning MOUNT_IMPL
 */
func NewMOUNT(config configuration.CONFIGURATION) *MOUNT_IMPL {
    client := new(MOUNT_IMPL)
    client.config = config
    return client
}
