package version

import (
	"encoding/json"
	"fmt"
	"github.com/gosuri/uitable"
	"runtime"
)

var (
	// BinVersion is binary version.
	BinVersion = ""
	// GitVersion is semantic version
	GitVersion = ""
	// BuildDate is ISO8601 format, output of $(date +"%Y-%m-%dT%H:%M:%SZ").
	BuildDate = "1970-01-01T00:00:00Z"
	// GitCommit sha1 from git, output of $(git rev-parse HEAD).
	GitCommit = "$Format:%H$"
)

type Info struct {
	Version    string `json:"version" yaml:"version"`
	GitVersion string `json:"gitVersion" yaml:"gitVersion"`
	GitCommit  string `json:"gitCommit" yaml:"gitCommit"`
	BuildDate  string `json:"buildDate" yaml:"buildDate"`
	GoVersion  string `json:"goVersion" yaml:"goVersion"`
	Compiler   string `json:"compiler" yaml:"compiler"`
	Platform   string `json:"platform" yaml:"platform"`
}

func Get() Info {
	return Info{
		Version:    BinVersion,
		GitVersion: GitVersion,
		GitCommit:  GitCommit,
		BuildDate:  BuildDate,
		GoVersion:  runtime.Version(),
		Compiler:   runtime.Compiler,
		Platform:   fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

func (info Info) ToString() string {
	if s, err := info.Text(); err != nil {
		return string(s)
	}

	return info.GitVersion
}

func (info Info) ToJson() string {
	s, _ := json.Marshal(info)
	return string(s)
}

func (info Info) Text() ([]byte, error) {
	table := uitable.New()
	table.RightAlign(0)
	table.MaxColWidth = 80
	table.Separator = " "
	table.AddRow("version", info.Version)
	table.AddRow("gitVersion", info.GitVersion)
	table.AddRow("gitCommit", info.GitCommit)
	table.AddRow("buildDate", info.BuildDate)
	table.AddRow("goVersion", info.GoVersion)
	table.AddRow("compiler", info.Compiler)
	table.AddRow("platform", info.Platform)

	return table.Bytes(), nil
}
