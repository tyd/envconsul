package version

import "fmt"

const Version = "0.7.4a (tyd)"

var (
	Name      string
	GitCommit string

	HumanVersion = fmt.Sprintf("%s v%s (%s)", Name, Version, GitCommit)
)
