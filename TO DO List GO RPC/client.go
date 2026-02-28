package main 
import (
	"fmt"
	"log"
	"net/rpc"
)
type Task struct{
	TaskID int
	Description string
	Completed bool
}
type TaskArgs struct{
	TaskID int 
}
func main(){
	client,err:=rpc.Dial("tcp","localhost:1234")
	if err!=nil{
		log.Fatal(err)
	}
	defer client.Close()
	var reply string 
	var reply1 []Task
	args:=&Task{TaskID:101,Description:"Code in GO"}
	if err:=client.Call("TODO.Add",args,&reply);err!=nil{
		log.Fatal(err)
	}
	fmt.Println(reply)
	args1:=&TaskArgs{TaskID:101}
	if err:=client.Call("TODO.ListTasks",args1,&reply1); err!=nil{
		log.Fatal(err)
	}
	fmt.Println(reply1)
	if err:=client.Call("TODO.Complete",args1,&reply); err!=nil{
		log.Fatal(err)
	}
	fmt.Println(reply)
	if err:=client.Call("TODO.ListTasks",args1,&reply1); err!=nil{
		log.Fatal(err)
	}
	fmt.Println(reply1)
	if err:=client.Call("TODO.Delete",args1,&reply); err!=nil{
		log.Fatal(err)
	}
	fmt.Println(reply)
		if err:=client.Call("TODO.ListTasks",args1,&reply1); err!=nil{
		log.Fatal(err)
	}
	fmt.Println(reply1)
}