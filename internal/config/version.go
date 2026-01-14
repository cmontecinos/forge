package config

// Version information - can be set at build time via ldflags:
// go build -ldflags "-X github.com/bigbytes/forge/internal/config.Version=1.0.0 \
//   -X github.com/bigbytes/forge/internal/config.BuildDate=$(date -u +%Y-%m-%dT%H:%M:%SZ) \
//   -X github.com/bigbytes/forge/internal/config.GitCommit=$(git rev-parse --short HEAD)"

var (
	// Version is the semantic version of forge
	Version = "0.1.0"

	// BuildDate is when the binary was built
	BuildDate = "unknown"

	// GitCommit is the git commit hash
	GitCommit = "unknown"
)
