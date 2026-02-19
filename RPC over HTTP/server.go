package main
import(
	"log"
	"net"
	"net/rpc"
	"net/http"
)
type Args struct{
	A int 
	B int 
}
type Arith struct{}
func (a *Arith) Add(args *Args,reply *int) error {*reply=args.A+args.B; return nil}
func (a *Arith) Mul(args *Args,reply *int) error {*reply=args.A*args.B; return nil}

func main(){
	if err:=rpc.Register(new(Arith)); err!=nil{
		log.Fatal(err)
	}
	rpc.HandleHTTP()
	ln,err:=net.Listen("tcp",":8080")
	if err!=nil{
		log.Fatal(err)
	}
	log.Println("HTTP RPC Server on :8080")
	log.Fatal(http.Serve(ln,nil))
}