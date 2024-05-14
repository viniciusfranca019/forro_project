package Location

type State struct {
	id           string
	name         string
	abbreviation string
	country      *Country
}

func (s *State) ID() string {
	return s.id
}

func (s *State) Name() string {
	return s.name
}

func (s *State) Country() *Country {
	return s.country
}

func (s *State) Abbreviation() string {
	return s.abbreviation
}
