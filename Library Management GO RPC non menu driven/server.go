package main 
import(
	"log"
	"errors"
	"net"
	"net/rpc"
)
type Book struct{
	BookID int 
	Title string 
	Available bool
}
type BookArgs struct{
	BookID int
}
type Library struct{
	Books map[int]Book
}
func (l* Library) Add(b *Book,reply *string) error{
	b.Available=true 
	l.Books[b.BookID]=*b 
	*reply="Book Added Successfully"
	return nil
}
func (l* Library) Issue(b *BookArgs,reply *string) error{
	b1,ok:=l.Books[b.BookID]
	if !ok{
		return errors.New("Book Not Found")
	}
	if !b1.Available{
		return errors.New("Book Not Available")
	}
	b1.Available=false 
	l.Books[b.BookID]=b1
	*reply="Book Issued Successfully"
	return nil
}
func (l *Library) Return(b *BookArgs,reply *string) error{
	b1,ok:=l.Books[b.BookID]
	if !ok{
		return errors.New("Book Not Found")
	}
	if b1.Available{
		return errors.New("Book Already Available")
	}
	b1.Available=true 
	l.Books[b.BookID]=b1 
	*reply="Book Returned Successfully"
	return nil
}
func (l *Library) CheckAvailability(b *BookArgs,reply *bool) error{
	b1,ok:=l.Books[b.BookID]
	if !ok{
		return errors.New("Book Not Found")
	}
	*reply=b1.Available
	return nil
}
func main(){
	library:=&Library{Books:make(map[int]Book)}
	if err:=rpc.Register(library);err!=nil{
		log.Fatal(err)
	}
	ln,err:=net.Listen("tcp",":1234")
	if err!=nil{
		log.Fatal(err)
	}
	log.Println("Library GO RPC listening on :1234")
	for{
		conn,err:=ln.Accept()
		if err!=nil{
			continue
		}
		go rpc.ServeConn(conn)
	}
}