package kernel

import "fmt"

func Println(p ...interface{}) error {
	_, err := fmt.Println(p...)
	return err
}
