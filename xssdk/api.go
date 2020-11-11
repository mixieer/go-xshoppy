package xssdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type XsApiClient struct {
	App     App
	baseURL *url.URL
	Order   OrderService
	Hook    HookService
	Shop    ShopService
	Client  *http.Client
}

var DefaultClient = &http.Client{}

func NewApiClient(a App) *XsApiClient {

	httpClient := http.DefaultClient
	c := &XsApiClient{
		App:    a,
		Client: httpClient,
	}
	c.Hook = &HookServiceOp{XsClient: c}
	c.Shop = &ShopServiceOp{XsClient: c}
	c.Order = &OrderServiceOp{XsClient: c}
	return c
}

func (xs *XsApiClient) NewRequest(method, urlStr string, body, options interface{}, customHeader map[string]string) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	// Make the full url based on the relative path
	u := xs.baseURL.ResolveReference(rel)

	// Add custom options
	if options != nil {
		optionsQuery, err := Values(options)
		if err != nil {
			return nil, err
		}

		for k, values := range u.Query() {
			for _, v := range values {
				optionsQuery.Add(k, v)
			}
		}
		u.RawQuery = optionsQuery.Encode()
	}

	// A bit of JSON ceremony
	var js []byte = nil

	if body != nil {
		js, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(js))
	if err != nil {
		return nil, err
	}

	if customHeader != nil {
		for k, v := range customHeader {
			req.Header.Add(k, v)
		}
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-SAIL-ACCESS-TOKEN", xs.App.SharedSecret)
	req.SetBasicAuth(xs.App.ApiKey, xs.App.Password)

	return req, nil
}

func (xs *XsApiClient) CreateAndDo(method, path string, data, options, resource interface{}, customHeader map[string]string) error {
	req, err := xs.NewRequest(method, path, data, options, customHeader)
	if err != nil {
		return err
	}
	err = xs.Do(req, resource)
	if err != nil {
		return err
	}

	return nil
}

func (xs *XsApiClient) Do(req *http.Request, v interface{}) error {

	resp, err := xs.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = CheckResponseError(resp)
	if err != nil {
		return err
	}
	if v != nil {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
		b := bytes.NewReader(body)
		decoder := json.NewDecoder(b)
		err := decoder.Decode(&v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (xs *XsApiClient) Get(path string, resource, options interface{}) error {
	return xs.CreateAndDo("GET", path, nil, options, resource, nil)
}

func (xs *XsApiClient) Post(path string, data, resource interface{}, customHeader map[string]string) error {
	return xs.CreateAndDo("POST", path, data, nil, resource, customHeader)
}
func (xs *XsApiClient) Put(path string, data, resource interface{}) error {
	return xs.CreateAndDo("PUT", path, data, nil, resource, nil)
}

func (xs *XsApiClient) Delete(path string, data interface{}) error {
	return xs.CreateAndDo("DELETE", path, data, nil, nil, nil)
}
