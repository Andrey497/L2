package pkg

type plane struct {
	model string
	sped  int
}

func (p *plane) Accept(v visitor) {
	v.visitorPlane(p)
}

func (p *plane) GetType() string {
	return "plane"
}
func NewPlane() *plane {
	return &plane{}
}
