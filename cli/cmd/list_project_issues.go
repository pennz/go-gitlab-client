package cmd

import (
	"fmt"

	"github.com/fatih/color"
	out "github.com/pennz/go-gitlab-client/cli/output"
	"github.com/pennz/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
)

func init() {
	listCmd.AddCommand(listProjectIssuesCmd)
}

func fetchProjectIssues(projectId string) {
	color.Yellow("Fetching project issues (project id: %s)…", projectId)

	o := &gitlab.IssuesOptions{}
	o.Page = page
	o.PerPage = perPage

	loader.Start()
	collection, meta, err := client.ProjectIssues(projectId, o)
	loader.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(collection.Items) == 0 {
		color.Red("No issue found for project %s", projectId)
	} else {
		out.Issues(output, outputFormat, collection)
	}

	printMeta(meta, true)

	handlePaginatedResult(meta, func() {
		fetchProjectIssues(projectId)
	})
}

var listProjectIssuesCmd = &cobra.Command{
	Use:     resourceCmd("project-issues", "project"),
	Aliases: []string{"pi"},
	Short:   "List project issues",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project", args)
		if err != nil {
			return err
		}

		fetchProjectIssues(ids["project_id"])

		return nil
	},
}
