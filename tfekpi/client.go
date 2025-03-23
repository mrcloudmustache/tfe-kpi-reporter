package tfekpi

import (
	"context"
	"log"

	"github.com/hashicorp/go-tfe"
)

type TFE struct {
	Client tfe.Client
	Org    string
}

func (t TFE) ListProjects() Projects {
	response, err := t.Client.Projects.List(context.Background(), t.Org, nil)
	if err != nil {
		log.Fatal(err)
	}
	projects := Projects{}
	for _, i := range response.Items {
		projects = append(projects, Project{
			ID:          i.ID,
			Name:        i.Name,
			Description: i.Description,
		})

	}
	return projects
}

func (t TFE) ListProjectWorkspaces(project_id string) Workspaces {
	options := &tfe.WorkspaceListOptions{ProjectID: project_id}
	response, err := t.Client.Workspaces.List(context.Background(), t.Org, options)
	if err != nil {
		log.Fatal(err)
	}

	workspaces := Workspaces{}
	for _, i := range response.Items {
		workspaces = append(workspaces, Workspace{
			ID:            i.ID,
			Name:          i.Name,
			ResourceCount: i.ResourceCount,
		})
	}
	return workspaces
}

func (t TFE) listWorkspacesRuns(workspace_id string) []*tfe.Run {
	response, err := t.Client.Runs.List(context.Background(), workspace_id, nil)
	if err != nil {
		log.Fatal(err)
	}
	return response.Items
}

func (t TFE) ListWorkspacesJobs(workspace_id string) Jobs {
	jobs := Jobs{}
	runs := t.listWorkspacesRuns(workspace_id)
	for _, i := range runs {
		apply, err := t.Client.Applies.Read(context.Background(), i.Apply.ID)
		if err != nil {
			log.Fatal(err)
		}
		jobs = append(jobs, Job{
			ID:         apply.ID,
			Status:     string(apply.Status),
			Adds:       apply.ResourceAdditions,
			Changes:    apply.ResourceChanges,
			Destroys:   apply.ResourceDestructions,
			StartedAt:  apply.StatusTimestamps.StartedAt.String(),
			FinishedAt: apply.StatusTimestamps.FinishedAt.String(),
		})

	}
	return jobs
}

func LoadProjects(client TFE) Projects {
	projects := client.ListProjects()
	for i, project := range projects {
		ws := client.ListProjectWorkspaces(project.ID)
		projects[i].Workspaces = ws

		for j := range projects[i].Workspaces {
			jobs := client.ListWorkspacesJobs(projects[i].Workspaces[j].ID)
			projects[i].Workspaces[j].Jobs = jobs
		}
	}
	return projects
}
