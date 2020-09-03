package commands

/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

import (
	"fmt"
	"time"

	"github.com/Starshine113/pkfronter/structs"
)

// PrettyPrintFronters prints prins a comma-separated list of all people currently fronting.
func PrettyPrintFronters(timestamp string, inputData []structs.Member, systemName string) {
	fmt.Printf("Current fronter(s) of %s:\n", systemName)
	for num, member := range inputData {
		if len(inputData) == 1 {
			fmt.Printf("%s", member.Name)
		} else {
			if num == len(inputData)-1 {
				fmt.Printf("and %s", member.Name)
			} else {
				fmt.Printf("%s, ", member.Name)
			}
		}
	}
	switchTime, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nSince: %s\n", switchTime.Format(time.RFC1123))
}
