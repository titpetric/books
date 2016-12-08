package subpackage1

import "fmt"
import "app/subpackage2"

func Hello() {
	return subpackage2.Hello();
}