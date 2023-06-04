package variables

import "fmt"

func MuestroEnteros() {

	x := 5
	isVerdad := true
	intde32bits := int32(5)
	intde64bits := int64(5)
	if isVerdad {
		fmt.Println(x, intde32bits, intde64bits)
	}
}
