package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type TransRequest struct {
	Method    string                 `json:"method"`
	Arguments map[string]interface{} `json:"arguments"`
	Tag       int                    `json:"tag"`
}

func TransAddTorrent(rpc string, url string, session string, path string) (out string, err error) {
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
	out = string(body)

	return
}

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
