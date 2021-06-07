package hw2lib

import (
	"github.com/jakehl/goid"
	"math"
	"strings"
)

func Tolower(str string) string {
	return strings.ToLower(str)
}

func Toupper(str string) string {
	return strings.ToUpper(str)
}

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func GetUUID() *goid.UUID {
	v4UUID := goid.NewV4UUID()
	return v4UUID
}

