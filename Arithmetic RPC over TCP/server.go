package main
import(
	"net"
	"net/rpc"
	"log"
	"errors"
)
type Args struct{
	A int
	B int 
}
type Arith struct{}
func (a *Arith) Add(args *Args,reply *int) error{*reply=args.A+args.B; return nil}
func (a *Arith) Sub(args *Args,reply *int) error{*reply=args.A-args.B; return nil}
func (a *Arith) Mul(args *Args,reply *int) error{*reply=args.A*args.B; return nil}
func (a *Arith) Div(args *Args,reply *int) error {
	if args.B==0{
		return errors.New("Division by zero")
	}
	*reply=args.A/args.B 
	return nil
}
func main(){
	if err:=rpc.Register(new(Arith)); err!=nil {
		log.Fatal(err)
	}
	ln,err :=net.Listen("tcp",":1234")
	if err!=nil{
		log.Fatal(err)
	}
	log.Println("TCP RPC on server :1234")
	for{
		conn,err:=ln.Accept()
		if err!=nil{
			log.Fatal(err)
		}
		go rpc.ServeConn(conn)
	}
}