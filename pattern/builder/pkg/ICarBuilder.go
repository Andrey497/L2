package pkg

type iCarBuilder interface {
	setWeight()
	setType()
	seMaxSpeed()
	getCar() *car
}

func GetCarBuilder(typeCarBuilder string) iCarBuilder {
	if typeCarBuilder == "truc" {
		return newTrucBuilder()
	}
	if typeCarBuilder == "passengerCar" {
		return newPassengerCarBuilder()
	}
	return nil
}
