package d

import (
	"strings"

	"github.com/just-install/just-install-updater-go/jiup/rules/c"
)

// Template creates a download link based on substituting {{.Version}}. Leave x64Tmpl empty is only x86.
func Template(x86Tmpl, x64Tmpl string) c.DownloadExtractorFunc {
	return func(version string) (string, *string, error) {
		r := func(i string) string {
			o := i
			o = strings.Replace(o, "{{.Version}}", version, -1)
			o = strings.Replace(o, "{{.VersionU}}", strings.Replace(version, ".", "_", -1), -1)
			o = strings.Replace(o, "{{.VersionD}}", strings.Replace(version, ".", "-", -1), -1)
			o = strings.Replace(o, "{{.Version0}}", strings.Split(version, ".")[0], -1)
			o = strings.Replace(o, "{{.VersionN}}", strings.Replace(version, ".", "", -1), -1)
			return o
		}
		if x64Tmpl != "" {
			x64 := r(x64Tmpl)
			return r(x86Tmpl), &x64, nil
		}
		return r(x86Tmpl), nil, nil
	}
}
