package cli

import (
	"strings"
)

func SplitInput(input string) (head string, tail []string) {
	input = strings.TrimSpace(input)
	tokens := strings.Fields(input)
	if len(tokens) == 0 {
		return "", []string{}
	}

	head = tokens[0]

	if len(tokens) > 1 {
		tail = tokens[1:]
	} else {
		tail = []string{}
	}

	return head, tail
}
