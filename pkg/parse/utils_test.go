package main

import (
	"testing"

	orderedmap "github.com/wk8/go-ordered-map"

	"github.com/stretchr/testify/assert"
)

var c = company{Employees: []employee{{
	Name:        "Danny",
	ManagerName: "Catherine",
	Salary:      150000,
	IsManager:   true,
},
	{
		Name:        "Cindy",
		ManagerName: "Danny",
		Salary:      100000,
		IsManager:   true,
	},
	{
		Name:      "Catherine",
		Salary:    200000,
		IsManager: true,
	},
	{
		Name:        "Linda",
		ManagerName: "Cindy",
		Salary:      95000,
		IsManager:   false,
	},
}}

func TestCalculateTotalSalary(t *testing.T) {
	assert.Equal(t, float64(545000), c.calculateTotalSalary())
}

func TestSortEmployees(t *testing.T) {
	c.sortEmployees()
	var employeeList []string
	for _, e := range c.Employees {
		employeeList = append(employeeList, e.Name)
	}
	assert.Equal(t, []string{"Catherine", "Cindy", "Danny", "Linda"}, employeeList)
}

func TestBuildMapOfManagersToSubordinates(t *testing.T) {
	ceo, mapOfMgrsToSubs := c.buildMapOfManagersToSubordinates()
	expectedMap := map[string][]string{"Catherine": {"Danny"}, "Cindy": {"Linda"}, "Danny": {"Cindy"}}
	assert.Equal(t, "Catherine", ceo)
	assert.Equal(t, expectedMap, mapOfMgrsToSubs)
}

func TestBuildEmployeeHierarchyTree(t *testing.T) {
	om := c.buildEmployeeHierarchyTree()

	expectedOm := orderedmap.New()
	expectedOm.Set("Catherine", []string{"Danny"})
	expectedOm.Set("Danny", []string{"Cindy"})
	expectedOm.Set("Cindy", []string{"Linda"})

	assert.Equal(t, expectedOm, om)
}
