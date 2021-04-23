# Employee Hierarchy

This Go based project prints out the Employee hierarchy for a given list of employees provided via a YAML file.
An example of the YAML file has been provided in the root folder of this repo, by the name employee_list.yaml.

## Build & Run

The YAML file needs to be provided as command line argument.

From the root folder of this repo, run
```go
go run ./... employee_list.yaml
```