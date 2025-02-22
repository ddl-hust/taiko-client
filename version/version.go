package version

// Version info.
const Version = "0.17.0" // x-release-please-version

var Meta = "dev"

// Git commit/date info, set via linker flags.
var (
	GitCommit = ""
	GitDate   = ""
)

// VersionWithCommit returns a textual version string including Git commit/date
// information.
func VersionWithCommit() string {
	vsn := Version + "-" + Meta
	if len(GitCommit) >= 8 {
		vsn += "-" + GitCommit[:8]
	}
	if (Meta != "stable") && (GitDate != "") {
		vsn += "-" + GitDate
	}
	return vsn
}
