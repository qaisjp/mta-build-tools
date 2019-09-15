// Package luadefs traverses a filesystem (just a repo) and provides information about defined Lua functions
package luadefs

import (
	"github.com/pkg/errors"
	"gopkg.in/src-d/go-billy.v4"
)

func ReadFuncs(fs *billy.Filesystem) (clientDefs []LuaDef, serverDefs []LuaDef, err error) {
	// Recover in case of indexing error
	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("Parse error: %+v", r)
		}
	}()

	return nil, nil, nil
}
