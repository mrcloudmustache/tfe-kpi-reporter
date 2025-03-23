package tfekpi

import "encoding/json"

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

func (r ProjectSummaryReports) ToJSON() string {
	marshaled, err := json.MarshalIndent(r, "", "   ")
	if err != nil {
		panic(err)
	}
	return string(marshaled)
}

func (w WorkspaceSummaryReports) ToJSON() string {
	marshaled, err := json.MarshalIndent(w, "", "   ")
	if err != nil {
		panic(err)
	}
	return string(marshaled)
}
