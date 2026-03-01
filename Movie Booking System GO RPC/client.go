package main 
import (
	"fmt"
	"net/rpc"
	"log"
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
func main(){
	client,err:=rpc.Dial("tcp","localhost:1234")
	if err!=nil{
		log.Fatal(err)
	}
	defer client.Close()
	var reply string 
	var reply1 int 
	var reply2 []Movie
	args:=&Movie{MovieID:101,Title:"Game of thrones",AvailableSeats:10}
	if err:=client.Call("System.Add",args,&reply);err!=nil{
		log.Fatal(err)
	}
	fmt.Println(reply)
	args2:=&MovieArgs{MovieID:101}
	if err:=client.Call("System.CheckAvailability",args2,&reply1);err!=nil{
		log.Fatal(err)
	}
	fmt.Println(reply1)
	args1:=&BookArgs{MovieID:101,TicketsToBook:10}
	if err:=client.Call("System.BookTickets",args1,&reply);err!=nil{
		log.Fatal(err)
	}
	fmt.Println(reply)
	if err:=client.Call("System.CheckAvailability",args2,&reply1);err!=nil{
		log.Fatal(err)
	}
	fmt.Println(reply1)
	if err:=client.Call("System.List",args2,&reply2);err!=nil{
		log.Fatal(err)
	}
	fmt.Println(reply2)
}