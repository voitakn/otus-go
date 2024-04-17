package order

type MyOrd struct {
	Name string
	Age  int
}

type Orders []MyOrd

func MyOrder() string {
	return "MyOrder()"
}
