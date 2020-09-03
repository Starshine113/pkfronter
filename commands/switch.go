package commands

/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

import (
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
	resp, err := client.Do()

	fmt.Println(resp.Code, string(resp.Data), err)
}
