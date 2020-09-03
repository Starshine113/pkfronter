package main

/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

import (
	"encoding/json"

	"github.com/Starshine113/pkfronter/structs"
	"github.com/monaco-io/request"
)

func getCurrentFronterData(api string, systemID string, token string) []byte {
	client := request.Client{
		URL:    api + "/s/" + systemID + "/fronters",
		Method: "GET",
		Header: map[string]string{"Authorization": token},
	}
	resp, err := client.Do()
	if err != nil {
		panic(err)
	}

	return resp.Data
}

func getFront(rawJSON []byte) (string, []structs.Member) {

	var frontWithTimestamp structs.Front
	if err := json.Unmarshal(rawJSON, &frontWithTimestamp); err != nil {
		panic(err)
	}

	var front []structs.Member

	front = frontWithTimestamp.Members
	timestamp := frontWithTimestamp.Timestamp

	return timestamp, front
}
