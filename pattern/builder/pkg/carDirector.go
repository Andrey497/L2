package pkg

type carDirector struct {
	builder iCarBuilder
}

func NewCarDirector(builder iCarBuilder) *carDirector {
	return &carDirector{
		builder: builder,
	}
}

func (d *carDirector) SetCarBuilder(builder iCarBuilder) {
	d.builder = builder
}

func (d *carDirector) BuildCar() *car {
	d.builder.setType()
	d.builder.setWeight()
	d.builder.seMaxSpeed()
	return d.builder.getCar()
}
