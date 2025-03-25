package main

import (
	"fmt"
	"log"
	"tfe-kpi-reporter/tfekpi"

	"github.com/hashicorp/go-tfe"
)

func main() {
	config := &tfe.Config{
		Address:           "https://app.terraform.io",
		Token:             "",
		RetryServerErrors: true,
	}

	client, err := tfe.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	org := "myorganization"
	tfe := tfekpi.TFE{Client: *client, Org: org}
	projects := tfekpi.LoadProjects(tfe)
	// fmt.Println(projects.SummaryReport().ToJSON())
	// fmt.Println(projects.WorkspaceSummaryReport().ToJSON())
	fmt.Println(projects.JobSummaryReport().ToJSON())
}
