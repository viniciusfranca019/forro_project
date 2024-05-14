package Location

type State struct {
	id           string
	name         string
	abbreviation string
	country      string
}

func (s *State) ID() string {
	return s.id
}

func (s *State) Name() string {
	return s.name
}

func (s *State) Country() string {
	return s.country
}

func (s *State) Abbreviation() string {
	return s.abbreviation
}
