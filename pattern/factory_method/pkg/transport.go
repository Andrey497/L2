package pkg

type transport struct {
	typeT  string
	model  string
	color  string
	weight int
	speed  int
}

func (t *transport) SetColor(color string) {
	t.color = color
}
func (t *transport) GetColor() string {
	return t.color
}
func (t *transport) SetModel(model string) {
	t.model = model
}
func (t *transport) GetModel() string {
	return t.model
}
func (t *transport) SetSpeed(speed int) {
	t.speed = speed
}
func (t *transport) GetSpeed() int {
	return t.speed
}
func (t *transport) GetWeight() int {
	return t.weight
}
func (t *transport) GetType() string {
	return t.typeT
}
