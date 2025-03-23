package tfekpi

type Job struct {
	ID         string `json:"id"`
	Status     string `json:"status"`
	Adds       int    `json:"adds"`
	Changes    int    `json:"changes"`
	Destroys   int    `json:"destroys"`
	StartedAt  string `json:"started_at"`
	FinishedAt string `json:"finished_at"`
}

type Jobs []Job

type JobMetrics interface {
	TotalJobs()
	FailedJobs()
	SuccessfulJobs()
	TotalResourceCount()
	TotalResourceAdds()
	TotalResourceChanges()
	TotalResourceDestroys()
}
