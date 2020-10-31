package main

import (
	"fmt"

	"github.com/emadghaffari/design-patterns/abstractFactory/packages"
)


func main()  {
	brand,_ := packages.GetSportsFactory("adidas")
	shoe := brand.MakeShoe()
	fmt.Println(shoe.GetLogo())

}