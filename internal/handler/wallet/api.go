package wallet

import (
	"errors"

	"github.com/bluele/gcache"
	"github.com/go-openapi/runtime/middleware"
	"github.com/massalabs/station-massa-wallet/api/server/models"
	"github.com/massalabs/station-massa-wallet/api/server/restapi/operations"
	"github.com/massalabs/station-massa-wallet/pkg/assets"
	"github.com/massalabs/station-massa-wallet/pkg/network"
	"github.com/massalabs/station-massa-wallet/pkg/prompt"
	"github.com/massalabs/station-massa-wallet/pkg/wallet/account"
	"github.com/massalabs/station-massa-wallet/pkg/walletmanager"
	"github.com/massalabs/station/pkg/logger"
)

// AppendEndpoints appends wallet endpoints to the API
// Note: the password prompter is mandatory for sign endpoint
func AppendEndpoints(api *operations.MassaWalletAPI, prompterApp prompt.WalletPrompterInterface, massaClient network.NodeFetcherInterface, AssetsStore *assets.AssetsStore, gc gcache.Cache) {
	api.CreateAccountHandler = NewCreateAccount(prompterApp, massaClient)
	api.DeleteAccountHandler = NewDelete(prompterApp, massaClient)
	api.ImportAccountHandler = NewImport(prompterApp, massaClient)
	api.AccountListHandler = NewGetAll(prompterApp.App().WalletManager, massaClient)
	api.SignHandler = NewSign(prompterApp, gc)
	api.SignMessageHandler = NewSignMessage(prompterApp, gc)
	api.GetAccountHandler = NewGet(prompterApp, massaClient)
	api.ExportAccountFileHandler = NewWalletExportFile(prompterApp.App().WalletManager)
	api.TransferCoinHandler = NewTransferCoin(prompterApp, massaClient)
	api.TradeRollsHandler = NewTradeRolls(prompterApp, massaClient)
	api.BackupAccountHandler = NewBackupAccount(prompterApp)
	api.UpdateAccountHandler = NewUpdateAccount(prompterApp, massaClient)
	api.AddAssetHandler = NewAddAsset(AssetsStore, massaClient)
	api.GetAllAssetsHandler = NewGetAllAssets(prompterApp.App().WalletManager, AssetsStore, massaClient)
	api.DeleteAssetHandler = NewDeleteAsset(AssetsStore)
}

// loadAccount loads a wallet from the file system or returns an error.
func loadAccount(wallet *walletmanager.Wallet, nickname string) (*account.Account, middleware.Responder) {
	w, err := wallet.GetAccount(nickname)
	if err == nil {
		return w, nil
	}

	errorObj := models.Error{
		Code:    errorGetWallets,
		Message: err.Error(),
	}

	if errors.Is(err, walletmanager.AccountNotFoundError) {
		return nil, operations.NewGetAccountNotFound().WithPayload(&errorObj)
	} else {
		return nil, operations.NewGetAccountBadRequest().WithPayload(&errorObj)
	}
}

func newInternalServerError(message, code string) middleware.Responder {
	logger.Error(message)

	return operations.NewGetAccountInternalServerError().WithPayload(&models.Error{
		Code:    code,
		Message: message,
	})
}

func newBadRequest(message, code string) middleware.Responder {
	logger.Error(message)

	return operations.NewSignBadRequest().WithPayload(
		&models.Error{
			Code:    code,
			Message: message,
		})
}
