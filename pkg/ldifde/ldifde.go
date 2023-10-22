package ldifde

import (
	"changeme/pkg/utils"
	"regexp"
	"strings"
)

var LdifDataItems []string

func Preprocessing(path string) (err error) {
	res, err := utils.ReadFile(path)
	if err != nil {
		return
	}

	fileData := string(res)
	fileData = strings.ReplaceAll(fileData, "\r\n", "\n")
	fileData = strings.ReplaceAll(fileData, "\n ", " ")
	LdifDataItems = strings.Split(fileData, "\n\n")
	return
}

func GetSamAccountNameWithDn(dn string) (samAccountName string) {
	var reSamAccountName = regexp.MustCompile(`(?i)(?m)^sAMAccountName: `)
	var reName = regexp.MustCompile(`(?i)(?m)^name: `)

	for _, item := range LdifDataItems {
		oneLine := strings.ToLower(strings.Split(item, "\n")[0])
		if oneLine == "dn: "+strings.ToLower(dn) {
			if reSamAccountName.MatchString(item) {
				var re = regexp.MustCompile(`(?i)(?m)^sAMAccountName: (.*)\n`)
				match := re.FindStringSubmatch(item)
				if len(match) == 2 {
					samAccountName = match[1]
				}
			} else if reName.MatchString(item) {
				var re = regexp.MustCompile(`(?i)(?m)^name: (.*)\n`)
				match := re.FindStringSubmatch(item)
				if len(match) == 2 {
					samAccountName = match[1]
				}
			}
		}
	}
	return
}
