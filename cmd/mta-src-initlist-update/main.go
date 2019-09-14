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
	"log"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/src-d/go-git.v4"
)

var srcFolder = flag.String("src", "", "Path to mtasa-blue git repository")
var remoteRepo = flag.String("remote", "git@github.com:qaisjp/mtasa-blue", "Remote to push + send a pull request to")
var prevRef = flag.String("prevref", "HEAD^", "the previous commit, which should be a reference to an object (tag, commit hash, branch, etc)")

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

	ref, err := r.Head()
	if err != nil {
		log.Fatalln("Couldn't get HEAD reference")
	}

	log.Println("It's", ref.String())
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
