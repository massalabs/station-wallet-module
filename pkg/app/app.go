package walletapp

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/awnumar/memguard"
	"github.com/massalabs/station-massa-wallet/pkg/utils"
	"github.com/massalabs/station-massa-wallet/pkg/wallet/account"
	"github.com/massalabs/station-massa-wallet/pkg/walletmanager"
	"github.com/massalabs/station/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type WalletApp struct {
	Ctx           context.Context
	CtrlChan      chan PromptCtrl
	PromptInput   chan interface{}
	WalletManager *walletmanager.Wallet
	Shutdown      bool
	IsListening   bool
}

func (a *WalletApp) cleanExit() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if a.Ctx != nil {
		a.Shutdown = true
		runtime.Quit(a.Ctx)
	}
}

func NewWalletApp(wallet *walletmanager.Wallet) *WalletApp {
	app := &WalletApp{
		CtrlChan:      make(chan PromptCtrl),
		PromptInput:   make(chan interface{}),
		WalletManager: wallet,
		Shutdown:      false,
		IsListening:   false,
	}

	go app.cleanExit()

	return app
}

func IsTestMode() bool {
	return os.Getenv("WALLET_PASSWORD") != ""
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *WalletApp) Startup(ctx context.Context) {
	a.Ctx = ctx
}

func (a *WalletApp) BeforeClose(ctx context.Context) bool {
	a.Hide()

	if a.Shutdown {
		return false
	}

	// Send a cancel message to the prompt and do NOT shutdown
	if a.IsListening {
		fmt.Println("canceling prompt before closing")
		a.CtrlChan <- Cancel
	}

	runtime.WindowReloadApp(a.Ctx)

	return true
}

// SendPromptInput is binded to the frontend
func (a *WalletApp) SendPromptInput(input string) {
	a.PromptInput <- input
}

// AbortAction is bound to the frontend
// It sends a cancel message to the prompt
func (a *WalletApp) AbortAction() {
	if a.IsListening {
		fmt.Println("Abort action")
		a.CtrlChan <- Cancel
	}
}

func (a *WalletApp) Show() {
	runtime.WindowShow(a.Ctx)
}

func (a *WalletApp) Hide() {
	runtime.WindowHide(a.Ctx)
}

type selectFileResult struct {
	Err         string `json:"err"`
	CodeMessage string `json:"codeMessage"`
	FilePath    string `json:"filePath"`
	Nickname    string `json:"nickname"`
}

func (a *WalletApp) SelectAccountFile() selectFileResult {
	filePath, err := runtime.OpenFileDialog(a.Ctx, runtime.OpenDialogOptions{
		Title:   "Import account file",
		Filters: []runtime.FileFilter{{DisplayName: "Account File (*.yaml;*.yml)", Pattern: "*.yaml;*.yml"}},
	})
	if err != nil {
		logger.Errorf("error while selecting a file: %v", err)
		return selectFileResult{Err: err.Error(), CodeMessage: utils.ErrAccountFile}
	}

	if filePath == "" {
		return selectFileResult{Err: utils.ErrNoFile, CodeMessage: utils.ErrNoFile}
	}

	acc, err := a.WalletManager.Load(filePath)
	if err != nil {
		logger.Errorf("error while loading file: %v", err)
		return selectFileResult{Err: err.Error(), CodeMessage: utils.WailsErrorCode(err)}
	}

	return selectFileResult{FilePath: filePath, Nickname: acc.Nickname}
}

type ImportFromPKey struct {
	PrivateKey *memguard.LockedBuffer
	Password   *memguard.LockedBuffer
	Nickname   string
}

func (a *WalletApp) ImportPrivateKey(privateKeyText string, nickname string, password string) {
	guardedPrivateKey := memguard.NewBufferFromBytes([]byte(privateKeyText))
	guardedPassword := memguard.NewBufferFromBytes([]byte(password))

	a.PromptInput <- ImportFromPKey{PrivateKey: guardedPrivateKey, Nickname: nickname, Password: guardedPassword}
}

func (a *WalletApp) IsNicknameUnique(nickname string) bool {
	return a.WalletManager.NicknameIsUnique(nickname) != nil
}

func (a *WalletApp) IsNicknameValid(nickname string) bool {
	return account.NicknameIsValid(nickname)
}
