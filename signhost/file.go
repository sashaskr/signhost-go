package signhost

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
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
	FilePath      string `json:"file_path"`
	ContentType   string `json:"ContentType"`
}

type FileMetaData struct {
	DisplayOrder int                    `json:"DisplayOrder,omitempty"`
	DisplayName  string                 `json:"DisplayName,omitempty"`
	SetParaph    bool                   `json:"SetParaph,omitempty"`
	Signers      map[string]FormSigners `json:"Signers,omitempty"`
	FormSets     map[string]FormSets    `json:"FormSets,omitempty"`
}

type FormSigners struct {
	FormSets []string `json:"FormSets,omitempty"`
}

type FormSets struct {
	Signatures map[string]Signature `json:"Signatures,omitempty"`
}

type Signature struct {
	Type     string   `json:"Type,omitempty"`
	Location Location `json:"Location,omitempty"`
}

type Location struct {
	Search     string `json:"Search,omitempty"`
	Occurence  int    `json:"Occurence,omitempty"`
	Top        int    `json:"Top,omitempty"`
	Right      int    `json:"Right,omitempty"`
	Bottom     int    `json:"Bottom,omitempty"`
	Left       int    `json:"Left,omitempty"`
	Width      int    `json:"Width,omitempty"`
	Height     int    `json:"Height,omitempty"`
	PageNumber int    `json:"PageNumber,omitempty"`
}

type FilePDF struct {
	FilePath   string
	FileDigest string
}

func (fs *FileService) Put(file File, meta interface{}) (v interface{}, err error) {
	u := fmt.Sprintf("transaction/%s/file/%s", file.TransactionID, file.FileID)
	var body interface{}
	switch file.ContentType {
	case RequestContentType:
		body = meta
	case RequestPdfContentType:
		body = GetPdfRequestBody(file)
	}
	req, err := fs.client.NewAPIRequest(http.MethodPut, u, body)
	//req.Header.Add("Digest", pdf.FileDigest) TODO: Add file digest functionality
	req.Header.Set("Content-Type", file.ContentType)

	res, err := fs.client.Do(req)
	if err = json.Unmarshal(res.content, &fs); err != nil {
		return
	}
	return
}

func GetPdfRequestBody(file File) interface{} {
	pdf := CreatePdfFile(file.FilePath)
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
	return body
}

func (file *File) SetFile(transactionID, fileID, filePath string) *File {
	file.FileID = fileID
	file.TransactionID = transactionID
	file.FilePath = filePath
	return file
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
	encoded := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	var h = fmt.Sprintf("SHA256=%s", encoded)

	return &FilePDF{
		FilePath:   path,
		FileDigest: h,
	}
}
