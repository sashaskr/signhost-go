package signhost

import (
	"encoding/json"
	"net/http"
	"signhost-client/testdata"
	"testing"
)

func TestTransactionService_Get(t *testing.T) {
	setEnv()
	setup()
	defer teardown()

	id := "67b9e140-b0f6-4c36-9b35-977041968185"
	_ = tClient.WithAuthenticationValue("banana_token", "dsfksdjfksjdlfs")
	tMux.HandleFunc("/transaction/"+id, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "APIKey banana_token")
		testHeader(t, r, ApplicationHeader, "APPKey dsfksdjfksjdlfs")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}
		if _, ok := r.Header[ApplicationHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetTransactionResponse))
	})

	c, err := tClient.Transaction.Get(id)
	if err != nil {
		t.Error(err)
	}

	if c.Id != id {
		t.Errorf("unexpected response: got %s, want %s", c.Id, id)
	}

	if c.Files["Contract.pdf"].Links[0].Link != "https://api.signhost.com/api/transaction/67b9e140-b0f6-4c36-9b35-977041968185/file/Contract.pdf" {
		t.Errorf("unexpected response: got %s, want %s", c.Files["Contract.pdf"].Links[0].Link, "https://api.signhost.com/api/transaction/67b9e140-b0f6-4c36-9b35-977041968185/file/Contract.pdf")
	}
	unsetEnv()
}

func TestTransactionService_Post(t *testing.T) {
	setEnv()
	setup()
	defer teardown()

	id := "67b9e140-b0f6-4c36-9b35-977041968185"
	_ = tClient.WithAuthenticationValue("banana_token", "dsfksdjfksjdlfs")
	tMux.HandleFunc("/transaction", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "APIKey banana_token")
		testHeader(t, r, ApplicationHeader, "APPKey dsfksdjfksjdlfs")
		testMethod(t, r, "POST")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}
		if _, ok := r.Header[ApplicationHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(testdata.CreateTransactionResponse))
	})

	var transaction = &Transaction{}
	err := json.Unmarshal([]byte(testdata.CreateTransactionRequest), &transaction)
	if err != nil {
		t.Error(err)
	}
	c, err := tClient.Transaction.Post(transaction)
	if err != nil {
		t.Error(err)
	}

	if c.Id != id {
		t.Errorf("unexpected response: got %s, want %s", c.Id, id)
	}

	if c.Signers[0].Email != transaction.Signers[0].Email {
		t.Errorf("unexpected response: got %s, want %s",c.Signers[0].Email, transaction.Signers[0].Email)
	}
	unsetEnv()
}