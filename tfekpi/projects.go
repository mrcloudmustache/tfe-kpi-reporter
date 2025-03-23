package tfekpi

import "encoding/json"

type Project struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Workspaces  Workspaces `json:"workspaces"`
}

type Projects []Project

func (p Project) TotalWorkspaces() int {
	return len(p.Workspaces)
}

func (p Projects) ToJson() string {
	marshaled, err := json.MarshalIndent(p, "", "   ")
	if err != nil {
		panic(err)
	}
	return string(marshaled)
}
