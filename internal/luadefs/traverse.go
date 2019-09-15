// Package luadefs traverses a filesystem (just a repo) and provides information about defined Lua functions
package luadefs

import (
	"bufio"
	"io"
	"path"
	"strings"

	"github.com/multitheftauto/build-tools/internal"

	"github.com/pkg/errors"
	"gopkg.in/src-d/go-billy.v4"
)

// ReadFuncs traverses the filesystem and looks for lua definitions in the obvious places, returning all the definitions
func ReadFuncs(fs billy.Filesystem) (defs []LuaDef, err error) {
	paths := map[string]FunctionType{
		"Server/mods/deathmatch/logic/luadefs":                 ServerFunctionType,
		"Server/mods/deathmatch/logic/lua/CLuaManager.cpp":     ServerFunctionType,
		"Client/mods/deathmatch/logic/luadefs":                 ClientFunctionType,
		"Client/mods/deathmatch/logic/lua/CLuaManager.cpp":     ClientFunctionType,
		"Shared/mods/deathmatch/logic/luadefs/CLuaBitDefs.cpp": SharedFunctionType,
	}

	// populate file lists
	type fileEntry struct {
		fpath string
		ftype FunctionType
	}
	fileEntries := []fileEntry{}

	for filepath, ftype := range paths {
		if path.Ext(filepath) == ".cpp" {
			file, err := fs.Stat(filepath)
			if err != nil {
				return nil, errors.Wrap(err, "Could not open file: "+filepath)
			} else if file.IsDir() {
				return nil, errors.Wrap(err, "Expected file but got directory for: "+filepath)
			}
			fileEntries = append(fileEntries, fileEntry{filepath, ftype})
		} else {
			files, err := fs.ReadDir(filepath)
			if err != nil {
				return nil, errors.Wrap(err, "could not read luadefs folder")
			}

			for _, file := range files {
				filename := file.Name()
				if !file.IsDir() && strings.HasPrefix(filename, "CLua") && strings.HasSuffix(filename, "Defs.cpp") {
					// fileclass := filename[4:len(filename)-8]
					fileEntries = append(fileEntries, fileEntry{path.Join(filepath, file.Name()), ftype})
				}
			}
		}
	}

	defs = []LuaDef{}
	for _, entry := range fileEntries {
		f, err := fs.Open(entry.fpath)
		if err != nil {
			return nil, errors.Wrap(err, "Could not open file: "+entry.fpath)
		}

		entrydefs, err := FindDefs(f, entry.ftype)
		if err != nil {
			return nil, errors.Wrap(err, "trouble reading "+entry.fpath)
		}

		for _, def := range entrydefs {
			defs = append(defs, def)
		}
	}

	return defs, nil
}

// FindDefs reads a file and spits out the associated function definitions
func FindDefs(r io.Reader, ftype FunctionType) (defs []LuaDef, err error) {
	// Recover in case of indexing error
	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("Parse error: %+v", r)
		}
	}()

	// Initialise empty array

	s := bufio.NewScanner(r)
	commentFound := false
	for s.Scan() {
		text := strings.TrimSpace(s.Text())

		if strings.HasSuffix(text, "*/") {
			commentFound = false
			continue
		} else if strings.HasPrefix(text, "/*") || commentFound {
			commentFound = true
			continue
		}

		if text == "std::map<const char*, lua_CFunction> functions{" {
			defs = []LuaDef{}
		} else if text == "};" {
			break
		} else if defs != nil {
			pair, err := internal.ExtractPair(text, false)
			if err != nil {
				return nil, err
			} else if pair != nil {
				defs = append(defs, LuaDef{pair[0], ftype})
			}
		}
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return defs, nil
}
