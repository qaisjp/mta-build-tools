package rescheck

import (
	"bufio"
	"io"
	"strings"

	"github.com/pkg/errors"
)

// VersionItem matches SVersionItem in mtasa-blue
// See https://github.com/multitheftauto/mtasa-blue/blob/2013c1fcdd2f9e5748bc1608fc6285c68e6a8991/Server/mods/deathmatch/logic/CResourceChecker.Data.h#L18
type VersionItem struct {
	functionName  string
	minMtaVersion string
}

// ReadInitLists reads a file in the manner of Server/mods/deathmatch/logic/CResourceChecker.Data.h and spews out the appropriate init lists
func ReadInitLists(r io.Reader) (clientItems []VersionItem, serverItems []VersionItem, err error) {
	// Recover in case of index error
	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("Parse error: %+v", r)
		}
	}()

	s := bufio.NewScanner(r)

	// Init lists
	var currList *[]VersionItem = nil
	clientItems = []VersionItem{}
	serverItems = []VersionItem{}

	for s.Scan() {
		text := s.Text()
		trimText := strings.TrimSpace(text)
		shouldContinue := true

		switch trimText {
		case "SVersionItem clientFunctionInitList[] = {":
			currList = &clientItems
		case "SVersionItem serverFunctionInitList[] = {":
			currList = &serverItems
		case "};":
			currList = nil
		default:
			shouldContinue = currList == nil
		}

		// We continue if we don't have currList, or if we *just* assigned the currList
		if shouldContinue {
			continue
		}

		// Trim a single trailing comma, for normalisation sake
		// Example: `{"toJSON", "1.1.1"},` becomes `{"toJSON", "1.1.1"}`
		n := len(trimText)
		if trimText[n] == ',' {
			trimText = trimText[:n-1]
			n--
		}

		// Error if first and last are not curly braces
		if trimText[0] != '{' && trimText[n-1] != '}' {
			return nil, nil, errors.New("Unexpected line `" + strings.TrimSpace(text) + "`")
		}

		// Trim surrounding curly braces
		// Example: `{"toJSON", "1.1.1"}` becomes `"toJSON", "1.1.1"`
		trimText = trimText[1 : n-1]
		n--

		// Split by comma, and validate length
		fields := strings.Split(trimText, ",")
		if len(fields) != 2 {
			return nil, nil, errors.New("Unexpected line `" + strings.TrimSpace(text) + "`")
		}

		// Trim spaces, ensure encapsulated in quotes, and add to init list
		item := VersionItem{}
		for i, field := range fields {
			field := strings.TrimSpace(field)

			// Ensure quote is valid
			quote := field[0]
			if quote != '\'' && quote != '"' && field[len(field)-1] != quote {
				return nil, nil, errors.New("Badly quoted line `" + strings.TrimSpace(text) + "`")
			}

			// Extract field content
			field = field[1 : len(field)-1]

			// And add to item appropriately
			if i == 0 {
				item.functionName = field
			} else if i == 1 {
				item.minMtaVersion = field
			} else {
				panic("this should never be reached")
			}
		}

		*currList = append(*currList, item)
	}

	if err := s.Err(); err != nil {
		return nil, nil, err
	}

	return nil, nil, nil
}
