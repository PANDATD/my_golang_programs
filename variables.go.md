# Variables in go 

+ As we know go languge is _statically_ typed and _compiled_ languge. 
+ So, we can not changed the datatypes as we will do in python programming languge.
  ### Python code example of Dynamic type changing 
  - ```python3 
      >>> name = "Tejas" # name is now string format 
      >>> name 
      'Tejas'
      >>> type(name)
      <class 'str'>
      >>> name = 20 #name is now int format 
      >>> name 
      20
      >>> type(20)
      <class 'int'>
      ```
   - We Done in following example beacuse python is Dynamically typed languge.
 + we can not change type like following beacuse go dose not 
   support _Dynamic Typing_ 
   
### Now we will learn declaretion of variable in Go
 ```go 
      package main 
      import "fmt"

      func main () {
      // First Way :- 
      var name string = "Tejas Dixit" 
      fmt.Println("Name: ", name)
      
      //Note :- Do not  RE-DECALARE VARIABLE 
     
      /* 
      I just puteed in comment else it gives you ERROR message  like command-line-arguments
          ./variable.go:8:9: name redeclared in this block
	        previous declaration at ./variable.go:5:8
        ./variable.go:12:10: no new variables on left side of :=
      */

      /* 
      // Second way:- ( As we Done in c/c++ languge [little be same] ) 
       var name string 
       name = "Vignesh" 
       fmt.Println("Name: ", name) 
       
      // Third way :-  (Using walrus opretor)
       name := "Kunal" 
       fmt.Println("Nmae: ", name) 
       
      //Fourth way :- (Normal method)
       name = "Datta" 
       fmt.Println("Name: ", name)
      */
    
     } 
     
   ``` 
### Run Program using following commond
```bash 
 go run variable.go 
``` 
### Output
``` 
Name:  Tejas Dixit
``` 

