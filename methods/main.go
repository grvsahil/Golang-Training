package main

import(
	"fmt"
)

type bike struct{
	name string
	category string
	displacement int
	engineType string
}

func (b bike) printInfo(){
	fmt.Printf("%s\n%s\n%d cc\n%s engine\n\n",b.name,b.category,b.displacement,b.engineType)
}

func main()  {
	bike1 := bike{
		"Bajaj Pulsar 220f",
		"Sports Tourer",
		220,
		"Single cylinder",
	}
	bike2 := bike{
		"Honda CBR 650f",
		"Sports Tourer",
		650,
		"Four cylinder",
	}
	bike3 := bike{
		"Honda Africa Twin",
		"Adventure Sports",
		1000,
		"Two cylinder",
	}
	bike4 := bike{
		"Kawasaki z900",
		"Naked Sports",
		900,
		"Four cylinder",
	}
	bike1.printInfo()
	bike2.printInfo()
	bike3.printInfo()
	bike4.printInfo()

}