package main

import (
	"log"
	"sort"
	"strings"

	orderedmap "github.com/wk8/go-ordered-map"
)

// Builds a map of managers to subordinates
func (c *company) buildMapOfManagersToSubordinates() (string, map[string][]string) {
	var mapOfMgrsToSubs = make(map[string][]string)

	var ceo string

	for _, e := range c.Employees {
		if e.ManagerName != "" {
			mapOfMgrsToSubs[e.ManagerName] = append(mapOfMgrsToSubs[e.ManagerName], e.Name)
		} else {
			ceo = e.Name
		}
	}

	return ceo, mapOfMgrsToSubs
}

// Builds the employee hierarchy tree
func (c *company) buildEmployeeHierarchyTree() *orderedmap.OrderedMap {
	ceo, mapOfMgrsToSubs := c.buildMapOfManagersToSubordinates()

	log.Print("****Hierarchy tree****")

	var managers []string

	// Use Ordered map to print hierarchy tree, because the native Go map does not guarantee order
	orderedMap := orderedmap.New()
	orderedMap.Set(ceo, mapOfMgrsToSubs[ceo])

	managers = append(managers, ceo)

	// For each level, add manager and his/her subordinates to the ordered map
	for len(managers) > 0 {
		mgrs := mapOfMgrsToSubs[managers[0]]
		managers = managers[1:]
		for _, v := range mgrs {
			if _, ok := mapOfMgrsToSubs[v]; ok {
				orderedMap.Set(v, mapOfMgrsToSubs[v])
				managers = append(managers, v)
			}
		}
	}

	return orderedMap
}

func printEmployeeHierarchyTree(om *orderedmap.OrderedMap) {

	// Starting from the top, print a list of managers and his/her subordinates at each level
	for pair := om.Oldest(); pair != nil; pair = pair.Next() {
		log.Print(pair.Key)
		log.Printf("Employees of %s: %s", pair.Key, strings.Join(pair.Value.([]string), " "))
	}

	log.Print("\r\n")
}

// Calculate the total salary of all the employees
func (c *company) calculateTotalSalary() float64 {
	var salary float64

	for _, e := range c.Employees {
		salary += e.Salary
	}

	return salary
}

// Sort all the employees alphabetically
func (c *company) sortEmployees() {
	sort.Slice(c.Employees, func(i, j int) bool {
		return c.Employees[i].Name < c.Employees[j].Name
	})
}
