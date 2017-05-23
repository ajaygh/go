//when a object contains many objects of its type which may
//also contain many objects of its type then this pattern is
//used. Basically it is useful to represent part or whole
//hierarchy of class tree.

package main

import "fmt"

type Employee struct {
	name         string
	designation  string
	salary       int
	subordinates []*Employee
}

func (emp *Employee) String() string {
	return fmt.Sprintf("[Name : %s Designation : %s Salary : $%d",
		emp.name, emp.designation, emp.salary)
}

func (emp *Employee) add(sub *Employee) {
	emp.subordinates = append(emp.subordinates, sub)
}

func (emp Employee) remove(sub Employee) {
	fmt.Println("Removed employee : ", sub)
}

func (emp *Employee) getSubordinates() []*Employee {
	return emp.subordinates
}

func main() {
	ceo := &Employee{"Mark", "CEO", 340000, nil}
	headSales := &Employee{"John", "HEAD SALES", 300000, nil}
	headMarketing := &Employee{"Tom", "HEAD Marketing", 300000, nil}
	se1 := &Employee{"Tony", "SALES EXECUTIVE", 270000, nil}
	se2 := &Employee{"Rony", "SALES EXECUTIVE", 270000, nil}

	clerk1 := &Employee{"Tim", "CLERK", 12000, nil}
	clerk2 := &Employee{"Tom", "CLERK", 12000, nil}

	ceo.add(headSales)
	ceo.add(headMarketing)

	headSales.add(se1)
	headSales.add(se2)

	headMarketing.add(clerk1)
	headMarketing.add(clerk2)

	fmt.Println(ceo)
	for _, headEmpl := range ceo.getSubordinates() {
		fmt.Println(headEmpl)
		for _, employee := range headEmpl.getSubordinates() {
			fmt.Println(employee)
		}
	}

	fmt.Println("Demo for composite pattern done.")
}
