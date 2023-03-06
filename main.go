package main 

import ("fmt")
type person struct{
	firstName string
	secondName string 
	fis []string
}
func main(){
	ej2()
}

func ej1(){
	person1 := person{
		firstName: "Johan",
		secondName: "Josue",
		fis: []string{"Chocolate","Mint","Vanilla"},
	}
	person2 := person{
		firstName: "J",
		secondName: "Cole",
		fis: []string{"Red","Blue","Grey"},
	}
	fmt.Println(person1)
	fmt.Println(person2)
	for i,k := range person1.fis{
		fmt.Printf("The position is:%v \t Name of the Ice Cream:%v\n",i,k)
	}
}

func ej2(){
	person1 := person{
		firstName: "Johan",
		secondName: "Josue",
		fis: []string{"Chocolate","Mint","Vanilla"},
	}
	person2 := person{
		firstName: "J",
		secondName: "Cole",
		fis: []string{"Red","Blue","Grey"},
	}

	m := map[string]person{
		person1.firstName:person1,
		person2.firstName:person2,
	}
	fmt.Println(m)
	fmt.Println("Now in the foor loop")
	for i,k := range m{
		fmt.Printf("The person is:%v\n",k)
		for x, z:= range i{
			fmt.Printf("This is the value:%v\t This is the position:%v\n ",z,x)
		}
	}
}