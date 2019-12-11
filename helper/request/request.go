package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	u "net/url"
	"strconv"
)

type Request struct {
	Client *http.Client
}

var (
	METHODGET    = http.MethodGet
	METHODPOST   = http.MethodPost
	METHODPUT    = http.MethodPut
	METHODDELETE = http.MethodDelete
)

func RequestGET(url string, header map[string]interface{}, data map[string]interface{}) (map[string]interface{}, error) {
	var value interface{}
	var err error
	var req *http.Request

	client := http.DefaultClient

	req, err = http.NewRequest(METHODGET, url, nil)
	if err != nil {
		return nil, err
	}
	if header != nil {
		if len(header) > 0 {
			if v, has := header["access_token"]; has {
				accessToken := "Bearer " + v.(string)
				req.Header.Add("Authorization", accessToken)
				delete(header, "access_token")
			}
			for key, val := range header {
				req.Header.Add(key, val.(string))
			}
		}
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := bodyBytes
	json.Unmarshal(bodyString, &value)

	if resp.StatusCode != http.StatusOK {
		errMessage, _ := json.Marshal(value)
		return nil, errors.New(resp.Status + " " + string(errMessage))
	}

	return value.(map[string]interface{}), nil
}

func RequestPOST(url string, header map[string]interface{}, data map[string]interface{}) (map[string]interface{}, error) {
	var value interface{}
	var err error
	var req *http.Request

	client := http.DefaultClient
	if data != nil {
		reqBody := u.Values{}
		var JsonData, _ = json.Marshal(data)
		reqBody.Set("data", string(JsonData))

		var body io.Reader
		body = bytes.NewBufferString(reqBody.Encode())
		rc, ok := body.(io.ReadCloser)
		if !ok && body != nil {
			rc = ioutil.NopCloser(body)
		}

		req, err = http.NewRequest(METHODPOST, url, rc)
		req.Header.Add("Content-Length", strconv.Itoa(len(reqBody.Encode())))
	} else {
		req, err = http.NewRequest(METHODPOST, url, nil)
	}

	if err != nil {
		return nil, err
	}
	if v, has := header["access_token"]; has {
		accessToken := "Bearer " + v.(string)
		req.Header.Add("Authorization", accessToken)

	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := bodyBytes
	json.Unmarshal(bodyString, &value)

	if resp.StatusCode != http.StatusOK {
		errMessage, _ := json.Marshal(value)
		return nil, errors.New(resp.Status + " " + string(errMessage))
	}

	return value.(map[string]interface{}), nil

}

func RequestPUT(url string, header map[string]interface{}, data map[string]interface{}) (map[string]interface{}, error) {
	var value interface{}
	var err error
	var req *http.Request

	client := http.DefaultClient
	if data != nil {
		reqBody := u.Values{}
		var JsonData, _ = json.Marshal(data)
		reqBody.Set("data", string(JsonData))

		var body io.Reader
		body = bytes.NewBufferString(reqBody.Encode())
		rc, ok := body.(io.ReadCloser)
		if !ok && body != nil {
			rc = ioutil.NopCloser(body)
		}

		req, err = http.NewRequest(METHODPUT, url, rc)
		req.Header.Add("Content-Length", strconv.Itoa(len(reqBody.Encode())))
	} else {
		req, err = http.NewRequest(METHODPUT, url, nil)
	}

	if err != nil {
		return nil, err
	}
	if v, has := header["access_token"]; has {
		accessToken := "Bearer " + v.(string)
		req.Header.Add("Authorization", accessToken)

	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := bodyBytes
	json.Unmarshal(bodyString, &value)

	if resp.StatusCode != http.StatusOK {
		errMessage, _ := json.Marshal(value)
		return nil, errors.New(resp.Status + " " + string(errMessage))
	}

	return value.(map[string]interface{}), nil
}

func RequestRevokeToken(url string, data map[string][]string) (map[string]interface{}, error) {
	var value interface{}

	client := http.DefaultClient
	req, err := http.NewRequest(METHODDELETE, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("refresh_token", data["refresh_token"][0])
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := bodyBytes
	json.Unmarshal(bodyString, &value)

	if resp.StatusCode != http.StatusOK {
		errMessage, _ := json.Marshal(value)
		return nil, errors.New(resp.Status + " " + string(errMessage))
	}

	return value.(map[string]interface{}), nil
}
