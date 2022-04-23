package main

import(
	"fmt"
)

func myFunc(ch chan string,str,str1 string){
	for _,v:=range str{
		ch <- string(v)
	}
	a := func(ch chan string,str1 string){
		for _,v:=range str1{
			ch <- string(v)
		}
		close(ch)
	}

	a(ch,str1)

}


func main()  {
	ch := make(chan string)
	str := "Gaurav"
	str1 := "Sahil"
	go myFunc(ch,str,str1)
	for i := range ch{
		fmt.Println(i)
	}
}
