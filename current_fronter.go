package main

/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/monaco-io/request"
)

func getCurrentFronterData(api string, systemID string) []byte {
	client := request.Client{
		URL:    api + "/s/" + systemID + "/fronters",
		Method: "GET",
	}
	resp, err := client.Do()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return resp.Data
}

func getFront(rawJSON []byte) {
	var data map[string]interface{}

	if err := json.Unmarshal(rawJSON, &data); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(data)
	fmt.Printf("\n")
	fmt.Println(data["members"])
	fmt.Println(data["timestamp"])
}
