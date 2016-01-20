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
	name        string
	description string
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
		spl := strings.Split(remainder, ": ")
		newVar := variable{}
		newVar.name = spl[0]
		if len(spl) > 1 {
			newVar.description = spl[1]
		}
		ret = append(ret, newVar)
	}
	return ret, nil
}
