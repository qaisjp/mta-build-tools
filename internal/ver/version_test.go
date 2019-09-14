package ver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func performTest(t *testing.T, verstr string, v MtaVersion) {
	assert.Equal(t, MtaVersionFromString(verstr), v)
}

func TestFull(t *testing.T) {
	performTest(t, "1.3.2-9.03931", MtaVersion{1, 3, 2, 3931})
}

func TestToPatch(t *testing.T) {
	performTest(t, "1.3.1", MtaVersion{1, 3, 1, 0})
}

func TestToMinor(t *testing.T) {
	performTest(t, "1.3", MtaVersion{1, 3, 0, 0})
}
