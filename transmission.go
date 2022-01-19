package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//TransRequest a request to Transmission.
type TransRequest struct {
	Method    string                 `json:"method"`
	Arguments map[string]interface{} `json:"arguments"`
	Tag       int                    `json:"tag"`
}

//TransAddResult a result from TransAddTorrent function.
type TransAddResult struct {
	Arguments interface{} `json:"arguments"`
	Result    string      `json:"result"`
	Tag       int         `json:"tag"`
}

//TransAddTorrent add a torrent into the transmission service.
func TransAddTorrent(rpc string, url string, session string, path string) (out TransAddResult, err error) {
	tr := TransRequest{
		Method: "torrent-add",
		Arguments: map[string]interface{}{
			"filename":     url,
			"download-dir": path,
		},
	}
	b, err := json.Marshal(tr)
	if err != nil {
		return
	}

	buf := bytes.NewBuffer(b)

	client := &http.Client{}

	req, err := http.NewRequest("POST", rpc, buf)
	if err != nil {
		return
	}
	req.Header.Add("X-Transmission-Session-Id", session)
	resp, err := client.Do(req)

	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	str := string(body)
	err = json.Unmarshal([]byte(str), &out)

	return
}

//TransGetSession get a transmission session id.
func TransGetSession(url string) (session string, err error) {
	b := make([]byte, 0)
	buf := bytes.NewBuffer(b)
	resp, err := http.Post(url, "image/jpeg", buf)

	if err != nil {
		return
	}

	session = resp.Header.Get("X-Transmission-Session-Id")

	return
}
