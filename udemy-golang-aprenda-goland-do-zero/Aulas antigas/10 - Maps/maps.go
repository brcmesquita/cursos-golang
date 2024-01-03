package main

import "fmt"

func main(){
	fmt.Println("Maps")

	usuario := map[string]string{ // precisa respeitar o tipo, senão não compila
		"nome": "Bruno",
		"Sobrenome": "Raphael",
	}

	fmt.Println(usuario) // map[Sobrenome:Raphael nome:Bruno]
	fmt.Println(usuario["nome"]) // Bruno

	// modo 2 - gambiarra para aninhar maps
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Bruno",
			"ultimo": "Mesquita",
		},
		"curso": {
			"nome": "Sistemas de Informação",
			"faculdade": "Estácio de Sá",
		},
	}

	fmt.Println(usuario2) // map[curso:map[faculdade:Estácio de Sá nome:Sistemas de Informação] nome:map[primeiro:Bruno ultimo:Mesquita]]

	// Se quiser deletar uma chave
	delete(usuario2, "nome")
	fmt.Println(usuario2) // map[curso:map[faculdade:Estácio de Sá nome:Sistemas de Informação]]

	// Se quiser adicionar uma chave
	usuario2["signo"] = map[string]string {
		"nome": "Câncer", // não precisa adicionar todos os campos do map neste caso
	}

	fmt.Println(usuario2) // map[curso:map[faculdade:Estácio de Sá nome:Sistemas de Informação] signo:map[nome:Câncer]]
}