package signhost

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type FileService service

type File struct {
	TransactionID string `json:"transaction_id,omitempty"`
	FileID        string `json:"file_id,omitempty"`
	filePath      string
}

type FilePDF struct {
	FilePath   string
	FileDigest string
}

func (fs *FileService) Put(file File) (f *File, err error) {
	u := fmt.Sprintf("transaction/%s/file/%s", file.TransactionID, file.FileID)
	pdf := CreatePdfFile(file.filePath)
	fileTemp, err := os.Open(pdf.FilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer fileTemp.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("application/pdf", filepath.Base(fileTemp.Name()))
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(part, fileTemp)
	writer.Close()

	req, err := fs.client.NewAPIRequest(http.MethodPut, u, body)
	req.Header.Add("Digest", pdf.FileDigest)
	req.Header.Set("Content-Type", "application/pdf")

	res, err := fs.client.Do(req)
	if err = json.Unmarshal(res.content, &fs); err != nil {
		return
	}
	return
}

func (f *File) SetFile(transactionID, fileID, filePath string) *File {
	f.FileID = fileID
	f.TransactionID = transactionID
	f.filePath = filePath
	return f
}

func CreatePdfFile(path string) *FilePDF {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}
	var h string = fmt.Sprintf("%x", hash.Sum(nil))

	return &FilePDF{
		FilePath:   path,
		FileDigest: h,
	}
}
