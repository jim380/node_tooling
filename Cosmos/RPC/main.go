package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Results struct {
	Results []Result
}

type Result struct {
	NodeInfo struct {
		ID      int    `json:"id"`
		Network string `json:"network"`
		Moniker string `json:"moniker"`
	} //`json:"node_info"`
	SyncInfo struct {
		NodeBlockHeight string `json:"latest_block_height"`
		CATCHUP         string `json:"catching_up"`
	} //`json:"sync_info"`
}

func main() {

	url := "http://207.246.72.35:26657/status"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	//var results Results
	//var result Result
	//json.Unmarshal(body, &results)
	people1 := &Results{}
	jsonErr := json.Unmarshal(body, people1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Println(people1.Results[0].NodeInfo.ID)
}
