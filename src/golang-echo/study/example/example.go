package main

import (
	"errors"
)

func main() {

}
func example(input string) (int, error) {
	if input == "hoge" {
		return 1, nil
	}
	return 0, errors.New("input must be hoge")
}
