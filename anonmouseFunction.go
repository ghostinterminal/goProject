// The usage of anonymouse function using golang.


package main
import "fmt"

func main(){

  square := func (n int) (int){
                  return n*n
                  }
  x := 1
  for x<10{
  
    fmt.Println(square(x))
    x++                 
           }
}
