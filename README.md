# Terraform Cloud KPI Reporter

A Go application that generates actionable metrics and reports from your Terraform Cloud organization. This tool helps DevOps teams and platform engineers track infrastructure changes, success rates, and resource utilization across projects and workspaces.

## Features

- **Project Summary Reports**: Aggregate metrics for each project including workspace count, resource totals, and job success rates
- **Workspace Detail Reports**: Detailed metrics for individual workspaces including resource counts and job statistics
- **Job Performance Analysis**: Track successful vs failed jobs across your organization
- **Resource Change Tracking**: Monitor resource additions, modifications, and deletions

## Requirements

- Go 1.23 or higher
- Terraform Cloud account with API access
- API token with read permissions for your organization

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/tfe-kpi-reporter.git
cd tfe-kpi-reporter

# Install dependencies
go mod download
```

## Configuration

Set your Terraform Cloud API token before running the application:

```go
// In main.go
config := &tfe.Config{
    Address: "https://app.terraform.io",
    Token:   "your-terraform-cloud-api-token", // Replace with your actual token
    RetryServerErrors: true,
}
```

Change the organization name if needed:

```go
// In main.go
org := "your-organization-name" // Replace with your actual organization name
```

## Usage

Run the application to generate reports:

```bash
go run main.go
```

### Example Output

#### Project Summary Report

```json
[
   {
      "project": "Production Infrastructure",
      "total_workspaces": 4,
      "total_resources": 156,
      "total_jobs": 27,
      "sucessful_jobs": 25,
      "failed_jobs": 2,
      "resource_adds": 145,
      "resource_changes": 38,
      "resource_destroys": 12
   },
   {
      "project": "Development Environment",
      "total_workspaces": 2,
      "total_resources": 42,
      "total_jobs": 15,
      "sucessful_jobs": 13,
      "failed_jobs": 2,
      "resource_adds": 42,
      "resource_changes": 21,
      "resource_destroys": 5
   }
]
```

#### Workspace Summary Report

```json
[
   {
      "project": "Production Infrastructure",
      "workspace": "networking",
      "resource_count": 38,
      "total_jobs": 8,
      "sucessful_jobs": 8,
      "failed_jobs": 0,
      "resource_adds": 38,
      "resource_changes": 12,
      "resource_destroys": 2
   },
   {
      "project": "Production Infrastructure",
      "workspace": "kubernetes-cluster",
      "resource_count": 64,
      "total_jobs": 12,
      "sucessful_jobs": 10,
      "failed_jobs": 2,
      "resource_adds": 64,
      "resource_changes": 18,
      "resource_destroys": 4
   }
]
```

## Architecture

### Core Components

- **Client**: Interface with Terraform Cloud API
- **Projects**: Representation of TFC projects with metrics
- **Workspaces**: Collection of workspaces with associated jobs and resources
- **Jobs**: Job execution details with success/failure statistics
- **Reports**: Generation of standardized output formats

### Code Structure

```
tfe-kpi-reporter/
├── main.go              # Application entry point
├── tfekpi/
│   ├── client.go        # TF Cloud API client
│   ├── projects.go      # Project data structures and methods
│   ├── workspaces.go    # Workspace data structures and metrics
│   ├── jobs.go          # Job data structures and statistics
│   └── reports.go       # Report generation and formatting
```

## Report Types

### Project Summary Report

Provides a high-level overview of each project:

| Field | Description |
|-------|-------------|
| `project` | Project name |
| `total_workspaces` | Number of workspaces in the project |
| `total_resources` | Total resources managed across all workspaces |
| `total_jobs` | Total number of jobs executed |
| `successful_jobs` | Number of successfully completed jobs |
| `failed_jobs` | Number of failed jobs |
| `resource_adds` | Total resources added across all jobs |
| `resource_changes` | Total resources modified across all jobs |
| `resource_destroys` | Total resources destroyed across all jobs |

### Workspace Summary Report

Provides detailed metrics for each workspace:

| Field | Description |
|-------|-------------|
| `project` | Project name |
| `workspace` | Workspace name |
| `resource_count` | Number of resources managed in this workspace |
| `total_jobs` | Total number of jobs executed |
| `successful_jobs` | Number of successfully completed jobs |
| `failed_jobs` | Number of failed jobs |
| `resource_adds` | Total resources added across all jobs |
| `resource_changes` | Total resources modified across all jobs |
| `resource_destroys` | Total resources destroyed across all jobs |

## Use Cases

- **Infrastructure Audit**: Track what resources are being provisioned and modified
- **Team Performance**: Monitor job success rates to identify problematic areas
- **Cost Analysis**: Track resource additions and deletions over time
- **Compliance Reporting**: Generate evidence for compliance audits

## Example: Generating Filtered Reports

You can modify the application to filter reports by project or date range:

```go
// Example: Filter for specific projects
filteredProjects := projects.FilterByName("Production")
fmt.Println(filteredProjects.SummaryReport().ToJSON())
```

## Integrating with Other Tools

### Export to CSV (Example Implementation)

```go
// Add to reports.go
func (r ProjectSummaryReports) ToCSV() string {
    var result strings.Builder
    result.WriteString("Project,Workspaces,Resources,Jobs,Successful,Failed,Adds,Changes,Destroys\n")
    
    for _, report := range r {
        row := fmt.Sprintf("%s,%d,%d,%d,%d,%d,%d,%d,%d\n",
            report.Project, report.TotalWorkspaces, report.TotalResources,
            report.TotalJobs, report.SucessfulJobs, report.FailedJobs,
            report.ResourceAdds, report.RsourceChanges, report.ResourceDestrorys)
        result.WriteString(row)
    }
    
    return result.String()
}
```

Example output:
```
Project,Workspaces,Resources,Jobs,Successful,Failed,Adds,Changes,Destroys
Production Infrastructure,4,156,27,25,2,145,38,12
Development Environment,2,42,15,13,2,42,21,5
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- [HashiCorp Terraform Cloud](https://www.terraform.io/cloud) for providing the API
- [go-tfe](https://github.com/hashicorp/go-tfe) for the Terraform Cloud Go client
