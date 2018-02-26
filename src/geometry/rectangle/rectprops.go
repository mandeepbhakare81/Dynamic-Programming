package  rectangle
import (
     "math"
)


func Diagonal(len, width float64) float64{
	diagonal := math.Sqrt((len*len)+(width*width))
	return diagonal
}
