package gonnie

import (
	"fmt"
)

func printConector(next func(), e *ExchangeMessage, out Message, u Uri, params ...interface{}) error {
	fmt.Println(u.options.Get("msg"))
	out <- e
	next()
	return nil
}