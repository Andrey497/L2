package pkg

func GetTransportByType(typ string) iTransport {
	if typ == "boat" {
		return newBoat()
	}
	if typ == "plane" {
		return newPlane()
	}
	if typ == "car" {
		return newCar()
	}
	return nil
}
