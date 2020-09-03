package main

/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

// Current PluralKit API version
const apiVersion string = "1"

func main() {
	api, id, token, cacheFile, systemName := readConfig()

	delegate(api, id, token, cacheFile, systemName)
}
