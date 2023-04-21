package wallet

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/massalabs/thyra-plugin-wallet/api/server/models"
	"github.com/massalabs/thyra-plugin-wallet/api/server/restapi/operations"
	"github.com/massalabs/thyra-plugin-wallet/pkg/wallet"
)

// HandleExportFile handles a export file request
// It will serve the yaml file so that the client can download it.
func HandleExportFile(params operations.RestWalletExportFileParams) middleware.Responder {
	wlt, err := wallet.Load(params.Nickname)
	if err == wallet.ErrorAccountNotFound {
		return operations.NewRestWalletExportFileNotFound().WithPayload(
			&models.Error{
				Code:    errorExportWallet,
				Message: err.Error(),
			})
	} else if err != nil {
		return operations.NewRestWalletExportFileBadRequest().WithPayload(
			&models.Error{
				Code:    errorExportWallet,
				Message: err.Error(),
			})
	}

	pathToWallet, err := wlt.FilePath()
	if err != nil {
		return operations.NewRestWalletExportFileInternalServerError().WithPayload(
			&models.Error{
				Code:    errorExportWallet,
				Message: err.Error(),
			})
	}

	responder := middleware.ResponderFunc(func(w http.ResponseWriter, _ runtime.Producer) {
		file, _ := os.Open(pathToWallet)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%q", filepath.Base(file.Name())))
		if _, err := io.Copy(w, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()
	})

	return responder
}
