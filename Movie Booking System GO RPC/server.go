package main 
import(
	"log"
	"errors"
	"net"
	"net/rpc"
)
type Movie struct{
	MovieID int 
	Title string 
	AvailableSeats int
}
type MovieArgs struct{
	MovieID int 
}
type BookArgs struct{
	MovieID int 
	TicketsToBook int 
}
type System struct{
	Movies map[int]Movie
}
func (s *System) Add(m *Movie,reply *string) error{
	s.Movies[m.MovieID]=*m 
	*reply="Movie added Successfully"
	return nil
}
func (s *System) BookTickets(b *BookArgs,reply *string) error{
	m,ok:=s.Movies[b.MovieID]
	if !ok{
		return errors.New("Movie not found")
	}
	if b.TicketsToBook>m.AvailableSeats{
		return errors.New("Insufficient Available Seats")
	}
	m.AvailableSeats-=b.TicketsToBook
	s.Movies[b.MovieID]=m 
	*reply="Tickets booked successfully"
	return nil
}
func (s *System) CheckAvailability(args *MovieArgs,reply *int) error{
	m,ok:=s.Movies[args.MovieID]
	if !ok{
		return errors.New("Movie not found")
	}
	*reply=m.AvailableSeats
	return nil
}
func (s *System) List(args *MovieArgs,reply *[]Movie) error{
	arr:=[]Movie{}
	for _,v := range s.Movies{
		arr=append(arr,v)
	}
	*reply=arr
	return nil 
}
func main(){
	system:=&System{Movies:make(map [int]Movie)}
	if err:=rpc.Register(system);err!=nil{
		log.Fatal(err)
	}
	ln,err:=net.Listen("tcp",":1234")
	if err!=nil{
		log.Fatal(err)
	}
	log.Println("System Listening at :1234")
	for{
		conn,err:=ln.Accept()
		if err!=nil{
			continue
		}
		go rpc.ServeConn(conn)
	}
}