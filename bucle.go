package main 

import ("fmt"
        
       )

func main (){
	fmt.Println("***** MI PROGRAMA CON GO *******")

	/*

     fmt.Println("Hola "+os.Args[1]+" bienvenido al programa en GO")

	edad,_ := strconv.Atoi(os.Args[2])


	

	}*/

	peliculas := []string{ "el club de la lucha", "cocon", "la vida es bella", "gran torino"}
	
	for i := 1; i<=len(peliculas); i++{
		if i%2 == 0 {
		fmt.Println("El numero es par", i)

	}else{
		fmt.Println("El numero es impar", i)

	}


	}
} 