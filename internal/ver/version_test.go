package ver_test

import (
	"testing"

	"github.com/multitheftauto/build-tools/internal/ver"
	"github.com/stretchr/testify/assert"
)

func performTest(t *testing.T, verstr string, v ver.MtaVersion) {
	assert.Equal(t, ver.MtaVersionFromString(verstr), v)
}

func TestFull(t *testing.T) {
	performTest(t, "1.3.2-9.03931", ver.MtaVersion{1, 3, 2, 3931})
}

func TestToPatch(t *testing.T) {
	performTest(t, "1.3.1", ver.MtaVersion{1, 3, 1, 0})
}

func TestToMinor(t *testing.T) {
	performTest(t, "1.3", ver.MtaVersion{1, 3, 0, 0})
}
