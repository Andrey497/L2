package pkg

type boat struct {
	transport
}

func newBoat() iTransport {
	return &boat{transport{typeT: "boat"}}
}
