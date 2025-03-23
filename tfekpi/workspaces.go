package tfekpi

type Workspace struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	ResourceCount int    `json:"resource_count"`
	Jobs          Jobs   `json:"jobs"`
}

type Workspaces []Workspace

func (w Workspaces) TotalResourcesCount() int {
	count := 0
	for _, ws := range w {
		count += ws.ResourceCount
	}
	return count
}

func (w Workspaces) TotalJobs() int {
	count := 0
	for _, ws := range w {
		count += len(ws.Jobs)
	}
	return count
}

func (w Workspaces) SucessfulJobs() int {
	count := 0
	for _, ws := range w {
		for _, j := range ws.Jobs {
			if j.Status == "finished" {
				count++
			}
		}
	}
	return count
}

func (w Workspaces) FailedJobs() int {
	count := 0
	for _, ws := range w {
		for _, j := range ws.Jobs {
			if j.Status != "finished" {
				count++
			}
		}
	}
	return count
}

func (w Workspaces) TotalResourceAdds() int {
	count := 0
	for _, ws := range w {
		for _, j := range ws.Jobs {
			count += j.Adds
		}
	}
	return count
}

func (w Workspaces) TotalResourceChanges() int {
	count := 0
	for _, ws := range w {
		for _, j := range ws.Jobs {
			count += j.Changes
		}
	}
	return count
}

func (w Workspaces) TotalResourceDestroys() int {
	count := 0
	for _, ws := range w {
		for _, j := range ws.Jobs {
			count += j.Destroys
		}
	}
	return count
}

func (w Workspace) TotalResourcesCount() int {
	return w.ResourceCount
}

func (w Workspace) TotalJobs() int {
	return len(w.Jobs)
}

func (w Workspace) SucessfulJobs() int {
	count := 0
	for _, ws := range w.Jobs {
		if ws.Status == "finished" {
			count++
		}
	}
	return count
}

func (w Workspace) FailedJobs() int {
	count := 0
	for _, ws := range w.Jobs {
		if ws.Status != "finished" {
			count++
		}
	}
	return count
}

func (w Workspace) TotalResourceAdds() int {
	count := 0
	for _, j := range w.Jobs {
		count += j.Adds
	}
	return count
}

func (w Workspace) TotalResourceChanges() int {
	count := 0
	for _, ws := range w.Jobs {
		count += ws.Changes
	}
	return count
}

func (w Workspace) TotalResourceDestroys() int {
	count := 0
	for _, ws := range w.Jobs {
		count += ws.Destroys
	}
	return count
}
