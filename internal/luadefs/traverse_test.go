package luadefs_test

import (
	"testing"

	"github.com/multitheftauto/build-tools/internal/luadefs"
	"github.com/stretchr/testify/assert"
)

func TestTraverse(t *testing.T) {
	fs := dummyFS()

	defs, err := luadefs.ReadFuncs(fs)
	assert.NoError(t, err, "must be able to read the filesystem")
	assert.NotNil(t, defs, "defs must exist")

	clientFuncNames := []string{}
	serverFuncNames := []string{}

	for _, d := range defs {
		if d.OnClient() {
			clientFuncNames = append(clientFuncNames, d.FunctionName)
		}
		if d.OnServer() {
			clientFuncNames = append(serverFuncNames, d.FunctionName)
		}
	}

	assert.ElementsMatch(t, dummyClientFuncNames, clientFuncNames)
	assert.ElementsMatch(t, dummyServerFuncNames, serverFuncNames)
}
