package hw2lib

import (
	"math"
	"strings"
	"github.com/twharmon/gouid"
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

func GetUUID() string {
	return gouid.String(16, gouid.MixedCaseAlpha)
}

