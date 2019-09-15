// mta-src-initlist-update is a tool that ensures the initlist is up to date
//
// It:
// - ensures a function exists for all items mentioned in the init list (removes nonsensical items)
// - only adds functions introduced through a changeset (it will not try to add stuff from previous majors)
// - knows the current version, and so removes stuff from previous majors (removes redundant items)
//
// Finally, if certain parameters are provided, it will also:
// - generate a patch
// - create a pull request to apply the patch
//
// Improvements:
// - this scrapes from buildinfo, but if this is included in the build pipeline, it can do a better job
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/multitheftauto/build-tools/internal/rescheck"
	"github.com/multitheftauto/build-tools/internal/ver"
	"github.com/pkg/errors"
	"gopkg.in/src-d/go-git.v4"
)

var srcFolder = flag.String("src", "", "Path to mtasa-blue git repository")
var remoteRepo = flag.String("remote", "git@github.com:qaisjp/mtasa-blue", "Remote to push + send a pull request to")
var prevRef = flag.String("prevref", "HEAD^", "the previous commit, which should be a reference to an object (tag, commit hash, branch, etc)")

// Tool is the tool
type Tool struct {
	r   *git.Repository
	src string
}

func main() {
	flag.Parse()

	path := *srcFolder
	if err := checkSourceFolder(path); err != nil {
		log.Fatalln("Invalid `src` provided:", path, err.Error())
		return
	}

	r, err := git.PlainOpen(path)
	if err != nil {
		log.Fatalln("Could not open git repo:", err.Error())
	}

	t := Tool{r, path}

	if err := t.run(); err != nil {
		log.Fatalln(err.Error())
	}
}

func checkSourceFolder(path string) error {
	// Get path info
	finfo, err := os.Stat(path)
	if err != nil {
		return errors.Wrap(err, "stat path failed")
	}

	// Ensure is directory
	if !finfo.IsDir() {
		return errors.New("expected directory, got file")
	}

	return nil
}

func (t *Tool) run() error {
	ref, err := t.r.Head()
	if err != nil {
		log.Fatalln("Couldn't get HEAD reference")
	}

	fmt.Println("repo is at:", ref.String())

	ver, err := t.getCurrentVersion()
	if err != nil {
		log.Fatalln("Could not get current version:", err.Error())
	}
	fmt.Println("repo version is:", ver)

	clientItems, serverItems, err := t.getInitLists()
	if err != nil {
		log.Fatalln("Could not read CResourceChecker.Data.h init lists:", err.Error())
	}

	// delClientItems, delServerItems, err := t.getInitLists()
	dupes := findDuplicates(clientItems)
	if len(dupes) == 0 {
		fmt.Println("ok(client): no dupes")
	} else {
		fmt.Println("fail(client):  dupes found")
	}

	dupes = findDuplicates(serverItems)
	if len(dupes) == 0 {
		fmt.Println("ok(server): no dupes")
	} else {
		fmt.Println("fail(server):  dupes found")
	}

	return nil
}

func (t *Tool) getCurrentVersion() (v ver.MtaVersion, err error) {
	path := filepath.Join(t.src, "Shared/sdk/version.h")

	f, err := os.Open(path)
	if err != nil {
		return v, err
	}

	return ver.ReadVersionHeader(f)
}

func (t *Tool) getInitLists() (clientItems []rescheck.VersionItem, serverItems []rescheck.VersionItem, err error) {
	path := filepath.Join(t.src, "Server/mods/deathmatch/logic/CResourceChecker.Data.h")

	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	return rescheck.ReadInitLists(f)
}
