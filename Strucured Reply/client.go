package main 
import(
	"net/rpc"
	"log"
	"fmt"
)
type Args struct{
	A int 
	B int 
}
type Reply struct{
	Sum int 
	Product int 
}
func main(){
	client,err:=rpc.Dial("tcp",":1234")
	if err!=nil{
		log.Fatal(err)
	}
	defer client.Close()
	args:=&Args{A:10,B:20}
	var reply Reply
	if err:=client.Call("Arith.SumProduct",args,&reply); err!=nil {
		log.Fatal(err)
	}
	fmt.Println("Sum: ",reply.Sum)
	fmt.Println("Product: ",reply.Product)
}