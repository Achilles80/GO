package main 
import(
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
	var reply string 
	var reply1 bool
	args:=&Book{BookID:101,Title:"Game of Thrones"}
	args1:=&Book{BookID:102,Title:"One Piece"}
	if err:=client.Call("Library.Add",args,&reply);err!=nil{
		fmt.Println(err)
	}
	fmt.Println(reply)
		if err:=client.Call("Library.Add",args1,&reply);err!=nil{
		fmt.Println(err)
	}
	fmt.Println(reply)
	args2:=BookArgs{BookID:102}
	if err:=client.Call("Library.Issue",args2,&reply);err!=nil{
		fmt.Println(err)
	}
	fmt.Println(reply)
	if err:=client.Call("Library.CheckAvailability",args2,&reply1);err!=nil{
		fmt.Println(err)
	}
	fmt.Println(reply1)
		if err:=client.Call("Library.Return",args2,&reply);err!=nil{
		fmt.Println(err)
	}
	fmt.Println(reply)
	if err:=client.Call("Library.CheckAvailability",args2,&reply1);err!=nil{
		fmt.Println(err)
	}
	fmt.Println(reply1)
}