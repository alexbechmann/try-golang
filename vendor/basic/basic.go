package basic

var (
	hra   int = 5
	tax   int = 2
	Basic int
)

func Calculation() (allowance int, deduction int) {
	allowance = (Basic * hra) / 100
	deduction = (Basic * tax) / 100
	return
}
