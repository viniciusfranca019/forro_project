package Location

type Country struct {
	id           string
	name         string
	abbreviation string
}

func (c *Country) ID() string {
	return c.id
}

func (c *Country) Name() string {
	return c.name
}

func (c *Country) Abbreviation() string {
	return c.abbreviation
}
