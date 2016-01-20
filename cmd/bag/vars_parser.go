package bag

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

var (
	errEmpty  = errors.New("reader was empty")
	errNoDeps = errors.New("no dependencies prefix found")
)

type variable struct {
	name string
}

func parseVars(reader io.Reader) ([]variable, error) {
	scanner := bufio.NewScanner(reader)
	if !scanner.Scan() {
		return nil, errEmpty
	}

	if scanner.Text() != lineCommented("DEPENDENCIES") {
		return nil, errNoDeps
	}

	var ret []variable
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, varListPrefix) {
			continue
		}
		remainder := line[len(varListPrefix):]
		ret = append(ret, variable{name: remainder})
	}
	return ret, nil
}
