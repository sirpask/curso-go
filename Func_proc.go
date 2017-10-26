package main 

import ("fmt"
       /* "time"*/
       )

func main (){
	/*var suma int = 8 + 9
	var resta int = 6 - 4
	var nombre string = "david"

	var apellidos string = "David WEB"
	var apellidos = "Pascual"

	pais:= "España"

	fmt.Println("Hola mundo desde Go con " + nombre + apellidos + pais)
	fmt.Println(suma)
	fmt.Println(resta

	time.Sleep(time.Second * 5)*/

	fmt.Println("Hola MUndo desde GO con Pask");
	pantalon("rojo", "largo", "sin bolsillos", "nike");
} 

func pantalon(caracteristicas ...string){
  	for _,caracteristica := range caracteristicas{
  	fmt.Println(caracteristica);
  }
};