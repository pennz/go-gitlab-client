package gitlab

import (
	"encoding/json"
	"io"
)

const (
	ProjectIssuesApiPath = "/projects/:id/issues"
)

type Issue struct {
	Id          int        `json:"id"`
	IId         int        `json:"iid"`
	ProjectId   int        `json:"project_id,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Labels      []string   `json:"labels,omitempty"`
	Milestone   *Milestone `json:"milestone,omitempty"`
	Assignee    *User      `json:"assignee,omitempty"`
	Author      *User      `json:"author,omitempty"`
	State       string     `json:"state,omitempty"`
	CreatedAt   string     `json:"created_at,omitempty"`
	UpdatedAt   string     `json:"updated_at,omitempty"`
}

type IssueCollection struct {
	Items []*Issue
}

type IssuesOptions struct {
	PaginationOptions

	AuthorID         int      `json:"author_id,omitempty"`
	AuthorUsername   int      `json:"author_username,omitempty"`
	AssigneeID       int      `json:"assignee_id,omitempty"`
	AssigneeUsername []string `json:"assignee_username,omitempty"`
	IIDs             []int    `json:"iids,omitempty"`
}

type IssueRequest struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	AssigneeId  int    `json:"assignee_id,omitempty"`
	MilestoneId int    `json:"milestone_id,omitempty"`
	Labels      string `json:"labels,omitempty"`
}

func (g *Gitlab) AddIssue(projectId string, req *IssueRequest) (issue *Issue, meta *ResponseMeta, err error) {
	params := map[string]string{
		":id": projectId,
	}
	u := g.ResourceUrl(ProjectIssuesApiPath, params)

	encodedRequest, err := json.Marshal(req)
	if err != nil {
		return
	}

	data, _, err := g.buildAndExecRequest("POST", u.String(), encodedRequest)
	if err != nil {
		return
	}

	issue = new(Issue)
	err = json.Unmarshal(data, issue)
	if err != nil {
		panic(err)
	}

	return
}

// ProjectIssues get all the issue for this project
func (g *Gitlab) ProjectIssues(projectId string, o *IssuesOptions) (collection *IssueCollection, meta *ResponseMeta, err error) {
	u := g.ResourceUrlQ(ProjectIssuesApiPath, map[string]string{
		":id": projectId,
	}, o)

	collection = new(IssueCollection)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &collection.Items)
	}

	return collection, meta, err
}

func (c *IssueCollection) RenderJson(w io.Writer) error {
	return renderJson(w, c.Items)
}

func (c *IssueCollection) RenderYaml(w io.Writer) error {
	return renderYaml(w, c.Items)
}
