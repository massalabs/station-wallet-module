package wallet

import (
	"fmt"
	"strconv"

	"github.com/awnumar/memguard"
	"github.com/go-openapi/runtime/middleware"
	"github.com/massalabs/station-massa-wallet/api/server/models"
	"github.com/massalabs/station-massa-wallet/api/server/restapi/operations"
	walletapp "github.com/massalabs/station-massa-wallet/pkg/app"
	"github.com/massalabs/station-massa-wallet/pkg/network"
	"github.com/massalabs/station-massa-wallet/pkg/prompt"
	"github.com/massalabs/station-massa-wallet/pkg/utils"
	"github.com/massalabs/station-massa-wallet/pkg/wallet/account"
	"github.com/massalabs/station-massa-wallet/pkg/walletmanager"
	sendOperation "github.com/massalabs/station/pkg/node/sendoperation"
	"github.com/massalabs/station/pkg/node/sendoperation/buyrolls"
	"github.com/massalabs/station/pkg/node/sendoperation/sellrolls"
)

func NewTradeRolls(prompterApp prompt.WalletPrompterInterface, massaClient network.NodeFetcherInterface) operations.TradeRollsHandler {
	return &tradeRolls{prompterApp: prompterApp, massaClient: massaClient}
}

type tradeRolls struct {
	prompterApp prompt.WalletPrompterInterface
	massaClient network.NodeFetcherInterface
}

func (t *tradeRolls) Handle(params operations.TradeRollsParams) middleware.Responder {
	acc, resp := loadAccount(t.prompterApp.App().WalletManager, params.Nickname)
	if resp != nil {
		return resp
	}

	// convert amount to uint64
	amount, err := strconv.ParseUint(string(params.Body.Amount), 10, 64)
	if err != nil {
		return operations.NewTradeRollsBadRequest().WithPayload(
			&models.Error{
				Code:    errorTransferCoin,
				Message: "Error during amount conversion",
			})
	}

	// convert fee to uint64
	fee, err := strconv.ParseUint(string(params.Body.Fee), 10, 64)
	if err != nil {
		return operations.NewTradeRollsBadRequest().WithPayload(
			&models.Error{
				Code:    errorTransferCoin,
				Message: "Error during fee conversion",
			})
	}

	promptRequest := prompt.PromptRequest{
		Action: walletapp.TradeRolls,
		Msg:    fmt.Sprintf("%s %s rolls , with fee %s nonaMassa", *params.Body.Side, string(params.Body.Amount), string(params.Body.Fee)),
	}

	promptOutput, err := prompt.WakeUpPrompt(t.prompterApp, promptRequest, acc)
	if err != nil {
		return operations.NewTradeRollsUnauthorized().WithPayload(
			&models.Error{
				Code:    errorCanceledAction,
				Message: "Unable to unprotect wallet",
			})
	}

	guardedPassword, _ := promptOutput.(*memguard.LockedBuffer)

	operation, tradeRollError := doTradeRolls(acc, guardedPassword, amount, fee, *params.Body.Side, t.massaClient)
	if tradeRollError != nil {
		errStr := fmt.Sprintf("error %sing rolls coin: %v", *params.Body.Side, tradeRollError.Err.Error())
		t.prompterApp.EmitEvent(walletapp.PromptResultEvent,
			walletapp.EventData{Success: false, CodeMessage: tradeRollError.CodeErr})

		return operations.NewTradeRollsInternalServerError().WithPayload(
			&models.Error{
				Code:    errorTransferCoin,
				Message: errStr,
			})
	}

	t.prompterApp.EmitEvent(walletapp.PromptResultEvent,
		walletapp.EventData{Success: true, CodeMessage: utils.MsgRollTradeSuccess})

	return operations.NewTradeRollsOK().WithPayload(
		&models.OperationResponse{
			OperationID: operation.OperationID,
		})
}

func doTradeRolls(
	acc *account.Account,
	guardedPassword *memguard.LockedBuffer,
	amount, fee uint64,
	side string,
	massaClient network.NodeFetcherInterface,
) (*sendOperation.OperationResponse, *walletmanager.WalletError) {
	var operation sendOperation.Operation
	if side == "buy" {
		operation = buyrolls.New(amount)
	} else {
		operation = sellrolls.New(amount)
	}

	return network.SendOperation(acc, guardedPassword, massaClient, operation, fee)
}
