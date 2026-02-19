package main 
import(
	"fmt"
	"net/rpc"
	"log"
)
type Args struct{
	A int 
	B int 
}
func main(){
	client,err := rpc.Dial("tcp","localhost:1234")
	if err!=nil{
		log.Fatal(err)
	}
	defer client.Close()
	args := &Args{A:20, B:0}
	var out int 
	client.Call("Arith.Add",args,&out); fmt.Println("Add: ",out)
	client.Call("Arith.Sub",args,&out); fmt.Println("Sub: ",out)
	client.Call("Arith.Mul",args,&out); fmt.Println("MUl: ",out)
	if err:=client.Call("Arith.Div",args,&out); err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Div: ",out)
}