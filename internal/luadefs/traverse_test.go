package luadefs_test

import (
	"testing"

	"github.com/multitheftauto/build-tools/internal/luadefs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTraverse(t *testing.T) {
	fs := dummyFS()

	defs, err := luadefs.FindFilesystemDefs(fs)
	require.NoError(t, err, "must be able to read the filesystem")
	require.NotNil(t, defs, "defs must exist")

	clientFuncNames := []string{}
	serverFuncNames := []string{}

	for _, d := range defs {
		if d.OnClient() {
			clientFuncNames = append(clientFuncNames, d.FunctionName)
		}
		if d.OnServer() {
			serverFuncNames = append(serverFuncNames, d.FunctionName)
		}
	}

	assert.ElementsMatch(t, dummyClientFuncNames, clientFuncNames, "client funcs should match")
	assert.ElementsMatch(t, dummyServerFuncNames, serverFuncNames, "server funcs should match")
}
