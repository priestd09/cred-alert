package git

import (
	"errors"
	"regexp"
	"strconv"
)

type Hunk struct {
	path      string
	startLine int
	length    int
}

func newHunk(path string, startLine int, length int) *Hunk {
	hunk := new(Hunk)
	hunk.path = path
	hunk.startLine = startLine
	hunk.length = length

	return hunk
}

func hunkHeader(rawLine string) (int, int, error) {
	matches := regexp.MustCompile(`^@@.*\+(\d+),?(\d+)?\s@@`).FindStringSubmatch(rawLine)
	if len(matches) < 3 {
		return 0, 0, errors.New("Not a hunk header")
	}

	startLine, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, 0, err
	}

	var length int
	if matches[2] == "" {
		length = 1
	} else {
		var err error
		if length, err = strconv.Atoi(matches[2]); err != nil {
			return 0, 0, err
		}
	}

	return startLine, length, nil
}

func (h Hunk) endOfHunk(lineNumber int) bool {
	return lineNumber >= h.startLine+h.length
}