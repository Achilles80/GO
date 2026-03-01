package main 
import (
	"log"
	"fmt"
	"net/rpc"
	"bufio"
	"os"
	"strings"
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
	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("Enter book id: ")
	var bookid int 
	fmt.Scanln(&bookid)
	fmt.Println("Enter book title: ")
	title,_:=reader.ReadString('\n')
	title=strings.TrimSpace(title)
	newBook:=&Book{BookID:bookid,Title:title}
	var msgReply string 
	var boolReply bool 
	client.Call("Library.Add",newBook,&msgReply)
	fmt.Println("Add: ",msgReply)
	args:=&BookArgs{BookID:bookid}
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