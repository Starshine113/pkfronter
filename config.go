package main

/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/monaco-io/request"
	"github.com/pelletier/go-toml"
)

func readConfig() (string, string, string, string) {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	configFile, err := ioutil.ReadFile(home + "/.pksystem.toml")
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
	cacheFile := home + "/" + config.Get("token").(string)

	return api, systemID, token, cacheFile
}

func getSystemID(api string, token string) string {
	var data map[string]interface{}

	client := request.Client{
		URL:    api + "/s",
		Method: "GET",
		Header: map[string]string{"Authorization": token},
	}

	resp, err := client.Do()
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(resp.Data, &data); err != nil {
		panic(err)
	}

	return string(data["id"].(string))
}
