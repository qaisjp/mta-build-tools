package ver

import (
	"fmt"
	"strconv"
	"strings"
)

// MtaVersion is a very primitive representation of an instance of an MTA version.
//
// - It assumes that the build type is "9" (release).
// - It ignores "build revision" (this is not the same as r9000 revisions)
type MtaVersion struct {
	Major       int
	Minor       int
	Patch       int
	BuildNumber int
}

func (v MtaVersion) String() string {
	if v.Major > 9 || v.Minor > 9 || v.Patch > 9 {
		panic("version major, minor or patch are too large to represent as a sortable string")
	}
	return fmt.Sprintf("%d.%d.%d-9.%05d", v.Major, v.Minor, v.Patch, v.BuildNumber)
}

// MtaVersionFromString parses MTA sortable version strings
func MtaVersionFromString(str string) (v MtaVersion) {
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
