package ver

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var versionPrefixes = []string{
	"MTASA_VERSION_MAJOR", "MTASA_VERSION_MINOR", "MTASA_VERSION_MAINTENANCE",
}

func ReadVersionHeader(r io.Reader) (v MtaVersion, err error) {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		text := scanner.Text()
		for _, prefix := range versionPrefixes {
			if strings.HasPrefix(text, "#define "+prefix) {
				fmt.Println(text)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return v, err
	}
	return v, nil
}
