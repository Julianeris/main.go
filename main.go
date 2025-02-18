package main

import (
	//"fmt" //Pacote padrão do Go

	"fmt"
	"strconv"
	"strings"

	"main.go/meet"
)

func main1 () { //a função main nunca retorna nada por padrão mas ela sempre precisa ser criada para garantir que vamos conseguir rodar o código
	meet.SayHello()/// Letra maiuscula permite utilizar em qualquer lugar, letra minuscula não
	meet.Say("Olá Jules")
}

//Para executar podemos rodar o Go run ou go build -o name para que possamos rodar o binário


// Go compila direto para código de maquina, sem a necessidadade de ter um compilador adicional



var texto string = "Olá"; //Podemos criar variaveis de duas formas, a primeira utilizando a palavra reservada var acompanhandada de nome, tipo e valor 

var textos = "Olá"; ///O Go assume que o valor de textos é string, isso acontece também com os demais tipos. 

var textinho string; //Conseguimos deixar setado apenas o nome e tipo da variavel. Com foco em garantir consistência a partir do tipo construido. 

var text = "ola"; // Aqui eu já consigo que a minha variavel nasça já com o tipo esperado, pois estamos setando o valor inicial

//constantes não são alterados em tempo de execução assim como em typescript 

const nome string = "Julia"

// porem é possível fazer uma declaraçao mais curta, permitindo que o Go assuma a informaçao de memoria


/// tipos básicos 
// int
func main2 () {
	var idade int = 30 // aqui eu nao falo o quanto vai ocupar e ele vai assumir que é a quantidade padrao do sistema
	var contador int32 = 2 // informo exatamente quantos bits vai ocupar
}

//bool
func main4 () {
	var novoTipo bool = true // se eu nao mencionar que o valor é true, por padrao ele será false

	fmt.Println(novoTipo)
}

//string 
func main() {
	var hello string  = "Olá mundo"
	var question string = "Como vai?"

	//concatenacao
	var meet = hello + question
	fmt.Println(meet)
	fmt.Println(strings.ToUpper(meet)) //pacote strings para fazer operaçoes basicas nos textos
	fmt.Println(strings.Contains(meet, "mundo"))	
}