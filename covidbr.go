package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
var resposta string
var password = "123456"

func main (){
var estado string
fmt.Print("Informe a sigla do estado: \n  Ex: MG - Minas Gerais \n Sigla: ")
fmt.Print(password)
fmt.Scanln(&estado)
buscaapi(estado)
}

func buscaapi(estado string) {	
	var url = "https://covid19-brazil-api.now.sh/api/report/v1/brazil/uf/" //api usada para consulta

	response, _:= http.Get(url+estado)

	// Leitura do Json
	responseJson, _ := ioutil.ReadAll(response.Body)

	// resultado
	fmt.Println(estado)
	fmt.Printf(string(responseJson))
	fmt.Println("\n Deseja consultar outro estado ? SIM(S) NÃO(N) ")
	fmt.Scanln(&resposta)
	if resposta == "S"{
		main()
	} else if resposta == "s"{
		main ()
	} else{}
	//fmt.Println("\n Evolução dos", len(dadosJson), "dias desde a primeira contaminação!")
}
