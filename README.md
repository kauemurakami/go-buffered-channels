[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-buffered-channels/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-buffered-channels/blob/main/README.md)

## Buffered Channels
A little different from a ```chan```, we also specify a capacity for this channel and this is very important.  
Example:  
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
Does it seem to be correct? the result of this is a *deadlock*, because the option of receiving and sending data are blocking changes, so we are sending the value to the channel, my program will wait for another line to receive the value from this channel, but this never happens in this code, it is blocked in ```channel <- "Hello world"``` and never reaches ```message := <-channel```, so it results in a deadlock, we send a value but there is no one to receive that value .  
Precisely for this reason, we use channels in separate functions, an alternative to being able to continue in the same function, which is to create a channel with ```buffer```, specify a capacity for ```chan```, changing only the line ```channel``` statement:  
```go
...
  channel := make(chan string, 2) // specifying the channel buffer/size
  channel <- "Hello world"
  message := <-channel
  fmt.Println(message) // output Hello World
```
This time it worked, the difference between a buffered channel and a normal one is that a buffered channel will only block when it reaches its maximum capacity, when it sends a value to it, as it has a capacity of 2 in this case it doesn't it will block, the program will continue, we are sending a value to it but I don't necessarily need to wait for someone to receive it, as we are working on a ```chan``` with ```buffer```.  
### Exemplo de uso correto e errado
Since our buffer in the ```channel := make(chan string, 2)``` example has a capacity of two, we can do this:
```go
func main() {
	channel := make(chan string, 2)

	channel <- "Hello world"
	channel <- "Olá mundo"

	message := <-channel
	fmt.Println(message)// output Hello world
}
```
So we can assign a value to the buffer twice, and this will work, as it has a ```buffer``` of 2, it prints once as it only receives the value of our ```chan``` once.<br /><br/>

Now, if we tried to pass more than 2 values ​​to our ```chan```:  
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
Because he sent it twice and reached its maximum capacity, so the next send, it hangs after the limit line and will never reach ```message <- channel```  

*Receiving and showing more than one value with buffer*  
```go
func main() {
	channel := make(chan string, 2)

	channel <- "Hello world"
	channel <- "Olá mundo"
	// channel <- "GO lang" //error deadlock

	message := <-channel
	message2 := <-channel
	fmt.Println(message) // output Hello world
	fmt.Println(message2) // Olá mundo
}
```

