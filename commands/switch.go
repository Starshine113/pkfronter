package commands

/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

import (
	"flag"
	"fmt"

	"github.com/monaco-io/request"
)

// Switch switches the IDs in `members` in.
func Switch(api string, token string, members []byte) {
	client := request.Client{
		URL:    api + "/s/switches",
		Method: "POST",
		Header: map[string]string{"Authorization": token},
		Body:   members,
	}
	resp, _ := client.Do()

	prettyPrintSwitchOutput(resp.Code, resp.Data)
}

func prettyPrintSwitchOutput(code int, resp []byte) {
	switch code {
	case 400:
		fmt.Printf("❌ ")
		if len(flag.Args()[1:]) == 0 {
			fmt.Printf("Already switched out.\n")
		} else if len(flag.Args()[1:]) != 1 {
			fmt.Printf("Members ")
			for i, member := range flag.Args()[1:] {
				if i == len(flag.Args()[1:])-1 {
					fmt.Printf("%s ", member)
				} else {
					fmt.Printf("%s, ", member)
				}
			}
			fmt.Printf("are already fronting.\n")
		} else {
			fmt.Printf("Member %s is already fronting.\n", flag.Args()[1])
		}
	case 204:
		if len(flag.Args()[1:]) == 0 {
			fmt.Printf("✅ Switch-out registered.\n")
		} else {
			fmt.Printf("✅ Switch registered.\n")
		}
	default:
		fmt.Printf("❌ %s\n", string(resp))
	}
}
