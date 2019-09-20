package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TransRequest struct {
	Method    string                 `json:"method"`
	Arguments map[string]interface{} `json:"arguments"`
	Tag       int                    `json:"tag"`
}

func TransAddTorrent(url string) {
	tr := TransRequest{
		Method: "torrent-add",
		Arguments: map[string]interface{}{
			"filename":     url,
			"download-dir": "/mnt/storage1/manga",
		},
	}
	b, _ := json.Marshal(tr)

	buf := bytes.NewBuffer(b)
	resp, err := http.Post("http://nas3.local:9091/transmission/rpc", "image/jpeg", buf)

	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}

func TransGetSession() string {
	b := make([]byte, 0)
	buf := bytes.NewBuffer(b)
	resp, err := http.Post("http://nas3.local:9091/transmission/rpc", "image/jpeg", buf)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return resp.Header.Get("X-Transmission-Session-Id")
}
