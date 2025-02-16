package main

import (
	//"fmt" //Pacote padrão do Go

	"main.go/meet"
)

func main () { //a função main nunca retorna nada por padrão mas ela sempre precisa ser criada para garantir que vamos conseguir rodar o código
	meet.SayHello()/// Letra maiuscula permite utilizar em qualquer lugar, letra minuscula não
	meet.Say("Olá Jules")
}

//Para executar podemos rodar o Go run ou go build -o name para que possamos rodar o binário


// Go compila direto para código de maquina, sem a necessidadade de ter um compilador adicional