package ldifde

import (
	"encoding/json"
	"regexp"
	"slices"
	"sort"
	"strings"
)

var Users []string

type UserGraph struct {
	Name     string      `json:"name"`
	Value    int         `json:"value"`
	Children []UserGraph `json:"children"`
}

func GetUsers() (users []string) {
	for _, item := range LdifDataItems {
		if strings.Contains(item, "\nobjectClass: user\n") {
			name := getUserName(item)
			if name != "" {
				users = append(users, name)
			}
		}
	}
	sort.Strings(users)
	Users = users
	return
}

func getUserName(item string) (name string) {
	var re = regexp.MustCompile(`(?m)^sAMAccountName: (.*)\n`)
	res := re.FindStringSubmatch(item)
	if len(res) == 2 {
		name = res[1]
	}
	return
}

func GetUserGraph(userName string) (data string) {
	userGraph := []UserGraph{}
	userGraph = append(userGraph, UserGraph{
		Name:  userName,
		Value: 1,
	})

	getMemberOf(&userGraph)
	dataBytes, err := json.Marshal(userGraph)
	if err != nil {
		return
	}
	data = string(dataBytes)
	return
}

func getMemberOf(userGraph *[]UserGraph) {
	for i, item := range *userGraph {
		samAccount := item.Name
		memberOf := getUserMemberOf(samAccount)
		for _, item := range memberOf {
			value := 1
			if slices.Contains(Users, item) {
				value = 2
			}
			(*userGraph)[i].Children = append((*userGraph)[i].Children, UserGraph{Name: item, Value: value})
		}
		getMemberOf(&(*userGraph)[i].Children)
	}
}

func getUserMemberOf(userName string) (memberOf []string) {
	var reMemberOf = regexp.MustCompile(`(?i)(?m)^memberOf: (CN=.*?)\n`)
	var reSamAccountName = regexp.MustCompile(`(?i)(?m)^sAMAccountName: ` + userName + "\n")

	for _, item := range LdifDataItems {
		if reSamAccountName.MatchString(item) {
			for _, match := range reMemberOf.FindAllStringSubmatch(item, -1) {
				dn := match[1]
				samAccountName := GetSamAccountNameWithDn(dn)
				memberOf = append(memberOf, samAccountName)
			}
		}
	}
	return
}
