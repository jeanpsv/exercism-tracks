package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	regex, _ := regexp.Compile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return regex.MatchString(text)
}

func SplitLogLine(text string) []string {
	regex, _ := regexp.Compile(`<[~*=-]*>`)
	return regex.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var count int
	// regex, _ := regexp.Compile(`(?i)"*password*"`)
	regex, _ := regexp.Compile(`(?i)".*password.*"`)
	for _, line := range lines {
		if regex.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	regex, _ := regexp.Compile(`end-of-line\d*`)
	return regex.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	regex, _ := regexp.Compile(`User\s+(\w+)`)
	var logs []string

	for _, line := range lines {
		founds := regex.FindStringSubmatch(line)
		if founds != nil {
			line = fmt.Sprintf("[USR] %s %s", founds[1], line)
		}
		logs = append(logs, line)
	}
	return logs
}
