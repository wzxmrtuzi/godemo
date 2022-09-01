package reference

import (
	"fmt"
)

func Quote() {
	num := 123
	changeQuote(num, 555)
	fmt.Println("quote", num)
}

func changeQuote(num, targeNum int) {
	num = targeNum
}
