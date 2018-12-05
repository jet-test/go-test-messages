package obj

import "fmt"

type Person struct {
	Name     string
	Contacts []Contact
}

func (p Person) String() string {
	return "person = { name: " + p.Name + ", " +
		"contacts: " + fmt.Sprint(p.Contacts) + " }"
}
