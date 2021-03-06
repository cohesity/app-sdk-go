// Copyright 2019 Cohesity Inc.
package tokenmanagement

import "github.com/cohesity/app-sdk-go/configuration"
import "github.com/cohesity/app-sdk-go/models"

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