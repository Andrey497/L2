package pkg

type plane struct {
	transport
}

func newPlane() iTransport {
	return &boat{transport{typeT: "plane"}}
}
