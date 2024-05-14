package Location

type City struct {
	id           string
	name         string
	state        *State
	cityIbgeCode string
}

func (c *City) ID() string {
	return c.id
}

func (c *City) Name() string {
	return c.name
}

func (c *City) State() *State {
	return c.state
}

func (c *City) CityIbgeCode() string {
	return c.cityIbgeCode
}
