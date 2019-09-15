package ver

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

var versionDefines = []string{
	"MTASA_VERSION_MAJOR", "MTASA_VERSION_MINOR", "MTASA_VERSION_MAINTENANCE",
}

// ReadVersionHeader reads a file in the form of version.h and extracts the MtaVersion struct
func ReadVersionHeader(r io.Reader) (v MtaVersion, err error) {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		text := scanner.Text()
		for _, def := range versionDefines {
			prefix := "#define " + def
			if strings.HasPrefix(text, prefix) {
				text = strings.TrimSpace(strings.TrimPrefix(text, prefix))
				n, err := strconv.Atoi(text)
				if err != nil {
					return v, err
				}

				switch def {
				case "MTASA_VERSION_MAJOR":
					v.Major = n
				case "MTASA_VERSION_MINOR":
					v.Minor = n
				case "MTASA_VERSION_MAINTENANCE":
					v.Patch = n
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return v, err
	}

	return v, nil
}
