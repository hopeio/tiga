package wrapper

type FloatGTValue float64

func (a FloatGTValue) Compare(b FloatGTValue) bool {
	return a > b
}

type FloatLTValue float64

func (a FloatLTValue) Compare(b FloatLTValue) bool {
	return a < b
}
