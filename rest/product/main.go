package main

func main() {
	a := App{}
	a.Initialize("product.db")
	a.Run(":8080")
}
