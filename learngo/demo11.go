package main

import (
	"fmt"
)

var container = []string{"zero", "one", "two"}

func main() {
	container := map[int]string{0:"zero", 1:"one", 2:"two"}

	v, ok := container[6]

	fmt.Printf("The element is %v, %v. (%v)\n", v, ok, container[7])

	value, ok := interface{}(container).(map[int]string)

	fmt.Printf("value:%v, ok:%v\n", value, ok)

	elem, err := getElement(container)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("The element is %q. (container type: %T)\n",
		elem, container)


	test := string(-1)
	fmt.Printf("%v\n", test)

	test = string([]byte{'\xe4', '\xbd', '\xa0', '\xe5', '\xa5', '\xbd'})
	fmt.Printf("test:%v\n", test)

}

func getElement(containerI interface{}) (elem string, err error) {
	switch t := containerI.(type) {
	case []string:
		elem = t[1]
	case map[int]string:
		fmt.Printf("======>t:%v\n", t)
		elem = t[1]
	default:
		err = fmt.Errorf("unsupported container type: %T", containerI)
		return
	}
	return
}
