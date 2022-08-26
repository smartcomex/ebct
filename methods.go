package ebct

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

const (
	ProductionUrl   = "https://api.correios.com.br/"
	SandBoxUrl      = "https://apihom.correios.com.br/"
	DefaultEndPoint = "packet/"
	Version         = "v1/"
)

// Map data por post in the request
type Params map[string]interface{}

// Map extra request headers
type Headers map[string]string

// Make request and return the response
func (c *Client) execute(method string, path string, params interface{}, headers Headers, model interface{}) error {

	var request *http.Request

	// mount endpoint
	var endpoint = c.GetEndpoint() + path

	spew.Dump(endpoint)

	// check for params
	if params != nil {

		// send as body
		if method != http.MethodGet {

			// marshal params
			b, err := json.Marshal(params)
			if err != nil {
				return err
			}

			// set payload with params
			payload := strings.NewReader(string(b))

			// set request with payload
			request, _ = http.NewRequest(method, endpoint, payload)

		} else {

			// init request
			request, _ = http.NewRequest(method, endpoint, nil)

			// init query string
			query := request.URL.Query()

			var param Params
			inrec, _ := json.Marshal(params)
			_ = json.Unmarshal(inrec, &param)

			// add params
			for key, value := range param {
				switch reflect.TypeOf(value).Kind() {
				case reflect.Slice:
					for _, val := range value.([]interface{}) {
						query.Add(key, val.(string))
					}
				default:
					query.Add(key, fmt.Sprintf("%v", value))
				}
			}

			// set query string
			request.URL.RawQuery = query.Encode()

		}

	} else {

		// set request without payload
		request, _ = http.NewRequest(method, endpoint, nil)

	}

	// set header
	if c.GetToken() != "" {
		request.Header.Add("Authorization", c.GetToken()) // Chave de autenticação
	} else {
		request.Header.Add("Authorization", c.BasicAuth)
	}
	request.Header.Add("platform", "smartcomex")     // Plataforma origem da requisição
	request.Header.Add("platform-version", "v1.0.0") // A versão atual da plataforma.
	request.Header.Add("plugin", "correios-plugin")  // Nome do conector utilizado.
	request.Header.Add("plugin-version", "v1.0.0")   // Versão do conector utilizado.
	request.Header.Add("accept", "application/json")
	request.Header.Add("content-type", "application/json") // Enviar sempre

	// add extra headers
	if headers != nil {
		for key, value := range headers {
			request.Header.Add(key, value)
		}
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	// read response
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// init error message
	erm := &ErrorMessage{}

	// check for error message
	if err = json.Unmarshal(data, erm); err == nil && erm.Error() != "" {
		return erm
	}

	// init generic error
	message := &GenericMessage{}

	// parse generic message
	_ = json.Unmarshal(data, message)

	// verify status code
	if NotIn(response.StatusCode, http.StatusOK, http.StatusCreated, http.StatusAccepted) {

		// return generic error message
		if message.Error() != "" {
			return message
		}

		// return body as error message
		if len(data) > 0 {
			return errors.New(string(data))
		}

		// return http status as error
		return errors.New(response.Status)

	}

	// some services have empty response array
	if len(data) == 2 && string(data) == "[]" {
		return nil
	}

	// some services have empty response
	if len(data) == 0 {
		return nil
	}

	// pdf
	if len(data) > 3 && string(data)[1:4] == "PDF" {
		*model.(*[]byte) = data
		return nil
	}

	// xml
	if len(data) > 5 && In(string(data)[:5], "<?xml", "<Nfse", "<NFe ", "<nfeP") {
		*model.(*string) = string(data)
		return nil
	}

	// parse data
	return json.Unmarshal(data, model)

}

// Execute GET requests
func (c *Client) Get(path string, params interface{}, headers Headers, model interface{}) error {
	return c.execute(http.MethodGet, path, params, headers, model)
}

// Execute POST requests
func (c *Client) Post(path string, params interface{}, headers Headers, model interface{}) error {
	return c.execute(http.MethodPost, path, params, headers, model)
}

// Execute PUT requests
func (c *Client) Put(path string, params interface{}, headers Headers, model interface{}) error {
	return c.execute(http.MethodPut, path, params, headers, model)
}

// Execute PATCH requests
func (c *Client) Patch(path string, params interface{}, headers Headers, model interface{}) error {
	return c.execute(http.MethodPatch, path, params, headers, model)
}

// Execute DELETE requests
func (c *Client) Delete(path string, params interface{}, headers Headers, model interface{}) error {
	return c.execute(http.MethodDelete, path, params, headers, model)
}
