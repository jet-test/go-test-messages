package obj

type Contact struct {
	Class string
	Value string
}

func (c Contact) String() string {
	return "Contact = { " +
		"Class: " + c.Class + ", " +
		"Value: " + c.Value + " }"
}
