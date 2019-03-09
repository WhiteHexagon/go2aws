package cognito

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var client = http.Client{}

// DoPost - web call to AWS
func DoPost(url string, target string, request interface{}, response interface{}) error {
	// Prepare Request
	jsonReq, err := json.Marshal(request)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-amz-json-1.1")
	req.Header.Set("X-AMZ-TARGET", target)
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	// Process Response
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, response)
	if err != nil {
		return err
	}
	return nil
}
