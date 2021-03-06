package version

import (
	"github.com/coreos/fleet/third_party/github.com/coreos/go-semver/semver"
)

const Version = "0.4.0+git"

var SemVersion semver.Version

func init() {
	sv, err := semver.NewVersion(Version)
	if err != nil {
		panic("bad version string!")
	}
	SemVersion = *sv
}
