package ldifde

import (
	"encoding/json"
	"regexp"
	"slices"
	"sort"
	"strings"
)

var Groups []string

type GroupGraph struct {
	Name     string       `json:"name"`
	Value    int          `json:"value"`
	Children []GroupGraph `json:"children"`
}

func GetGroups() (groups []string) {
	for _, item := range LdifDataItems {
		if strings.Contains(item, "\nobjectClass: group\n") {
			name := getGroupName(item)
			if name != "" {
				groups = append(groups, name)
			}
		}
	}
	sort.Strings(groups)
	Groups = groups
	return
}

func getGroupName(item string) (name string) {
	var re = regexp.MustCompile(`(?m)^sAMAccountName: (.*)\n`)
	res := re.FindStringSubmatch(item)
	if len(res) == 2 {
		name = res[1]
	}
	return
}

func GetGroupGraph(groupName string) (data string) {
	groupGraph := []GroupGraph{}
	groupGraph = append(groupGraph, GroupGraph{
		Name:  groupName,
		Value: 1,
	})

	getMembers(&groupGraph)
	dataBytes, err := json.Marshal(groupGraph)
	if err != nil {
		return
	}
	data = string(dataBytes)
	return
}

func getMembers(groupGraph *[]GroupGraph) {
	for i, item := range *groupGraph {
		samAccount := item.Name
		if slices.Contains(Groups, samAccount) {
			members := getGroupMembers(samAccount)
			for _, member := range members {
				value := 2
				if slices.Contains(Groups, member) {
					value = 1
				}
				(*groupGraph)[i].Children = append((*groupGraph)[i].Children, GroupGraph{Name: member, Value: value})
			}
			getMembers(&(*groupGraph)[i].Children)
		}
	}
}

func getGroupMembers(groupName string) (members []string) {
	var reMember = regexp.MustCompile(`(?i)(?m)^member: (CN=.*?)\n`)
	var reSamAccountName = regexp.MustCompile(`(?i)(?m)^sAMAccountName: ` + groupName + "\n")

	for _, item := range LdifDataItems {
		if reSamAccountName.MatchString(item) {
			for _, match := range reMember.FindAllStringSubmatch(item, -1) {
				dn := match[1]
				samAccountName := GetSamAccountNameWithDn(dn)
				members = append(members, samAccountName)
			}
		}
	}
	return
}
