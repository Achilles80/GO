package main 
import(
	"log"
	"net"
	"net/rpc"
	"errors"
	"strings"
)
type TextArgs struct{
	Text string
}
type CipherArgs struct{
	Text string 
	Shift int 
}
type Tproc struct{}
func (t *Tproc) Reverse(s *TextArgs,reply *string) error {
	r:=""
	for i:=len(s.Text)-1;i>=0;i--{
		r=r+string(s.Text[i])
	}
	*reply=r
	return nil
}
func (t *Tproc) Count(s *TextArgs,reply *int) error{
	cnt:=0
	for i:=0;i<len(s.Text);i++{
		if s.Text[i]=='a' || s.Text[i]=='e' || s.Text[i]=='i' || s.Text[i]=='o' || s.Text[i]=='u' || s.Text[i]=='A' || s.Text[i]=='E' || s.Text[i]=='I' || s.Text[i]=='O' || s.Text[i]=='U'{
			cnt++
		}
	}
	*reply=cnt 
	return nil
}
func (t *Tproc) Caeser(s *CipherArgs,reply *string) error{
	rev:=""
	for i:=0;i<len(s.Text);i++{
		rev+= string(s.Text[i]+byte(s.Shift))
	}
	*reply=rev
	return nil
}
func (t *Tproc) Format(s *TextArgs,reply *string) error{
	parts:=strings.Split(s.Text," ")
	if len(parts)<2{
		return errors.New("Invalid name")
	}
	res:=parts[1]+ ", "+parts[0]
	*reply=res
	return nil
}
func main(){
	if err:=rpc.Register(new(Tproc));err!=nil{
		log.Fatal(err)
	}
	ln,err:=net.Listen("tcp",":1234")
	if err!=nil{
		log.Fatal(err)
	}
	log.Println("Text Processing RPC listening on :1234")
	for{
		conn,err:=ln.Accept()
		if err!=nil{
			continue
		}
		go rpc.ServeConn(conn)
	}
}