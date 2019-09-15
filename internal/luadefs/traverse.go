// Package luadefs traverses a filesystem (just a repo) and provides information about defined Lua functions
package luadefs

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/src-d/go-billy.v4"
)

// ReadFuncs traverses the filesystem and looks for lua definitions in the obvious places, returning all the definitions
func ReadFuncs(fs billy.Filesystem) (defs []LuaDef, err error) {
	// Recover in case of indexing error
	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("Parse error: %+v", r)
		}
	}()

	defs = []LuaDef{}

	// value true = is client, value false = is server
	paths := map[string]bool{
		"Server/mods/deathmatch/logic/luadefs":             false,
		"Server/mods/deathmatch/logic/lua/CLuaManager.cpp": false,
		"Client/mods/deathmatch/logic/luadefs":             true,
		"Client/mods/deathmatch/logic/lua/CLuaManager.cpp": true,
		// todo(qaisjp): Shared defs!!!!
	}

	// populate file lists
	fileinfos := map[bool][]os.FileInfo{}

	for filepath, csided := range paths {
		if path.Ext(filepath) == ".cpp" {
			file, err := fs.Stat(filepath)
			if err != nil {
				return nil, errors.Wrap(err, "Could not open file: "+filepath)
			}
			fileinfos[csided] = append(fileinfos[csided], file)
		} else {
			files, err := fs.ReadDir(filepath)
			if err != nil {
				return nil, errors.Wrap(err, "could not read luadefs folder")
			}

			for _, file := range files {
				filename := file.Name()
				if !file.IsDir() && strings.HasPrefix(filename, "CLua") && strings.HasSuffix(filename, "Defs.cpp") {
					// fileclass := filename[4:len(filename)-8]
					fileinfos[csided] = append(fileinfos[csided], file)
				}
			}
		}
	}

	for b, fs := range fileinfos {
		for _, f := range fs {
			fmt.Printf("%+v, %v\n", b, f)
		}
	}

	return nil, nil
}
