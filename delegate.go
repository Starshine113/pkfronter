package main

import (
	"flag"

	"github.com/Starshine113/pkfronter/commands"
)

/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

func delegate(api string, id string, token string, cacheFile string) {
	flag.Parse()

	if len(flag.Args()) != 0 {
		switch flag.Args()[0] {
		case `s`, `switch`:
			members := getMemberIDs(api, token, id, flag.Args()[1:])
			commands.Switch(api, token, memberIDsToJSON(members))
			currentFrontersCmd(api, id, token)
		case `current`:
			currentFrontersCmd(api, id, token)
		default:
			currentFrontersCmd(api, id, token)
		}
	} else {
		currentFrontersCmd(api, id, token)
	}

}

func currentFrontersCmd(api string, id string, token string) {
	timestamp, fronters := getFront(getCurrentFronterData(api, id, token))
	commands.PrettyPrintFronters(timestamp, fronters)
}
