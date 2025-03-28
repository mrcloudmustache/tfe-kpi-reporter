package tfekpi

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ProjectSummaryReport struct {
	Project           string `json:"project"`
	TotalWorkspaces   int    `json:"total_workspaces"`
	TotalResources    int    `json:"total_resources"`
	TotalJobs         int    `json:"total_jobs"`
	SucessfulJobs     int    `json:"sucessful_jobs"`
	FailedJobs        int    `json:"failed_jobs"`
	ResourceAdds      int    `json:"resource_adds"`
	RsourceChanges    int    `json:"resource_changes"`
	ResourceDestrorys int    `json:"resource_destroys"`
}

type ProjectSummaryReports []ProjectSummaryReport

type WorkspaceSummaryReport struct {
	Project           string `json:"project"`
	Workspace         string `json:"workspace"`
	ResourceCount     int    `json:"resource_count"`
	TotalJobs         int    `json:"total_jobs"`
	SucessfulJobs     int    `json:"sucessful_jobs"`
	FailedJobs        int    `json:"failed_jobs"`
	ResourceAdds      int    `json:"resource_adds"`
	RsourceChanges    int    `json:"resource_changes"`
	ResourceDestrorys int    `json:"resource_destroys"`
}

type WorkspaceSummaryReports []WorkspaceSummaryReport

type JobSummaryReport struct {
	Project   string `json:"project"`
	Workspace string `json:"workspace"`
	Job
}

type JobSummaryReports []JobSummaryReport

func (p Projects) SummaryReport() ProjectSummaryReports {
	reports := []ProjectSummaryReport{}
	for _, project := range p {
		w := project.Workspaces
		reports = append(reports, ProjectSummaryReport{
			Project:           project.Name,
			TotalWorkspaces:   project.TotalWorkspaces(),
			TotalJobs:         w.TotalJobs(),
			FailedJobs:        w.FailedJobs(),
			SucessfulJobs:     w.SucessfulJobs(),
			TotalResources:    w.TotalResourcesCount(),
			ResourceAdds:      w.TotalResourceAdds(),
			RsourceChanges:    w.TotalResourceChanges(),
			ResourceDestrorys: w.TotalResourceDestroys(),
		})
	}

	return reports
}

func (p Projects) WorkspaceSummaryReport() WorkspaceSummaryReports {
	reports := []WorkspaceSummaryReport{}
	for _, project := range p {
		for _, w := range project.Workspaces {
			reports = append(reports, WorkspaceSummaryReport{
				Project:           project.Name,
				Workspace:         w.Name,
				TotalJobs:         w.TotalJobs(),
				FailedJobs:        w.FailedJobs(),
				SucessfulJobs:     w.SucessfulJobs(),
				ResourceCount:     w.TotalResourcesCount(),
				ResourceAdds:      w.TotalResourceAdds(),
				RsourceChanges:    w.TotalResourceChanges(),
				ResourceDestrorys: w.TotalResourceDestroys(),
			})
		}

	}

	return reports
}

func (p Projects) JobSummaryReport() JobSummaryReports {
	reports := JobSummaryReports{}
	for _, project := range p {
		for _, w := range project.Workspaces {
			for _, j := range w.Jobs {
				reports = append(reports, JobSummaryReport{
					Project:   project.Name,
					Workspace: w.Name,
					Job:       j,
				})
			}
		}
	}

	return reports
}

func (r ProjectSummaryReports) ToJSON() string {
	marshaled, err := json.MarshalIndent(r, "", "   ")
	if err != nil {
		panic(err)
	}

	return string(marshaled)
}

func (r WorkspaceSummaryReports) ToJSON() string {
	marshaled, err := json.MarshalIndent(r, "", "   ")
	if err != nil {
		panic(err)
	}

	return string(marshaled)
}

func (r JobSummaryReports) ToJSON() string {
	marshaled, err := json.MarshalIndent(r, "", "   ")
	if err != nil {
		panic(err)
	}

	return string(marshaled)
}

func (r ProjectSummaryReports) ToCSV() string {
	var result strings.Builder
	result.WriteString("Project,TotalWorkspaces,TotalResources,TotalJobs,SuccessfulJobs,FailedJobs,ResourceAdds,ResourceChanges,ResourceDestroys\n")

	for _, report := range r {
		row := fmt.Sprintf("%s,%d,%d,%d,%d,%d,%d,%d,%d\n",
			report.Project, report.TotalWorkspaces, report.TotalResources,
			report.TotalJobs, report.SucessfulJobs, report.FailedJobs,
			report.ResourceAdds, report.RsourceChanges, report.ResourceDestrorys)
		result.WriteString(row)
	}

	return result.String()
}

func (r WorkspaceSummaryReports) ToCSV() string {
	var result strings.Builder
	result.WriteString("Project,Workspace,ResourceCount,TotalJobs,SuccessfulJobs,FailedJobs,ResourceAdds,ResourceChanges,ResourceDestroys\n")

	for _, report := range r {
		row := fmt.Sprintf("%s,%s,%d,%d,%d,%d,%d,%d,%d\n",
			report.Project, report.Workspace, report.ResourceCount,
			report.TotalJobs, report.SucessfulJobs, report.FailedJobs,
			report.ResourceAdds, report.RsourceChanges, report.ResourceDestrorys)
		result.WriteString(row)
	}

	return result.String()
}

func (r JobSummaryReports) ToCSV() string {
	var result strings.Builder
	result.WriteString("Project,Workspaces,ID,Status,Adds,Changes,Destroys,StartedAt,FinishedAt\n")

	for _, report := range r {
		row := fmt.Sprintf("%s,%s,%s,%s,%d,%d,%d,%s,%s\n",
			report.Project, report.Workspace, report.ID,
			report.Status, report.Adds, report.Changes,
			report.Destroys, report.StartedAt, report.FinishedAt)
		result.WriteString(row)
	}

	return result.String()
}
