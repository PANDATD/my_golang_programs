# Loop in Golang 
---
```go
package main 
import "fmt"

// Main func to run code 
func main(){
        //For loop 
        for i:=1; i<100;i++{
                // statement to execute repitedly
                fmt.Println("Hello World! ",i)
                }
          }
```
---

### Note

- In Go programming languge their is only one loop<br>
  such as *for loop*<br>
  if we have to exceuter a block of code repetedly then we will use for loop.<br>
  along with for keyword initalze varable put semicolon put condition again semicolon;<br>
  and increment or decrement.<br>
  Why we use semicolon ? -> To sepretout the conditions.<br>

### SYNTAX:

        for initalization; condition; increment/decrement{
        statement to execute repitedliy
        }

### EXAMPLE:
```go
        for i := 1; i<=10; i++{
        fmt.Println("Hello World! ",i)
        }
```
