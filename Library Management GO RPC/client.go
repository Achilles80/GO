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
	for{
		fmt.Println("Enter what to do: ")
		fmt.Println("1. Add Book")
		fmt.Println("2.Check Availability Book")
		fmt.Println("3. Issue Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. Exit")
		var opt int 
		fmt.Scanln(&opt)
		var msgReply string 
		var boolReply bool 
		var bookid int 
		reader:=bufio.NewReader(os.Stdin)
		if opt==1{
			fmt.Println("Enter book id: ")
			fmt.Scanln(&bookid)
			fmt.Println("Enter book title: ")
			title,_:=reader.ReadString('\n')
			title=strings.TrimSpace(title)
			newBook:=&Book{BookID:bookid,Title:title}
			client.Call("Library.Add",newBook,&msgReply)
			fmt.Println("Add: ",msgReply)
		}else if opt==2{
			fmt.Println("Enter book id: ")
			fmt.Scanln(&bookid)
			args:=&BookArgs{BookID:bookid}
			client.Call("Library.CheckAvailability",args,&boolReply)
			fmt.Println("CheckAvailability: ",boolReply)
		}else if opt==3{
			fmt.Println("Enter book id: ")
			fmt.Scanln(&bookid)
			args:=&BookArgs{BookID:bookid}
			client.Call("Library.Issue",args,&msgReply)
			fmt.Println("Issue",msgReply)
		}else if opt==4{
			fmt.Println("Enter book id: ")
			fmt.Scanln(&bookid)
			args:=&BookArgs{BookID:bookid}
			client.Call("Library.Return",args,&msgReply)
			fmt.Println("Return",msgReply)

		}else{
			break
		}
	}		
}