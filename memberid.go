package main

/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

import (
	"encoding/json"
	"strings"

	"github.com/Starshine113/pkfronter/structs"
	"github.com/monaco-io/request"
)

func getMemberIDs(api string, token string, systemID string, memberInput []string) []string {
	client := request.Client{
		URL:    api + "/s/" + systemID + "/members",
		Method: "GET",
		Header: map[string]string{"Authorization": token},
	}

	resp, err := client.Do()
	if err != nil {
		panic(err)
	}

	var members []structs.Member

	if err := json.Unmarshal(resp.Data, &members); err != nil {
		panic(err)
	}

	memberMap := makeMemberMap(members)
	memberIDs := make([]string, 0)

	for _, memberName := range memberInput {
		if memberMap[strings.ToLower(memberName)] != "" {

		}
		if memberID, ok := memberMap[memberName]; ok {
			memberIDs = append(memberIDs, memberID)
		}
	}

	return memberIDs
}

func makeMemberMap(memberStruct []structs.Member) map[string]string {
	memberMap := make(map[string]string)

	for _, member := range memberStruct {
		memberMap[strings.ToLower(member.Name)] = member.ID
	}

	return memberMap
}

func memberIDsToJSON(members []string) []byte {
	switches := "{\"members\": ["
	for i, member := range members {
		switches += "\"" + member + "\""
		if i != len(members)-1 {
			switches += ", "
		}
	}
	switches += "]}"
	return []byte(switches)
}
