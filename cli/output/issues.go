package output

import (
	"fmt"
	"io"

	"github.com/olekukonko/tablewriter"
	"github.com/pennz/go-gitlab-client/gitlab"
)

func Issues(w io.Writer, format string, collection *gitlab.IssueCollection) {
	if format == "json" {
		collection.RenderJson(w)
	} else if format == "yaml" {
		collection.RenderYaml(w)
	} else {
		fmt.Fprintln(w, "")
		table := tablewriter.NewWriter(w)
		table.SetHeader([]string{
			"IId",
			"Author",
			"Title",
		})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		for _, issue := range collection.Items {
			table.Append([]string{
				fmt.Sprint(issue.IId),
				issue.Author.Username,
				issue.Title,
			})
		}
		table.Render()
		fmt.Fprintln(w, "")
	}
}
