package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Get(url string, respBody interface{}) (int, string, error) {

	req, _ := http.NewRequest("GET", url, nil)
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
