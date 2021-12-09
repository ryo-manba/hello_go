package main

import "fmt"

func AorAn(s string) string {
	if s[0] == 'i' || s[:2] == "*i" {
		return "an"
	} else {
		return "a"
	}
}

func PutVars(v, varType string) {
	fmt.Printf("%s is %s %s.\n", v, AorAn(varType), varType)
}

func main() {
	type FortyTwo struct {}
	vars := []interface{}{
		"42",
		uint(42),
		42,
		uint8(42),
		int16(42),
		uint32(42),
		int64(42),
		false,
		float32(42),
		float64(42),
		complex64(42 + 0i),
		42 + 21i,
		FortyTwo{},
		[...]int{42},
		map[string]int{"42": 42},
		(*int)(nil),
		[]int{},
		chan int(nil),
		nil,
	}

	for _, v := range vars {
		if v == (*int)(nil) {
			PutVars(fmt.Sprintf("%p", v), fmt.Sprintf("%T", v))
			continue
		}
		PutVars(fmt.Sprintf("%v", v), fmt.Sprintf("%T", v))
	}
}
