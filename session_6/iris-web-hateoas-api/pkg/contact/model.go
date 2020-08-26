package contact

import "time"

type Contact struct {
	ID         string  `json:"id"`
	Type       string  `json:"type"`
	Value      string  `json:"value"`
	SiteId     string  `json:"site"`
	UserOwner  *User   `json:"user"`
	Status     string  `json:"status,omitempty"`
	LastUpdate string  `json:"last_modified_date"`
	Links      map[string]*Link `json:"links,omitempty"`
}

// User
type User struct {
	ID        string  `json:"id"`
	FirstName string  `json:"first_name,omitempty"`
	LastName  string  `json:"last_name,omitempty"`
	Links     map[string]*Link `json:"links,omitempty"`
}

// User
type Link struct {
	Href    string `json:"href"`
	Method  string `json:"method"`
}

func MockContacts() []*Contact {
	cnts := make([]*Contact, 0)
	cnt1 := Contact{
		ID: "1",
		Type: "email",
		Value: "luis.stubbia@mercadolibre.com",
		SiteId: "MLA",
		UserOwner: &User{
			ID: "10000",
			FirstName: "Luis",
			LastName: "Stubbia",
		},
		Status: "active",
		LastUpdate: time.Now().String(),
	}
	cnts = append(cnts, &cnt1)

	cnt2 := Contact{
		ID: "2",
		Type: "email",
		Value: "pancho.salmaso@mercadolibre.com",
		SiteId: "MLA",
		UserOwner: &User{
			ID: "20000",
			FirstName: "Pancho",
			LastName: "Salmaso",
		},
		Status: "active",
		LastUpdate: time.Now().String(),
	}
	cnts = append(cnts, &cnt2)
	return cnts
}
