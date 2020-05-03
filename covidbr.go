package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
var resposta string

func main (){
var estado string
fmt.Print("Informe a sigla do estado: \n  Ex: MG - Minas Gerais \n Sigla: ")
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
	fmt.Println("\n Deseja consultar outro estado ? SIM(S) NÃO(N) \t ")
	fmt.Scanln(&resposta)
	if resposta == "S"{
		main()
	} else {}
}
