package main

import "github.com/apaxa-go/eval"

func Evaluate(item NyaaTorrentItem, expr *eval.Expression) (result bool, err error) {
	arg := eval.Args{
		"item": eval.MakeDataRegularInterface(item),
	}

	r, err := expr.EvalToInterface(arg)
	if err != nil {
		return
	}

	result = r.(bool)

	return

}

// FilterNyaaItems filter out items that does not match the criteria.
func FilterNyaaItems(items []NyaaTorrentItem, expr string) []NyaaTorrentItem {
	var out []NyaaTorrentItem

	exprObj, err := eval.ParseString(expr, "")
	if err != nil {
		return out
	}

	for _, i := range items {
		if result, e := Evaluate(i, exprObj); e != nil || !result {
			continue
		}

		out = append(out, i)
	}

	return out
}
