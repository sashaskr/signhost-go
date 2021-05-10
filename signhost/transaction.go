package signhost

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type TransactionService service

type Transaction struct {
	Id                     string   `json:"Id,omitempty"`
	Files                  []File   `json:"FileEntry,omitempty"`
	Seal                   bool     `json:"Seal,omitempty"`
	Signers                []Signer `json:"Signers"`
	Reference              string   `json:"Reference,omitempty"`
	PostbackUrl            string   `json:"PostbackUrl,omitempty"`
	SignRequestMode        int32    `json:"SignRequestMode,omitempty"`
	DaysToExpire           int32    `json:"DaysToExpire,omitempty"`
	SendEmailNotifications bool     `json:"SendEmailNotifications,omitempty"`
	Status                 int32    `json:"Status,omitempty"`
	CancellationReason     string   `json:"CancellationReason,omitempty"`
	Context                string   `json:"Context,omitempty"`
}

type Signer struct {
	Id                   string `json:"Id,omitempty"`
	Email                string `json:"Email,omitempty"`
	IntroText            string `json:"IntroText,omitempty"`
	Authentications      []Authentication
	Verifications        []Verification
	SendSignRequest      bool       `json:"SendSignRequest,omitempty"`
	SignUrl              string     `json:"SignUrl,omitempty"`
	SignRequestSubject   string     `json:"SignRequestSubject,omitempty"`
	SignRequestMessage   string     `json:"SignRequestMessage,omitempty"`
	SendSignConfirmation bool       `json:"SendSignConfirmation,omitempty"`
	Language             string     `json:"language,omitempty"`
	ScribbleName         string     `json:"ScribbleName,omitempty"`
	DaysToRemind         int        `json:"DaysToRemind,omitempty"`
	Expires              string     `json:"Expires,omitempty"`
	Reference            string     `json:"Reference,omitempty"`
	RejectReason         string     `json:"RejectReason,omitempty"`
	ReturnUrl            string     `json:"ReturnUrl,omitempty"`
	Context              string     `json:"Context,omitempty"`
	Activities           []Activity `json:"Activities,omitempty"`
	Receivers            []Receiver `json:"Receivers,omitempty"`
}

type Authentication struct {
	Type string `json:"Type,omitempty"`
}

type Verification struct {
	Type string `json:"Type,omitempty"`
}

type File struct {
	Links       []Link `json:"Links,omitempty"`
	DisplayName string `json:"DisplayName,omitempty"`
}

type Link struct {
	Rel  string `json:"Rel,omitempty"`
	Type string `json:"Type,omitempty"`
	Link string `json:"Link,omitempty"`
}

type Activity struct {
	Id              string `json:"Id,omitempty"`
	Code            int    `json:"Code,omitempty"`
	Info            string `json:"Info,omitempty"`
	CreatedDateTime string `json:"CreatedDateTime,omitempty"`
}

type Receiver struct {
	Name      string `json:"Name,omitempty"`
	Email     string `json:"Email,omitempty"`
	Language  string `json:"Language,omitempty"`
	Subject   string `json:"Subject,omitempty"`
	Message   string `json:"Message,omitempty"`
	Reference string `json:"Reference,omitempty"`
	Context   string `json:"Context,omitempty"`
}

func (ts *TransactionService) Post(t *Transaction) (tt *Transaction, err error) {
	req, err := ts.client.NewAPIRequest(http.MethodPost, "transaction", t)
	if err != nil {
		return
	}

	res, err := ts.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &tt); err != nil {
		return
	}
	return
}

func (ts *TransactionService) Get(transactionID string) (tt *Transaction, err error) {
	u := fmt.Sprintf("transaction/%s", transactionID)
	req, err := ts.client.NewAPIRequest(http.MethodGet, u, nil)
	if err != nil {
		return
	}

	res, err := ts.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &tt); err != nil {
		log.Fatal(err)
		return
	}

	return
}

func (tr *Transaction) AddSigner(signer *Signer) []Signer {
	return append(tr.Signers, *signer)
}

func (s *Signer) AddVerification(verification *Verification) []Verification {
	return append(s.Verifications, *verification)
}
