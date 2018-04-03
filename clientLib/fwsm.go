package fwsmAPIClient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	ErrNotImplemented = fmt.Errorf("Not implemented (yet?)")
)

type FwsmAPIClient struct {
	httpClient *http.Client

	host string
	port int
	user string
	pass string

	scheme string
}

type FwsmAPIClientNewArgs struct {
	Host string
	Port int
	User string
	Pass string

	Scheme string
}

func New(args *FwsmAPIClientNewArgs) *FwsmAPIClient {
	if args == nil {
		return &FwsmAPIClient{}
	}

	if args.Scheme == "" {
		args.Scheme = "http"
	}

	return &FwsmAPIClient{
		httpClient: &http.Client{
		},
		host: args.Host,
		port: args.Port,
		user: args.User,
		pass: args.Pass,

		scheme: args.Scheme,
	}
}

func parseError(body string) error {
	result := map[string]interface{}{}
	err := json.Unmarshal([]byte(body), &result)
	if err != nil {
		return err
	}
	status, ok := result["status"].(string)
	if !ok {
		return fmt.Errorf("There's no \"status\" field")
	}
	if status != "OK" {
		return fmt.Errorf("\"status\" value is not \"OK\": %v: %v", status, result["error_description"])
	}
	return nil
}

func (c *FwsmAPIClient) writeRequest(method, uri string) error {
	body, err := c.request(method, uri)
	if err != nil {
		return err
	}
	return parseError(body)
}

func (c *FwsmAPIClient) Reload() error {
	return c.writeRequest("PUT", "fwsm/reload")
}

func (c *FwsmAPIClient) Apply() error {
	return c.writeRequest("PUT", "fwsm/apply")
}

func (c *FwsmAPIClient) readRequest(uri string) (string, error) {
	body, err := c.request("GET", uri)
	if err != nil {
		return string(body), err
	}
	err = parseError(body)
	return string(body), err
}

func (c *FwsmAPIClient) CheckConnection() error {
	_, err := c.readRequest("fwsm/config")
	return err
}

func (c *FwsmAPIClient) request(method, uri string) (string, error) {
	req, err := http.NewRequest(method, c.scheme+"://"+c.host+"/"+uri, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.user, c.pass)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}

	r, e := ioutil.ReadAll(res.Body)
	res.Body.Close()

	return string(r), e
}

