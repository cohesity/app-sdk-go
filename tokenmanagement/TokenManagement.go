package tokenmanagement

import "github.com/cohesity/app-sdk-go/models"
import "github.com/cohesity/app-sdk-go/configuration"

/*
 * Interface for the TOKENMANAGEMENT_IMPL
 */
type TOKENMANAGEMENT interface {
    CreateManagementAccessToken () (*models.ManagementAccessToken, error)
}

/*
 * Factory for the TOKENMANAGEMENT interaface returning TOKENMANAGEMENT_IMPL
 */
func NewTOKENMANAGEMENT(config configuration.CONFIGURATION) *TOKENMANAGEMENT_IMPL {
    client := new(TOKENMANAGEMENT_IMPL)
    client.config = config
    return client
}
