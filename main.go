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
func main3() {
	var hello string  = "Olá mundo"
	var question string = "Como vai?"

	//concatenacao
	var meet = hello + question
	fmt.Println(meet)
	fmt.Println(strings.ToUpper(meet)) //pacote strings para fazer operaçoes basicas nos textos
	fmt.Println(strings.Contains(meet, "mundo"))	
}


//array
func main5() {
	var gavetas [2]string //damos qual o tamanho do array
	gavetas[0] = "copos" // adicionamos o valor e qual o indice que ele ocupa
	gavetas[1] = "panos" //	nao e possivel adicionar mais itens do que o array permite, dara erro de compilacao

	fmt.Println(gavetas[1]) // para ler o array eu preciso informar qual o indice que quero ler 
}

//slices, tem tamanhos flexiveis para armazenar itens do mesmo tipo 
func main5() {
	var gavetinhas []string // Aqui eu gero o slice
	gavetinhas = append(gavetinhas, "copinhos", "paninhos") // Para popular um slice é preciso utilizar a função append
	fmt.Println(gavetinhas)
	fmt.Println(len(gavetinhas)) // Podemos validar qual o tamanho do slice gavetinhas
	fmt.Println(gavetinhas[0]) // Para acessar o slice, também é por meio de indices.
	fmt.Println(gavetinhas[0:1]) // Para fazer uma divisão do valor final do slice precisamos informar qual o indice que se inicia, qual que termina SEM CONTAR. Se eu não passar o valor de indice inicial, ele vai pegar o na posição de 0. Se eu não passo algo para o final, ele considera o último indice
}

//Maps sao estrutura chave valor com tipos na chave e no valor para inserção rápida
func main6() {
	var pessoinhas = map[string]int{} /// Aqui criamos a variavel e informamos qual para o map qual será o tipo da chave (string) e qual será o tipo do valor {int}
	pessoinhas ["lais"] = 26 //inserção do valor
	pessoinhas ["leo"] = 32
	fmt.Println(pessoinhas)
	fmt.Println(pessoinhas["lais"]) //podemos fazer ref a estrutura do map para conseguir acessar o valor. Caso eu queira validar antes de acessar, posso fazer uma validação com SE 
	if  idade, ok := pessoinhas ["lais"]; ok { // a informação básica é "Se a idade for OK e a pessoa  =  a LAIS for OK, faça isso, se não faça aquilo."
		fmt.Println("Pessoa existe no map", idade, ok)
	} else {
		fmt.Println("Pessoa não existe no map")
	}
}

/// fluxo de controle  (if/else/switch/for/range)

func main () {
	nota := 75

	if nota >= 90 {
		fmt.Println("Aprovado com distinção")
	} else if  nota >= 70 {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")
	}
}