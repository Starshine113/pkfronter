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
		case `current`:
			currentFrontersCmd(api, id)
		case `switch`:
			members := getMemberIDs(api, token, id, flag.Args()[1:])
			commands.Switch(api, token, memberIDsToJSON(members))
			/* commands.Switch(token, id, flag.Args()[1:]) */
		default:
			currentFrontersCmd(api, id)
		}
	} else {
		currentFrontersCmd(api, id)
	}

}

func currentFrontersCmd(api string, id string) {
	fronters := getFront(getCurrentFronterData(api, id))
	commands.PrettyPrintFronters(fronters)
}
