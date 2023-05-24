package prompt

import (
	"fmt"
	"strings"

	walletapp "github.com/massalabs/thyra-plugin-wallet/pkg/app"
	"github.com/massalabs/thyra-plugin-wallet/pkg/utils"
	"github.com/massalabs/thyra-plugin-wallet/pkg/wallet"
)

func handleImportPrompt(prompterApp WalletPrompterInterface, input interface{}) (*wallet.Wallet, bool, error) {
	filePath, ok := input.(string)
	if ok {
		return handleImportFile(prompterApp, filePath)
	}

	walletInfo, ok := input.(walletapp.ImportFromPKey)
	if ok {
		return handleImportPrivateKey(prompterApp, walletInfo)
	}

	return nil, false, InputTypeError(prompterApp)
}

func handleImportFile(prompterApp WalletPrompterInterface, filePath string) (*wallet.Wallet, bool, error) {
	if !strings.HasSuffix(filePath, ".yml") {
		prompterApp.EmitEvent(walletapp.PromptResultEvent,
			walletapp.EventData{Success: false, CodeMessage: utils.ErrAccountFile})
		return nil, false, fmt.Errorf(InvalidAccountFileErr)
	}
	account, loadErr := wallet.LoadFile(filePath)
	if loadErr != nil {
		errStr := fmt.Sprintf("%v: %v", AccountLoadErr, loadErr.Err)
		prompterApp.EmitEvent(walletapp.PromptResultEvent,
			walletapp.EventData{Success: false, CodeMessage: loadErr.CodeErr})
		return nil, false, fmt.Errorf(errStr)
	}

	// Validate nickname
	if !wallet.NicknameIsValid(account.Nickname) {
		errorCode := utils.ErrInvalidNickname
		fmt.Println(errorCode)
		prompterApp.EmitEvent(walletapp.PromptResultEvent,
			walletapp.EventData{Success: false, CodeMessage: errorCode})
		return nil, false, fmt.Errorf(errorCode)
	}

	// Validate nickname uniqueness
	err := wallet.NicknameIsUnique(account.Nickname)
	if err != nil {
		errorCode := utils.ErrDuplicateNickname
		fmt.Println(errorCode)
		prompterApp.EmitEvent(walletapp.PromptResultEvent,
			walletapp.EventData{Success: false, CodeMessage: errorCode})
		return nil, false, fmt.Errorf(errorCode)
	}

	return &account, false, nil
}

func handleImportPrivateKey(prompterApp WalletPrompterInterface, walletInfo walletapp.ImportFromPKey) (*wallet.Wallet, bool, error) {
	wallet, importErr := wallet.Import(walletInfo.Nickname, walletInfo.PrivateKey, walletInfo.Password)
	if importErr != nil {
		errStr := fmt.Sprintf("%v: %v", ImportPrivateKeyErr, importErr.Err)
		prompterApp.EmitEvent(walletapp.PromptResultEvent,
			walletapp.EventData{Success: false, CodeMessage: importErr.CodeErr})
		return nil, false, fmt.Errorf(errStr)
	}

	return wallet, false, nil
}
