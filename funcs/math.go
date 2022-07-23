package mathFunction

import (
	"math"
	"math/rand"
	"time"
)

type MathFunction interface {
	Calculate(arg float64) float64
	GetName() string
}

type Sin struct {
	Name string
}

func (s Sin) Calculate(arg float64) float64 {
	return math.Sin(arg * (math.Pi / 180))
}

func (s Sin) GetName() string {
	return s.Name
}

type Cos struct {
	Name string
}

func (c Cos) Calculate(arg float64) float64 {
	return math.Cos(arg * (math.Pi / 180))
}

func (c Cos) GetName() string {
	return c.Name
}

type Log struct {
	Name string
}

func (l Log) Calculate(arg float64) float64 {
	return math.Log10(arg)
}

func (l Log) GetName() string {
	return l.Name
}

func MathFunctionFactory() MathFunction {
	var mf MathFunction
	rand.Seed(time.Now().UnixNano())
	i := rand.Int31n(3)

	switch i {
	case 0:
		mf = Sin{"Sinus"}
	case 1:
		mf = Cos{"Cosinus"}
	case 2:
		mf = Log{"Log"}
	}
	return mf
}
