package main

/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

import (
	"fmt"
	"os"
)

const apiVersion string = "1"

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	api, id, token := readConfig(home + "/.pksystem.toml")

	fronter := getCurrentFronterData(api, id)
	//fmt.Println(string(fronter))
	//fmt.Printf("ID: %s\nToken: %s\n\n\n%s\n", id, token, fronter)
	getFront(fronter)
	getSystemID(api, token)
}
