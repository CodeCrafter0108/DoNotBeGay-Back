package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/eriscrypt/donotbegay/common/blockchain"
	"github.com/eriscrypt/donotbegay/gateway/internal"
	// "github.com/ethereum/go-ethereum/common"
	"github.com/go-chi/chi/v5"
)

const MaxImageUrlSymbols = 150

type TokenSaveResponse struct {
	Message string `json:"message"`
}

type TokenPreSaveResponse struct {
	Url string `json:"url"`
}

const (
	ErrMaxImageUrlSymbols    = "max image url symbols"
	ErrInsufficientBalance   = "insufficient balance"
	ErrorChainIdNotSupported = "chain not supported"
)

//	@Summary		Pre-save token image
//	@Description	Pre-save token image. Select a file to upload in formData.
//	@Tags			static
//	@Produce		json
//	@Success		200				{object}	TokenPreSaveResponse
//	@Failure		500				{object}	ErrorResponse
//	@Param			chainId			path		blockchain.ChainId	true	"Chain ID"
//	@Param			userAddress		path		string				true	"User address"
//	@Param			randomBytes		path		string				true	"Random bytes"
//	@Param			tournamentId	path		string				true	"Tournament ID"
//	@Param			file			formData	file				true	"File to upload (max 250KB)"
//	@Router			/static/tokens/create/{chainId}/{userAddress}/{randomBytes}/{tournamentId} [post]
func (app *application) tokenPreSaveHandler(w http.ResponseWriter, r *http.Request) {
	chainId, err := app.extractChainId(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// client, err := blockchain.NewClient(chainId)
	// if err != nil {
	// 	app.internalServerError(w, r, err)
	// 	return
	// }

	// userAddress := chi.URLParam(r, "userAddress")
	randomBytes := chi.URLParam(r, "randomBytes")
	tournamentId := chi.URLParam(r, "tournamentId")

	// userBalance, err := blockchain.GetUserBalance(client, common.HexToAddress(userAddress))
	// if err != nil {
	// 	app.internalServerError(w, r, err)
	// 	return
	// }

	// initialLiquidity, err := blockchain.GetInitialLiquidity(client, chainId)
	// if err != nil {
	// 	app.internalServerError(w, r, err)
	// 	return
	// }

	// // if user balance is less than initial liquidity then return bad request
	// if userBalance.Cmp(initialLiquidity) == -1 {
	// 	app.badRequestResponse(w, r, errors.New(ErrInsufficientBalance))
	// 	return
	// }

	url, err := internal.PreSaveFile(r, randomBytes, chainId, tournamentId)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if len(url) > MaxImageUrlSymbols {
		app.badRequestResponse(w, r, errors.New(ErrMaxImageUrlSymbols))
		return
	}

	err = app.jsonResponse(w, http.StatusOK, TokenPreSaveResponse{Url: url})
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

//	@Summary		Save token image
//	@Description	Save token image to the server
//	@Tags			static
//	@Produce		json
//	@Success		200			{object}	TokenSaveResponse
//	@Failure		500			{object}	ErrorResponse
//	@Param			randomBytes	path		string	true	"Random bytes"
//	@Router			/static/tokens/save/{randomBytes} [put]
func (app *application) tokenSaveHandler(w http.ResponseWriter, r *http.Request) {
	randomBytes := chi.URLParam(r, "randomBytes")

	err := internal.SaveFile(randomBytes)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = app.jsonResponse(w, http.StatusOK, TokenSaveResponse{Message: "File saved successfully"})
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

//	@Summary		Get token image
//	@Description	Get token image to the server
//	@Tags			static
//	@Produce		json
//	@Success		200				{file}		file
//	@Failure		500				{object}	ErrorResponse
//	@Param			chainId			path		blockchain.ChainId	true	"Chain ID"
//	@Param			tournamentId	path		string				true	"Tournament ID"
//	@Param			filename		path		string				true	"Filename"
//	@Router			/static/tokens/{chainId}/{tournamentId}/{filename} [get]
func (app *application) tokenGetHandler(w http.ResponseWriter, r *http.Request) {
	chainId, err := app.extractChainId(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	filename := chi.URLParam(r, "filename")
	tournamentId := chi.URLParam(r, "tournamentId")
	filePath := "static/" + strconv.Itoa(int(chainId)) + "/" + tournamentId + "/" + filename

	w.Header().Set("Cache-Control", "public, max-age=31536000")
	http.ServeFile(w, r, filePath)
}

// function to extract the chain id from the URL
func (app *application) extractChainId(r *http.Request) (blockchain.ChainId, error) {
	_chainId, err := strconv.Atoi(chi.URLParam(r, "chainId"))
	if err != nil {
		return 0, err
	}

	if !blockchain.IsSupportedChain(blockchain.ChainId(_chainId)) {
		return 0, errors.New(ErrorChainIdNotSupported)
	}

	return blockchain.ChainId(_chainId), nil
}
