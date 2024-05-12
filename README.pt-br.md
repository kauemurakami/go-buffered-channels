[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-buffered-channels/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-buffered-channels/blob/main/README.md)

## Canais com Buffer
Um pouco diferente de um ```chan```, a gente especifica também uma capacidade para este canal e isso é muito importante.  
Exemplo:  
```go
package main

import "fmt"

func main() {
	channel := make(chan string)
	channel <- "Hello world"
	message := <-channel
	fmt.Println(message)
}
```  
Parece estar correto? o resultado disso é um *deadlock*, pois a opeção de receber e enviar dados, são alterações bloqueantes, então estamos enviando o valor pro canal, meu programa vai esperar alguma outro linha receber o valor desse canal, mas isso nunca acontece nesse código, ele fica bloqueado em ```channel <- "Hello world"``` e nunca chega em ```message := <-channel```, então resulta em deadlock, enviamos um valor mas não tem ninguém pra receber esse valor.  
Justamente por isso, usamos os canais em funções separadas, uma alternativa para poder continuar na mesma função, que é criar um canal com ```buffer```, especificar uma capacidade pro ```chan```, alterando apenas a linha de de claração do ```channel```:  
```go
...
  channel := make(chan string, 2) // especificando o buffer/tamanho do canal
  channel <- "Hello world"
  message := <-channel
  fmt.Println(message) // output Hello World
```
Desta vez funcionou, a diferença de um canal com buffer e um normal, é que um canal com buffer só vai bloquear quando atingir a capacidade máxima dele, quando ele envia um valor pra ele, como ele tem uma capacidade de 2 nesse caso ele não vai bloquear, vai continuar o programa, estamos enviando um valor pra ele mas não necessariamente preciso esperara alguem receber, pois estamos trabalhando em um ```chan``` com ```buffer```.  

### Exemplo de uso correto e errado
Como nosso buffer no exemplo ```channel := make(chan string, 2)``` tem capacidade de dois, podemos fazer isso:  
```go
func main() {
	channel := make(chan string, 2)

	channel <- "Hello world"
	channel <- "Olá mundo"

	message := <-channel
	fmt.Println(message)// output Hello world
}
```
Então podemos atribuir um valor no buffer duas vezes, e isso funcionará, pois possui o ```buffer``` de 2, ele printa uma vez pois só recebe uma vez o valor de nosso ```chan```.<br/><br/>

Agora, caso tentássemos passar mais de 2 valores para nosso ```chan```:  
```go
func main() {
	channel := make(chan string, 2)

	channel <- "Hello world"
	channel <- "Olá mundo"
	channel <- "Go lang"

	message := <-channel
	fmt.Println(message)// output deadlock
}
```
Pois ele enviou duas vezes e chegou na capacidade máxima dele, portanto o proximo envio, ele trava após a linha limite e nunca chegará em ```message <- channel```  

*Recebendo e mostrando mais de um valor com buffer*  
```go
func main() {
	channel := make(chan string, 2)

	channel <- "Hello world"
	channel <- "Olá mundo"
	// channel <- "GO lang" //erro deadlock

	message := <-channel
	message2 := <-channel
	fmt.Println(message) // output Hello world
	fmt.Println(message2) // Olá mundo
}
```

