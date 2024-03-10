package main

import (
	"fmt"
	"log"
	"math"
)

type sqrtErr struct {
	lat  string
	long string
	err  error
}

func (se sqrtErr) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(value float64) (float64, error) {
	if value < 0 {
		e := fmt.Errorf("value cannot be negative - value is %f", value)
		return 0, sqrtErr{"50.36537", "90.9393", e}
	}
	return math.Sqrt(value), nil
}
