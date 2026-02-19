package main 
import(
	"log"
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)
type Args struct{
	A int 
	B int 
}
func main(){
	conn,err:=net.Dial("tcp",":1234")
	if err!=nil{
		log.Fatal(err)
	}
	defer conn.Close()
	client:=jsonrpc.NewClient(conn)
	args:=&Args{A:20,B:20}
	var out int 
	if err:=client.Call("Arith.Add",args,&out); err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Add: ",out)
}