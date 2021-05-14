package cmd

import (
	out "github.com/pennz/go-gitlab-client/cli/output"
	"github.com/pennz/go-gitlab-client/gitlab"
)

func printMeta(meta *gitlab.ResponseMeta, withPagination bool) {
	if verbose {
		out.Meta(meta, withPagination)
	}
}
