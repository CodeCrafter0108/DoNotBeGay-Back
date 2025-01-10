package internal

import (
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/eriscrypt/donotbegay/common/blockchain"
)

const MaxImageSize = 250 * 1024 // 250KB

type File struct {
	Path string
	File multipart.File
}

var files = make(map[string]File)

const (
	ErrFileAlreadyExists = "file already exists"
	ErrFileNotPreSaved   = "file not pre-saved"
	ErrFileExceedsLimit  = "file size exceeds 250KB limit"
)

func getUrl(r *http.Request, chainId blockchain.ChainId, tournamentId string, fileName string) string {
	return "https://" + r.Host + "/v1/static/tokens/" + strconv.Itoa(int(chainId)) + "/" + tournamentId + "/" + fileName
}

func PreSaveFile(r *http.Request, randomBytes string, chainId blockchain.ChainId, tournamentId string) (string, error) {
	err := r.ParseMultipartForm(1 << 20) // 1 MB
	if err != nil {
		return "", err
	}

	// Retrieve the file from form data
	file, header, err := r.FormFile("file")
	if err != nil {
		return "", err
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	// Check the file size
	fileSize := r.ContentLength
	if fileSize > MaxImageSize {
		return "", errors.New(ErrFileExceedsLimit)
	}

	fileExtension := filepath.Ext(header.Filename)
	fileName := randomBytes + fileExtension
	filePath := filepath.Join("static", strconv.Itoa(int(chainId)), tournamentId, fileName)

	// Check if the file already exists
	if files[randomBytes].File != nil {
		return "", errors.New(ErrFileAlreadyExists)
	}

	// Check if the file already exists
	fileInfo, _ := os.Stat(filePath)
	if fileInfo != nil {
		return "", errors.New(ErrFileAlreadyExists)
	}

	files[randomBytes] = File{
		Path: filePath,
		File: file,
	}

	url := getUrl(r, chainId, tournamentId, fileName)

	return url, nil
}

func SaveFile(bytes string) error {
	file := files[bytes]
	if file.File == nil {
		return errors.New(ErrFileNotPreSaved)
	}

	// Create directories if they do not exist
	err := os.MkdirAll(filepath.Dir(file.Path), os.ModePerm)
	if err != nil {
		return err
	}

	// Check if the file already exists
	if _, err := os.Stat(file.Path); err == nil {
		return err
	}

	// Create the file
	out, err := os.Create(file.Path)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			return
		}
	}(out)

	// Copy the file content to the new file
	_, err = io.Copy(out, file.File)
	if err != nil {
		return err
	}

	// Delete the file from the map
	delete(files, bytes)

	return nil
}
