package luadefs_test

import (
	"testing"

	"github.com/multitheftauto/build-tools/internal/luadefs"
	"github.com/stretchr/testify/assert"
)

func TestTraverse(t *testing.T) {
	fs := dummyFS()

	clientDefs, serverDefs, err := luadefs.ReadFuncs(fs)
	assert.NoError(t, err, "must be able to read the filesystem")
	assert.NotNil(t, clientDefs, "clientDefs must exist")
	assert.NotNil(t, serverDefs, "serverDefs must exist")

	clientFuncNames := []string{}
	serverFuncNames := []string{}

	assert.ElementsMatch(t, dummyClientFuncNames, clientFuncNames)
	assert.ElementsMatch(t, dummyServerFuncNames, serverFuncNames)
}
