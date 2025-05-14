package exception

type ItemCounting struct{}

func (*ItemCounting) Error() string {
	return "Fail to count item"
}
