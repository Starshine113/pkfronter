package commands

/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

import (
	"fmt"

	"github.com/Starshine113/pkfronter/structs"
)

// PrettyPrintFronters prints prins a comma-separated list of all people currently fronting.
func PrettyPrintFronters(inputData []structs.Member) {
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
	fmt.Printf("\n")
}
