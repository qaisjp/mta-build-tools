package internal

import (
	"errors"
	"strings"
)

// ExtractPair will take a line like:
// - {"first", "last"}
// - {"first", Last}
// - // some comment
// - # some pragma
// - <empty line>
//
// And return something like []string{"first", "last"}. Nil if it's comment or empty.
func ExtractPair(line string, secondQuote bool) (pair []string, err error) {
	originalLine := line

	if strings.HasPrefix(line, "//") || strings.HasPrefix(line, "#") || line == "" {
		return nil, nil
	}

	// Trim a single trailing comma, for normalisation sake
	// Example: `{"toJSON", "1.1.1"},` becomes `{"toJSON", "1.1.1"}`
	n := len(line)
	if line[n-1] == ',' {
		line = line[:n-1]
		n--
	}

	// Error if first and last are not curly braces
	if line[0] != '{' && line[n-1] != '}' {
		return nil, errors.New("Unexpected line `" + originalLine + "`")
	}

	// Trim surrounding curly braces
	// Example: `{"toJSON", "1.1.1"}` becomes `"toJSON", "1.1.1"`
	line = line[1 : n-1]
	n--

	// Split by comma, and validate length
	fields := strings.Split(line, ",")
	if len(fields) != 2 {
		return nil, errors.New("Unexpected line `" + originalLine + "`")
	}

	// Trim spaces, ensure encapsulated in quotes, and add to init list
	item := []string{"", ""}
	for i, field := range fields {
		field := strings.TrimSpace(field)

		if i == 0 || (i == 1 && secondQuote) {
			// Ensure quote is valid
			quote := field[0]
			if quote != '\'' && quote != '"' && field[len(field)-1] != quote {
				return nil, errors.New("Badly quoted line `" + originalLine + "`")
			}

			// Extract field content
			field = field[1 : len(field)-1]
		}

		// And add to item appropriately
		if i == 0 {
			item[0] = field
		} else if i == 1 {
			item[1] = field
		} else {
			panic("this should never be reached")
		}
	}

	return item, nil
}
