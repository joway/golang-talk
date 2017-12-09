package benchmark

import "github.com/joway/golang-talk/utils"

func Calc() {
	beg := utils.Now()
	sum := 0.0
	for i := 0; i < 1000000000; i++ {
		sum += 1000000000.123 * 1000000000.321 * 321.123456789
	}
	end := utils.Now()
	utils.Logf("Go : Calc used %d ms\n", end-beg)
}
