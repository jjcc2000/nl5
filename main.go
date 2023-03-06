package main 

import ("fmt")
type person struct{
	firstName string
	secondName string 
	fis []string
}

func main(){
	ej4()
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
	//Here you got a  Map that has a Struct named person and inside that struct there is a slice
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
	// fmt.Println(m)
	// fmt.Println("Now in the foor loop")
	for i,k := range m{
		fmt.Printf("The person is:%v\n",i)
		for x, z:= range k.fis{
			fmt.Printf(" This is the position:%v\t This is the value:%v\n ",x,z)
		}
	}
}
func ej3(){
	//using a compose literal create truck and sedan
	type vehicle struct{
		doors int
		color string 
	}
	type sedan struct{
		 v vehicle
		 luxury bool
	}
	type truck struct{
		v vehicle
		fourWheel bool
	}
	corolla := sedan{
		v: vehicle{4, "Red"},
		luxury: true,
	}
	frontier:= truck{
		v: vehicle{4, "Grey"},
		fourWheel: true,
	}
	fmt.Println(corolla)
	fmt.Println(frontier)

}
func ej4(){
	s := struct{
		name string
		friends map[string]string
		Colors []string
	}{
		name: "Johan",
		friends:map[string]string{
				"Sergio": "Es su mejor amigo",
				"Pedro": "No es su mejor amigo",
		} ,
		Colors: []string{"Red","Blue","Grey"},
	}
	fmt.Println("Name:",s.name)
	for k, row := range s.friends{
		fmt.Printf("Key:%v\tContent:%v\n",k,row)
	}
	for i,k := range s.Colors{
		fmt.Println("Position Index:",i, "Content inside Index:",k)
	}
}