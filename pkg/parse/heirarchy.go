package main

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type employee struct {
	Name        string  `yaml:"name"`
	ManagerName string  `yaml:"manager_name"`
	Salary      float64 `yaml:"salary"`
	IsManager   bool    `yaml:"isManager"`
}

type company struct {
	Employees []employee `yaml:"employees"`
}

func (c *company) processEmployees() {
	// Do not set any flags, keep the logging simple
	log.SetFlags(0)

	/* Read the details of all the employees via a YAML file supplied via
	   a command line argument */
	yamlFile, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("Error reading yaml file: %#v", err)
	}

	// Unmarshall into an array of employees
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Error unmarshalling: %#v", err)
	}

	// Print the Employee hierarchy tree, along with total salary
	om := c.buildEmployeeHierarchyTree()

	printEmployeeHierarchyTree(om)

	totalSalary := c.calculateTotalSalary()

	log.Printf("Total salary: %f", totalSalary)

	// Now sort all the employees and print out the names of the sorted employees
	c.sortEmployees()

	log.Print("\r\n")

	log.Print("Sorted Employee list:")

	for _, e := range c.Employees {
		log.Print(e.Name)
	}
}
