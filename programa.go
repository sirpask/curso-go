package main 

import ("fmt"
        "os"
        "strconv"
       )

func main (){
	fmt.Println("***** MI PROGRAMA CON GO *******")

	/*

     fmt.Println("Hola "+os.Args[1]+" bienvenido al programa en GO")

	edad,_ := strconv.Atoi(os.Args[2])


	if edad > 18 && edad !=25 && edad != 99 {
		fmt.Println("Eres mayor de edad")
	}else if edad == 25 {
		fmt.Println("Tienes 25 años")
	}else if edad == 99 {
		fmt.Println("pronto moriras")
	}else{
		fmt.Println("eres menor de edad")
	}
	
	*/
	numero,_ := strconv.Atoi(os.Args[1])

	if numero%2 == 0 {
		fmt.Println("El numero es par")

	}else{
		fmt.Println("El numero es impar")

	}
} 