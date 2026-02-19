package main 
import(
	"log"
	"net"
	"net/rpc"
)
type Args struct{
	A int 
	B int 
}
type Reply struct{
	Sum int 
	Product int 
}
type Arith struct{}
func (a *Arith) SumProduct(args *Args,reply *Reply) error {
	reply.Sum=args.A+args.B 
	reply.Product=args.A*args.B 
	return nil 
}
func main(){
	if err:=rpc.Register(new(Arith)); err!=nil{
		log.Fatal(err)
	}
	ln,err:=net.Listen("tcp",":1234")
	if err!=nil{
		log.Fatal(err)
	}
	log.Println("TCP RPC Server on :1234")
	for{
		conn,err:=ln.Accept()
		if err!=nil{
			log.Fatal(err)
		}
		go rpc.ServeConn(conn)
	}
}