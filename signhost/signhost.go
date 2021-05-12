package signhost

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	BaseURL                string = "https://api.signhost.com/api/"
	AuthHeader             string = "Authorization"
	ApplicationHeader      string = "Application"
	DigestHeader           string = "Digest"
	RequestContentType     string = "application/json"
	RequestPdfContentType  string = "application/pdf"
	RequestAccept          string = "application/vnd.signhost.v1+json"
	TokenTypeAuthorization string = "APIKey"
	TokenTypeApplication   string = "APPKey"
	Connection             string = "keep-alive"
	APITokenEnv            string = "SIGNHOST_API_TOKEN"
	AppKeyEnv              string = "SIGNHOST_APP_KEY"
	ContentTypeHeader      string = "Content-Type"
)

type Response struct {
	*http.Response
	content []byte
}

type Client struct {
	BaseURL        *url.URL
	authentication string
	application    string
	userAgent      string
	client         *http.Client
	config         *Config
	common         service
	Transaction    *TransactionService
	File           *FileService
}

var (
	errEmptyAuthKey = errors.New("you must provide a non-empty authentication key")
	errEmptyAppKey  = errors.New("you mush provide a non-empty application key")
	errBadBaseURL   = errors.New("malformed base url, it must contain a trailing slash")
)

type service struct {
	client *Client
}

func (c *Client) WithAuthenticationValue(k string, p string) error {
	if k == "" {
		return errEmptyAuthKey
	}

	if p == "" {
		return errEmptyAppKey
	}

	c.authentication = strings.TrimSpace(k)
	c.application = strings.TrimSpace(p)
	return nil
}

func (c *Client) NewAPIRequest(method string, uri string, body interface{}) (req *http.Request, err error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, errBadBaseURL
	}

	u, err := c.BaseURL.Parse(uri)
	if err != nil {
		return nil, err
	}

	if c.config.testing {
		u.Query().Add("testmode", "true")
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err = http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add(AuthHeader, strings.Join([]string{TokenTypeAuthorization, c.authentication}, " "))
	req.Header.Add(ApplicationHeader, strings.Join([]string{TokenTypeApplication, c.application}, " "))
	req.Header.Set("Content-Type", RequestContentType)
	req.Header.Set("Accept", RequestAccept)
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Connection", Connection)
	return
}

func (c *Client) Do(req *http.Request) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response, _ := newResponse(resp)
	err = CheckResponse(resp)
	if err != nil {
		return response, err
	}

	return response, nil
}

func NewClient(baseClient *http.Client, c *Config) (signhost *Client, err error) {
	if baseClient == nil {
		baseClient = http.DefaultClient
	}

	u, _ := url.Parse(BaseURL)

	signhost = &Client{
		BaseURL: u,
		client:  baseClient,
		config:  c,
	}

	signhost.common.client = signhost

	// services for resources
	signhost.Transaction = (*TransactionService)(&signhost.common)
	signhost.File = (*FileService)(&signhost.common)
	// services end
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	viper.SetConfigFile(basepath + "/../.env")
	//viper.AutomaticEnv()
	_ = viper.ReadInConfig()
	var token, appKey string
	var okToken, okAppKey bool
	if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		log.Println("Not .env file found, searching for OS ENV...")
		token, okToken = os.LookupEnv(c.auth)
		if !okToken {
			log.Fatalf("Error while reading os.env=%s", c.auth)
		}
		appKey, okAppKey = os.LookupEnv(c.appKey)
		if !okAppKey {
			log.Fatalf("Error while reading os.env=%s", c.appKey)
		}
	} else {
		token, okToken = viper.Get(c.auth).(string)
		appKey, okAppKey = viper.Get(c.appKey).(string)
	}

	if okToken && okAppKey {
		signhost.authentication = token
		signhost.application = appKey
		signhost.userAgent = strings.Join([]string{
			runtime.GOOS,
			runtime.GOARCH,
			runtime.Version(),
		}, ";")
	}
	return
}

type Error struct {
	Code     int            `json:"code"`
	Message  string         `json:"message"`
	Content  string         `json:"content,omitempty"`
	Response *http.Response `json:"response"` // the full response that produced the error
}

func (e *Error) Error() string {
	return fmt.Sprintf("response failed with status %s\npayload: %v", e.Message, e.Content)
}

func newError(r *http.Response) *Error {
	var e Error
	e.Response = r
	e.Code = r.StatusCode
	e.Message = r.Status
	c, err := ioutil.ReadAll(r.Body)
	if err == nil {
		e.Content = string(c)
	}
	r.Body = ioutil.NopCloser(bytes.NewBuffer(c))
	return &e
}

func newResponse(r *http.Response) (*Response, error) {
	var res Response
	c, err := ioutil.ReadAll(r.Body)
	if err == nil {
		res.content = c
	}
	err = json.NewDecoder(r.Body).Decode(&res)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(c))
	res.Response = r
	return &res, err
}

func CheckResponse(r *http.Response) error {
	if r.StatusCode >= http.StatusMultipleChoices {
		return newError(r)
	}
	return nil
}
