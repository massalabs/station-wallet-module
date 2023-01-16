package wallet

import (
	"errors"
	"sync"

	"fyne.io/fyne/v2"
	"github.com/go-openapi/runtime/middleware"
	"github.com/massalabs/thyra-plugin-massa-wallet/api/server/models"
	"github.com/massalabs/thyra-plugin-massa-wallet/api/server/restapi/operations"
	"github.com/massalabs/thyra-plugin-massa-wallet/pkg/password"
	"github.com/massalabs/thyra-plugin-massa-wallet/pkg/wallet"
)

const fileModeUserRW = 0o600

//nolint:nolintlint,ireturn
func NewImport(walletStorage *sync.Map, app *fyne.App) operations.RestWalletImportHandler {
	return &wImport{walletStorage: walletStorage, app: app}
}

type wImport struct {
	walletStorage *sync.Map
	app           *fyne.App
}

//nolint:nolintlint,ireturn,funlen
func (c *wImport) Handle(operations.RestWalletImportParams) middleware.Responder {
	password, walletName, privateKey, err := password.AskWalletInfo(c.app)
	if err != nil {
		return NewWalletError(err.Error(), err.Error())
	}

	if len(walletName) == 0 {
		return operations.NewRestWalletCreateBadRequest().WithPayload(
			&models.Error{
				Code:    err.Error(),
				Message: "Error: nickname field is mandatory.",
			})
	}

	_, inStore := c.walletStorage.Load(walletName)
	if inStore {
		return NewWalletError(err.Error(), "Error: a wallet with the same nickname already exists.")
	}

	if len(password) == 0 {
		return NewWalletError(err.Error(), "Error: password field is mandatory.")
	}

	newWallet, err := wallet.Imported(walletName, privateKey, password)
	if err != nil {
		if errors.Is(err, err) {
			return NewWalletError(err.Error(), err.Error())
		}

		return NewWalletError(err.Error(), err.Error())
	}

	return CreateNewWallet(newWallet)
}

//nolint:nolintlint,ireturn
func NewWalletError(code string, message string) middleware.Responder {
	return operations.NewRestWalletCreateInternalServerError().WithPayload(
		&models.Error{
			Code:    code,
			Message: message,
		})
}
