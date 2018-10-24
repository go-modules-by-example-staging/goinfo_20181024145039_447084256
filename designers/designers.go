package designers

import "github.com/go-modules-by-example-staging/goinfo/contributors"

func Names() []string {
	var res []string
	for _, p := range contributors.Details() {
		switch p.FullName {
		case "Rob Pike", "Ken Thompson", "Robert Griesemer":
			res = append(res, p.FullName)
		}
	}
	return res
}
