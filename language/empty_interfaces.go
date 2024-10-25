package language

import (
	"fmt"
	"reflect"
)

func init() {
	emptyInterfaceSample()
}

func emptyInterfaceSample() {
	sum, _ := add(1, 2)
	fmt.Println("SUM:", sum)
	concat, _ := add("1", "2")
	fmt.Println("CONCAT:", concat)
	conjunction, _ := add(true, true)
	fmt.Println("CONJUNCTION", conjunction)
	_, err := add(1.0, 2.4)
	fmt.Println("ERROR:", err)
}

func add(a interface{}, b interface{}) (result interface{}, err any) {
	defer func() {
		if r := recover(); r != nil {
			err = r
		}
	}()

	switch a.(type) {
	case int:
		result = a.(int) + b.(int)
		return
	case bool:
		result = a.(bool) && b.(bool)
		return
	case string:
		result = a.(string) + b.(string)
		return
	default:
		panic("Unsupported type of arguments identified: " + reflect.TypeOf(a).Name())
	}
}
