package _struct

type House struct {
	NumberRoom, Price int
	City              string
}

func (h House) Evaluate() int {
	return h.Price
}

func (h House) IsTwoRoomApartment() bool {
	return h.NumberRoom == 2
}
