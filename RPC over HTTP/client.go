package main 
import(
	"fmt"
	"log"
	"net/rpc"
)
type Args struct{
	A int 
	B int 
}
func main(){
	client,err:=rpc.DialHTTP("tcp",":8080")
	if err!=nil{
		log.Fatal(err)
	}
	defer client.Close()
	args:=&Args{A:20,B:30}
	var out int 
	client.Call("Arith.Add",args,&out); fmt.Println(out)
	client.Call("Arith.Mul",args,&out); fmt.Println(out)
}