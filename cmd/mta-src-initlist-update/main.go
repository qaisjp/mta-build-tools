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
	"strconv"
	"strings"

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

	log.Println("It's", ref.String())

	return nil
}

func (t *Tool) getCurrentVersion() {

}

// This might
type mtaVersion struct {
	Major       int
	Minor       int
	Patch       int
	BuildNumber int
}

func mtaVersionFromString(str string) (v mtaVersion) {
	parts := strings.Split(str, "-")

	ver := parts[0]
	nums := strings.Split(ver, ".")

	var err error
	v.Major, err = strconv.Atoi(nums[0])
	if err != nil {
		// todo: be helpful
		panic(err)
	}

	v.Minor, err = strconv.Atoi(nums[1])
	if err != nil {
		// todo: be helpful
		panic(err)
	}

	if len(nums) == 3 {
		v.Patch, err = strconv.Atoi(nums[2])
		if err != nil {
			// todo: be helpful
			panic(err)
		}
	} else if len(nums) > 3 {
		panic("malformed nums, str: " + str)
	}

	if len(parts) == 2 {
		second := parts[1]
		parts := strings.Split(second, ".")
		if parts[0] != "9" {
			// todo: be helpful
			panic("unexpected build type: " + parts[0] + " in version string " + str)
		}
		v.BuildNumber, err = strconv.Atoi(parts[1])
		if err != nil {
			// todo: be helpful
			panic(err)
		}
	}

	return
}
