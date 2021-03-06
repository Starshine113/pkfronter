package structs

/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

type Member struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Displayname string `json:"display_name"`
}

type Front struct {
	Timestamp string   `json:"timestamp"`
	Members   []Member `json:"members"`
}
