# golang

GO
Next Gen Language

Compiled: 
Go Lang is compiled
Go tool can run file directly, without VM
Executables are different for OS

What:
System apps to web apps- cloud

Object Oriented:
 Yes and No
What you see on the screen is the code

-In go lang 
We don’t have classes in go-lang, we have structs which king of alternative of classes
Don’t have overloading overriding in go-lang

Missing:
 Is it really missing?
No try Catch
lexer does a lot of work

--------------------------------------------------------------------------------------------------------------------------------
Golang Installation and hello world

create a file with name main.go
	1. go mod init  hello 
	2. func main() is the starting point of the program
	3. Main. Go file :
	package main
	import "fmt"
	func main() {
	    fmt.Println("Hello World !!!!!!!")
	}
	4. go  run main.go
--------------------------------------------------------------------------------------------------------------------------------

GOPATH and reading go docs

go help path: this will you give you the documentation how we can set go path

go env GOPATH: this will give you the current go path


--------------------------------------------------------------------------------------------------------------------------------
Lexer in golang and Types

Types:
Case insensitive; almost;
Variable type should be known in advance
Everything is a Type
 

	


--------------------------------------------------------------------------------------------------------------------------------

Variables, type and constants

package main
import "fmt"
func main() {
    fmt.Println("Variables !!!!!")
    var userName string = "golang"
    fmt.Println(userName)
    fmt.Printf("Usename is of type %T \n", userName)
}

For Boolean:

Var isBoolean bool =True



// defaultValues
Ex: int : 0
We don’t have garbage  values in golang ,some default values are assigned

//implicit type
If you will not define variable type, it will assign based upon the values assigned
Ex:
var website ="www.google.com" 
 fmt.Println(website)

But later on you can't change it value.

// no var type
numberOfUsers := 3000  (This declaration is only allowed inside method)
fmt.Println(numberOfUsers)

// constants
const LoginToken string ="asdksadfnkjsdf"  
(here variable name starts with Capital letter indicates this variable type is Public)


--------------------------------------------------------------------------------------------------------------------------------
Comma ok syntax and package in golang

Taking Input from User

You can use OS package or Bufio for taking input from user. 

Documentation we can find it here:
https://pkg.go.dev/bufio
https://pkg.go.dev/os


package main
import (
    "bufio"
    "fmt"
    "os"
)
func main() {
    welcome := "Welcome to user Input"
    fmt.Println(welcome)
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter User Input: ")
    // comma ok // err err
    input, _ := reader.ReadString('\n')
    fmt.Println("Thanks for input,", input)
    fmt.Printf("Type of this rating is %T", input)
}

Here if we give input as integer , it will still go as string so this is problem for further operations.
Here input, _ act in such a way we only care about input not about error, vice versa for  _, input
input, err this will act as a try catch block, if input is fine then ok, else it will throw error
--------------------------------------------------------------------------------------------------------------------------------








![image](https://github.com/jatinpardhi/golang/assets/48383934/ed082056-af0d-49ea-86ca-e902262ca7dc)
