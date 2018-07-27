package command

// Default build-time variable.
// These values are overridden via ldflags
var (
	Version     = "1.0.0"
	GitCommit   string
	GitDescribe string
	BuildTime   string
)
