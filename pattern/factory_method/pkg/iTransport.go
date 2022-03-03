package pkg

type iTransport interface {
	SetColor(string)
	GetColor() string
	SetModel(string)
	GetModel() string
	SetSpeed(int)
	GetSpeed() int
	GetWeight() int
	GetType() string
}
