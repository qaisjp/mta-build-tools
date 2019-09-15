package luadefs_test

import (
	"fmt"
	"testing"

	"github.com/multitheftauto/build-tools/internal/luadefs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTraverse(t *testing.T) {
	fs := dummyFS()

	defs, err := luadefs.ReadFuncs(fs)
	require.NoError(t, err, "must be able to read the filesystem")
	require.NotNil(t, defs, "defs must exist")

	clientFuncNames := []string{}
	serverFuncNames := []string{}

	for _, d := range defs {
		fmt.Println("ok")
		if d.OnClient() {
			fmt.Println("c", d)
			clientFuncNames = append(clientFuncNames, d.FunctionName)
		}
		if d.OnServer() {
			fmt.Println("s", d)
			serverFuncNames = append(serverFuncNames, d.FunctionName)
		}
	}

	assert.ElementsMatch(t, dummyClientFuncNames, clientFuncNames, "client funcs should match")
	assert.ElementsMatch(t, dummyServerFuncNames, serverFuncNames, "server funcs should match")
}
