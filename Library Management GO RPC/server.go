package main 
import(
	"log"
	"net"
	"net/rpc"
	"errors"
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
func (l *Library) Add(book *Book,reply *string) error{
	book.Available=true 
	l.Books[book.BookID]=*book 
	*reply="Book Successfully Added"
	return nil 
}
func (l *Library) Issue(args *BookArgs, reply *string) error{
	b,ok:=l.Books[args.BookID]
	if !ok{
		return errors.New("Book not found")
	}
	if !b.Available{
		return errors.New("Book not available")
	}
	b.Available=false 
	l.Books[args.BookID]=b
	*reply="Book Successfully issued"
	return nil 
}
func (l *Library) Return(args *BookArgs,reply *string) error{
	b,ok:=l.Books[args.BookID]
	if !ok{
		return errors.New("Book not found")
	}
	b.Available=true 
	l.Books[args.BookID]=b
	*reply="Book Successfully returned"
	return nil
}
func (l *Library) CheckAvailability(args *BookArgs,reply *bool) error{
	b,ok:=l.Books[args.BookID]
	if !ok{
		return errors.New("Book not found")
	}
	*reply=b.Available
	return nil
}
func main(){
	lib:=&Library{Books:make(map[int]Book)}
	if err:=rpc.Register(lib); err!=nil{
		log.Fatal(err)
	}
	ln,err:=net.Listen("tcp",":1234")
	if err!=nil{
		log.Fatal(err)
	}
	log.Println("Library RPC on TCP :1234")
	for{
		conn,err:=ln.Accept()
		if err!=nil{
			log.Fatal(err)
		}
		go rpc.ServeConn(conn)
	}
}