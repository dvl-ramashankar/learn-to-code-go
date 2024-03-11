// Package mymath provides math solutions relevant to our company.
package mymath

//Return summ of all numbers
func Sum(xi []int) (f int) {
	for _, x := range xi {
		f += x
	}
	return f
}
