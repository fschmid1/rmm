package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Post(url string, data interface{}, respBody interface{}) (int, string, error) {

	byteData, _ := json.MarshalIndent(data, "", " ")
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(byteData))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err.Error(), err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode != 200 {
		return resp.StatusCode, string(body), nil
	}
	json.Unmarshal(body, respBody)
	return resp.StatusCode, "", nil
}
