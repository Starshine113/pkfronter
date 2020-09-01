package main

/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

import (
	"encoding/json"
	"io/ioutil"

	"github.com/monaco-io/request"
	"github.com/pelletier/go-toml"
)

func readConfig(configPath string) (string, string, string) {
	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	configString := string(configFile)
	config, err := toml.Load(configString)
	if err != nil {
		panic(err)
	}
	token := config.Get("token").(string)
	api := "https://api.pluralkit.me/v" + apiVersion
	systemID := getSystemID(api, token)

	return api, systemID, token
}

func getSystemID(api string, token string) string {
	client := request.Client{
		URL:    api + "/s",
		Method: "GET",
		Header: map[string]string{"Authorization": token},
	}
	resp, err := client.Do()
	if err != nil {
		panic(err)
	}

	var data map[string]interface{}

	if err := json.Unmarshal(resp.Data, &data); err != nil {
		panic(err)
	}

	return string(data["id"].(string))
}
