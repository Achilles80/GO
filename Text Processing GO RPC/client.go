package main 
import(
	"fmt"
	"net/rpc"
	"log"
)
type TextArgs struct{
	Text string
}
type CipherArgs struct{
	Text string
	Shift int
}
func main(){
	client,err:=rpc.Dial("tcp","localhost:1234")
	if err!=nil{
		log.Fatal(err)
	}
	defer client.Close()
	var srep string 
	var num int 
	args:=&TextArgs{Text:"hello"}
	client.Call("Tproc.Reverse",args,&srep)
	fmt.Println(srep)
	client.Call("Tproc.Count",args,&num)
	fmt.Println(num)
	args1:=&CipherArgs{Text:"abcd",Shift:1}
	client.Call("Tproc.Caeser",args1,&srep)
	fmt.Println(srep)
	args2:=&TextArgs{Text:"Hello World"}
	client.Call("Tproc.Format",args2,&srep)
	fmt.Println(srep)
}