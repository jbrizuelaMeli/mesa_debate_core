package contact

type Service interface {
	GetAllContacts() []*Contact
	GetContact(id string) *Contact
	ChangeStatus(id string) *Contact
}

type service struct {
	cnts []*Contact
}

func NewService() Service {
	return &service{
		cnts: MockContacts(),
	}
}

func (s *service) GetAllContacts() []*Contact {
	return s.cnts
}

func (s *service) GetContact(id string) *Contact {
	for _, cnt := range s.cnts {
		if cnt.ID == id {
			return cnt
		}
	}
	return nil
}

func (s *service) ChangeStatus(id string) *Contact {
	for _, cnt := range s.cnts {
		if cnt.ID == id {
			status := "active"
			if cnt.Status == "active" {
				status = "disabled"
			}
			cnt.Status = status
			return cnt
		}
	}
	return nil
}