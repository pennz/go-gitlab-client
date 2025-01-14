module github.com/pennz/go-gitlab-client/cli

go 1.16

require (
	github.com/briandowns/spinner v1.12.0
	github.com/fatih/color v1.11.0
	github.com/manifoldco/promptui v0.8.0
	github.com/olekukonko/tablewriter v0.0.5
	github.com/pennz/go-gitlab-client/gitlab v0.0.0-00010101000000-000000000000
	github.com/pennz/go-gitlab-client/test v0.0.0-20210514061436-cc1e84d1b235
	github.com/plouc/textree v1.0.0
	github.com/spf13/cobra v1.1.3
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/pennz/go-gitlab-client/gitlab => ../gitlab
