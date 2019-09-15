package rescheck

import (
	"bufio"
	"io"
	"strings"

	"github.com/multitheftauto/build-tools/internal"

	"github.com/pkg/errors"
)

// VersionItem matches SVersionItem in mtasa-blue
// See https://github.com/multitheftauto/mtasa-blue/blob/2013c1fcdd2f9e5748bc1608fc6285c68e6a8991/Server/mods/deathmatch/logic/CResourceChecker.Data.h#L18
type VersionItem struct {
	FunctionName  string
	MinMtaVersion string
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

		pair, err := internal.ExtractPair(trimText, true)
		if err != nil {
			return nil, nil, err
		}

		if pair != nil {
			*currList = append(*currList, VersionItem{pair[0], pair[1]})
		}
	}

	if err := s.Err(); err != nil {
		return nil, nil, err
	}

	return
}
