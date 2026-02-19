package main 
import(
	"log"
	"net/rpc"
	"time"
	"fmt"
)
type Args struct{
	A int
	B int 
}
func main(){
	client,err:=rpc.Dial("tcp",":1234")
	if err!=nil{
		log.Fatal(err)
	}
	defer client.Close()
	reply:=new(int)
	args:=&Args{A:10,B:10}
	call:=client.Go("Arith.Add",args,reply,nil)
	time.Sleep(300*time.Millisecond)
	done:=<-call.Done 
	if done.Error!=nil{
		log.Fatal(done.Error)
	}
	fmt.Println("Add:",*reply)
}