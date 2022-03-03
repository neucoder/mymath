package main

import (
	"errors"
)

func Add(a int, b int) (ret int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("必须是非负的两个数")
	}
	return a + b, err
}
