package main

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

func main() {

	params := make(map[string]interface{})
	params["instrument-1.parameter-1"] = 15
	// params["instrument-1.parameter-2"] = 100
	expression, err := govaluate.NewEvaluableExpression("(2 * [instrument-1.parameter-1] + [instrument-1.parameter-2] ** 2)")
	if err != nil {
		fmt.Println(err)
	}
	result, err := expression.Evaluate(params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
