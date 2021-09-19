package algo

import (
	"fmt"

	"github.com/pastequo/services.fizzbuzz/internal/entity"
)

// ComputeFizzBuzzString computes fizzbuzz algorithm result.
// params must be a valid entity.
func ComputeFizzBuzzString(params entity.FizzBuzzParams) string {
	ret := ""

	for i := 1; i <= params.Max; i++ {
		if i > 1 {
			ret = fmt.Sprintf("%v,", ret)
		}

		shouldAppend1 := i%params.FirstWord.Multiple == 0
		shouldAppend2 := i%params.SecondWord.Multiple == 0

		if !shouldAppend1 && !shouldAppend2 {
			ret = fmt.Sprintf("%v%d", ret, i)

			continue
		}

		if shouldAppend1 {
			ret = fmt.Sprintf("%v%v", ret, params.FirstWord.Word)
		}

		if shouldAppend2 {
			ret = fmt.Sprintf("%v%v", ret, params.SecondWord.Word)
		}
	}

	return ret
}
