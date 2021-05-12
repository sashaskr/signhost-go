package signhost

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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
	if file.ContentType == RequestPdfContentType {
		pdfDigest := CreatePdfFile(file.FilePath)
		req.Header.Add("Digest", pdfDigest.FileDigest)
	}
	req.Header.Set("Content-Type", file.ContentType)

	res, err := fs.client.Do(req)
	if err = json.Unmarshal(res.content, &fs); err != nil {
		return
	}
	return
}

func (fs *FileService) Get(file File) (interface{}, error) {
	u := fmt.Sprintf("transaction/%s/file/%s", file.TransactionID, file.FileID)
	req, err := fs.client.NewAPIRequest(http.MethodGet, u, nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := fs.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(file.FilePath, res.content, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return nil, nil
}

func GetPdfRequestBody(file File) interface{} {
	fileTemp, err := os.Open(file.FilePath)
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
	fileDigest := GenerateDigest(file)

	return &FilePDF{
		FilePath:   path,
		FileDigest: fileDigest,
	}
}

func GenerateDigest(file *os.File) string {
	h := sha256.New()
	if _, err := io.Copy(h, file); err != nil {
		log.Fatal(err)
	}
	sha256hash := fmt.Sprintf("%x", h.Sum(nil))
	sEnc := base64.StdEncoding.EncodeToString([]byte(sha256hash))
	return fmt.Sprintf("SHA256=%s", sEnc)
}
