package pkg

type car struct {
	transport
}

func newCar() iTransport {
	return &boat{transport{typeT: "car"}}
}
