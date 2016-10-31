package lib

import (
	"fmt"
)

func P(z ...interface{}) {
	var x []interface{}
	x = append(x, "[ Secretary ]")
	for i := range z {
		x = append(x, z[i])
	}
	fmt.Println(x...)
}

func PN() {
	fmt.Println("")
}
