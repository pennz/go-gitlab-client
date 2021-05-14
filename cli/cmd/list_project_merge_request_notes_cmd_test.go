package cmd

import (
	"testing"

	"github.com/pennz/go-gitlab-client/test"
)

func TestListProjectMergeRequestNotesCmd(t *testing.T) {
	test.RunCommandTestCases(t, "notes", []*test.CommandTestCase{
		{
			[]string{"list", "project-merge-request-notes", "--help"},
			nil,
			//configs["default"],
			"list_project_merge_request_notes_help",
			false,
			nil,
		},
		{
			[]string{"list", "project-merge-request-notes"},
			nil,
			//configs["default"],
			"list_project_merge_request_notes_no_project_id",
			true,
			nil,
		},
		{
			[]string{"list", "project-merge-request-notes", "1"},
			nil,
			//configs["default"],
			"list_project_merge_request_notes_no_merge_request_iid",
			true,
			nil,
		},
		{
			[]string{"list", "project-merge-request-notes", "1", "3"},
			nil,
			//configs["default"],
			"list_project_merge_request_notes",
			false,
			nil,
		},
		{
			[]string{"list", "project-merge-request-notes", "1", "3", "-f", "json"},
			nil,
			//configs["default"],
			"list_project_merge_request_notes_json",
			false,
			nil,
		},
		{
			[]string{"list", "project-merge-request-notes", "1", "3", "-f", "yaml"},
			nil,
			//configs["default"],
			"list_project_merge_request_notes_yaml",
			false,
			nil,
		},
	})
}
