package ver

import (
	"strconv"
	"strings"
)

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
