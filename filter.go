package main

import (
	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"
)

// Evaluate if the current item pass the thresshold
func Evaluate(item NyaaTorrentItem, program *vm.Program) (result bool, err error) {
	res, err := expr.Run(program, item)
	if err != nil {
		return
	}

	result = res.(bool)

	return
}

// FilterNyaaItems filter out items that does not match the criteria.
func FilterNyaaItems(expression string, items []NyaaTorrentItem) (out []NyaaTorrentItem, err error) {

	program, err := expr.Compile(expression, expr.Env(NyaaTorrentItem{}))
	if err != nil {
		return
	}

	for _, i := range items {
		if result, e := Evaluate(i, program); e != nil || !result {
			continue
		}

		out = append(out, i)
	}

	return
}
