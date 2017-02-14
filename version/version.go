package version

var (
	name        = "axiom"
	version     = "0.1.0"
	description = "metadata service for docker"
	GitCommit   = "HEAD"
)

func Name() string {
	return name
}

func Version() string {
	return version + " (" + GitCommit + ")"
}

func Description() string {
	return description
}

func FullVersion() string {
	return Name() + " " + Version()
}
