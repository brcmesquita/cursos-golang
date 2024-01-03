package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
}

// Criando um método
func (u usuario) salvar() {
	//fmt.Println("Dentro do método salvar")
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade(idade int) {
	if idade <= 17 {
		fmt.Println("O usuário é menor de idade")
	} else {
		fmt.Println("O usuário é maior de idade")
	}
}

// função fazer Aniversário (eu fazendo)
func (u usuario) fazerAniversario(idade int) {
	novaIdade := idade + 1
	fmt.Printf("Feliz aniversário! Você tinha %d, e agora tem %d! Parabéns!\n", idade, novaIdade )
}

// Função fazerAniversário (professor fazendo)
func (u *usuario) fazerAniversario2() {
	u.idade++
}



func main() {
	usuario1 := usuario{"Usuário 1", 16}
	//fmt.Println(usuario1)
	usuario1.salvar()
	usuario1.maiorDeIdade(int(usuario1.idade))
	usuario1.fazerAniversario(int(usuario1.idade))
	usuario1.fazerAniversario2()
	fmt.Println(usuario1.idade)

	usuario2 := usuario{"Usuario 2", 30}
	usuario2.salvar()
	usuario2.maiorDeIdade(int(usuario2.idade))
	usuario2.fazerAniversario(int(usuario2.idade))
	usuario2.fazerAniversario2()
	fmt.Println(usuario2.idade)
}