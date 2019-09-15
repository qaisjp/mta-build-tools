package luadefs_test

import (
	"testing"

	"github.com/multitheftauto/build-tools/internal/luadefs"
	"github.com/stretchr/testify/assert"
)

func TestTraverse(t *testing.T) {
	fs := dummyFS()

	clientFuncs, serverFuncs, err := luadefs.ReadFuncs(fs)
	assert.NoError(t, err, "must be able to read the filesystem")
	assert.NotNil(t, clientFuncs, "clientFuncs must exist")
	assert.NotNil(t, serverFuncs, "serverFuncs must exist")

	clientFuncNames := []string{}
	serverFuncNames := []string{}

	assert.ElementsMatch(t, dummyClientFuncNames, clientFuncNames)
	assert.ElementsMatch(t, dummyServerFuncNames, serverFuncNames)
}
