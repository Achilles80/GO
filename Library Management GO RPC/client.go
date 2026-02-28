package main 
import (
	"log"
	"fmt"
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
func main(){
	client,err:=rpc.Dial("tcp","localhost:1234")
	if err!=nil{
		log.Fatal(err)
	}
	defer client.Close()
	newBook:=&Book{BookID:101,Title:"Hello World"}
	var msgReply string 
	var boolReply bool 
	client.Call("Library.Add",newBook,&msgReply)
	fmt.Println("Add: ",msgReply)
	args:=&BookArgs{BookID:101}
	client.Call("Library.CheckAvailability",args,&boolReply)
	fmt.Println("CheckAvailability: ",boolReply)
	client.Call("Library.Issue",args,&msgReply)
	fmt.Println("Issue",msgReply)
	client.Call("Library.CheckAvailability",args,&boolReply)
	fmt.Println("CheckAvailability: ",boolReply)
	client.Call("Library.Return",args,&msgReply)
	fmt.Println("Return",msgReply)
	client.Call("Library.CheckAvailability",args,&boolReply)
	fmt.Println("CheckAvailability: ",boolReply)
}