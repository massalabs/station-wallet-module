package handler

import (
	"github.com/go-openapi/loads"
	"github.com/massalabs/thyra-plugin-massa-wallet/api/server/restapi"
	"github.com/massalabs/thyra-plugin-massa-wallet/api/server/restapi/operations"
	"github.com/massalabs/thyra-plugin-massa-wallet/internal/handler/html"
	"github.com/massalabs/thyra-plugin-massa-wallet/internal/handler/wallet"
	"github.com/massalabs/thyra-plugin-massa-wallet/pkg/password"
	"github.com/massalabs/thyra-plugin-massa-wallet/pkg/privateKey"
)

// InitializeAPI initializes the API handlers
func InitializeAPI(passwordPrompter password.PasswordAsker, privateKeyPrompter privateKey.PrivateKeyAsker) (*operations.MassaWalletAPI, error) {
	// Load the Swagger specification
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		return nil, err
	}

	// Create a new MassaWalletAPI instance
	api := operations.NewMassaWalletAPI(swaggerSpec)

	// Set web endpoints
	html.AppendEndpoints(api)

	// Set wallet API endpoints
	wallet.AppendEndpoints(api, privateKeyPrompter, passwordPrompter)

	return api, nil
}
