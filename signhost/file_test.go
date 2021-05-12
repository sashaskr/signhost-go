package signhost

import (
	"encoding/json"
	"net/http"
	"os"
	"signhost-client/testdata"
	"testing"
)

func TestFileService_PutWithMetadata(t *testing.T) {
	setEnv()
	setup()
	defer teardown()

	id := "67b9e140-b0f6-4c36-9b35-977041968185"
	fileID := "Contract.pdf"
	_ = tClient.WithAuthenticationValue("banana_token", "dsfksdjfksjdlfs")
	tMux.HandleFunc("/transaction/"+id+"/file/"+fileID, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "APIKey banana_token")
		testHeader(t, r, ApplicationHeader, "APPKey dsfksdjfksjdlfs")
		testHeader(t, r, ContentTypeHeader, RequestContentType)

		testMethod(t, r, "PUT")

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}
		if _, ok := r.Header[ApplicationHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.PutFileResponse))
	})

	var fileMetaData = &FileMetaData{}
	err := json.Unmarshal([]byte(testdata.PutFileWithMetaRequest), fileMetaData)
	if err != nil {
		t.Error(err)
	}

	file := File{
		TransactionID: id,
		FileID:        fileID,
		ContentType:   "application/json",
	}

	_, err = tClient.File.Put(file, fileMetaData)
	if err != nil {
		t.Error(err)
	}
	// TODO: more checks
}

func TestGenerateDigest(t *testing.T) {
	file, err := os.Open("../testdata/sample.pdf")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	digest := GenerateDigest(file)
	if digest != testdata.SampleFileDigest {
		t.Errorf("Digests of sample.pdf files are not mathed. Expected %s, got %s", testdata.SampleFileDigest, digest)
	}
}

func TestFileService_PutWithPdfFile(t *testing.T) {
	setEnv()
	setup()
	defer teardown()

	id := "67b9e140-b0f6-4c36-9b35-977041968185"
	fileID := "Contract.pdf"
	_ = tClient.WithAuthenticationValue("banana_token", "dsfksdjfksjdlfs")

	tMux.HandleFunc("/transaction/"+id+"/file/"+fileID, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "APIKey banana_token")
		testHeader(t, r, ApplicationHeader, "APPKey dsfksdjfksjdlfs")
		testHeader(t, r, ContentTypeHeader, RequestPdfContentType)
		testHeader(t, r, DigestHeader, "SHA256=OGRlY2M4NTcxOTQ2ZDRjZDcwYTAyNDk0OWUwMzNhMmEyYTU0Mzc3ZmU5ZjFjMWI5NDRjMjBmOWVlMTFhOWU1MQ==")
		testMethod(t, r, "PUT")

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		if _, ok := r.Header[ApplicationHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		if _, ok := r.Header[DigestHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.PutFileResponse))
	})

	file := File{
		TransactionID: id,
		FileID:        fileID,
		ContentType:   "application/pdf",
		FilePath:      "../testdata/sample.pdf",
	}
	var meta interface{}
	_, err := tClient.File.Put(file, meta)
	if err != nil {
		t.Error(err)
	}
}
