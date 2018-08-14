package rule

import (
	"strings"

	"github.com/z0mbie42/flint/lint"
)

type NoTrailingUnderscores struct{}

func (r NoTrailingUnderscores) Apply(file lint.File) []lint.Issue {
	parts := strings.Split(strings.TrimSuffix(file.Name, file.Ext), ".")
	issues := []lint.Issue{}

	for _, part := range parts {
		if strings.TrimLeft(part, "_") != part {
			issue := lint.Issue{
				File:         file,
				RuleName:     r.Name(),
				Explaination: "Unexpected leading '_'",
			}
			issues = append(issues, issue)
		}
	}

	return issues
}

func (_ NoTrailingUnderscores) Name() string {
	return "no_trailing_underscores"
}