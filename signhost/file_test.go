package signhost

import (
	"encoding/json"
	"net/http"
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
	tMux.HandleFunc("/transaction/" + id + "/file/" + fileID, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "APIKey banana_token")
		testHeader(t, r, ApplicationHeader, "APPKey dsfksdjfksjdlfs")
		testHeader(t, r, ContentTypeHeader, RequestContentType)

		//testHeader(t, r, DigestHeader, "") TODO: Digest test
		testMethod(t, r, "PUT")

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}
		if _, ok := r.Header[ApplicationHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.PutFileWithMetaResponse))
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
