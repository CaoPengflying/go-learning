package combination

import "strconv"

type IOrganization interface {
	Count() int
}

type Employee struct {
	Name string
}

func (e Employee) Count() int {
	return 1
}

type Department struct {
	Name      string
	employees []Employee
	subOrgs   []IOrganization
}

func (d *Department) Count() int {
	count := len(d.employees)
	for _, org := range d.subOrgs {
		count += org.Count()
	}
	return count
}

func (d *Department) AddSub(org IOrganization) {
	d.subOrgs = append(d.subOrgs, org)
}

func (d *Department) AddEmployee(employee Employee) {
	d.employees = append(d.employees, employee)
}

func NewOrg() IOrganization {
	root := &Department{Name: "root"}

	for i := 0; i < 10; i++ {
		name := "000000" + strconv.Itoa(i)
		root.AddEmployee(Employee{Name: name[len(name)-6:]})
		root.AddSub(&Department{Name: "sub" + strconv.Itoa(i), employees: []Employee{Employee{Name: name[len(name)-6:]}}})
	}
	return root
}
