package redis

import "strings"

type RedisHelp struct {
	Command string
	Args    string
	Desc    string
}

func RedisMatchedCommands(text string) []RedisHelp {
	value := strings.ToUpper(text)
	segments := strings.SplitN(value, " ", 3)

	var matched = make([]RedisHelp, 0)

	for _, help := range commands {
		cmd := strings.SplitN(help.Command, " ", 2)

		if len(cmd) > 1 && len(segments) > 1 {
			if help.Command == segments[0]+" "+segments[1] {
				matched = []RedisHelp{help}
				break
			}

			if strings.HasPrefix(help.Command, segments[0]+" "+segments[1]) {
				matched = append(matched, help)
				continue
			}
		} else {
			if help.Command == segments[0] {
				matched = []RedisHelp{help}
				break
			}

			if strings.HasPrefix(help.Command, segments[0]) {
				matched = append(matched, help)
				continue
			}
		}

	}

	return matched
}

func RedisHelpMatch(text string, matchedFunc func(help RedisHelp)) bool {
	value := strings.ToUpper(text)
	segments := strings.SplitN(value, " ", 3)

	var matchedIndex = -1

	for i, help := range commands {
		cmd := strings.SplitN(help.Command, " ", 2)
		if segments[0] != cmd[0] {
			continue
		}

		if len(cmd) > 1 && len(segments) > 1 && !strings.HasPrefix(cmd[1], segments[1]) {
			continue
		}

		matchedIndex = i
		break
	}

	if matchedIndex > -1 {
		help := commands[matchedIndex]
		matchedFunc(help)
	}

	return matchedIndex > -1
}
