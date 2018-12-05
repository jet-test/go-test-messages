package obj

import "testing"

func TestContact_String(t *testing.T) {
	tests := []struct {
		name    string
		contact Contact
		want    string
	}{
		{"empty", Contact{}, "Contact = { Class: , Value:  }"},
		{"mail", Contact{"mail", "mail@mail.ru"}, "Contact = { Class: mail, Value: mail@mail.ru }"},
		{"phone", Contact{"phone", "+7913..."}, "Contact = { Class: phone, Value: +7913... }"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Contact{
				Class: tt.contact.Class,
				Value: tt.contact.Value,
			}
			if got := c.String(); got != tt.want {
				t.Errorf("Contact.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
