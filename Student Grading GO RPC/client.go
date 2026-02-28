package main
import(
	"fmt"
	"log"
	"net/rpc"
)
type Record struct{
	StudentID int 
	Name string 
	Grades []float64
}
type StudArgs struct{
	StudentID int 
}
type GradeArgs struct{
	StudentID int 
	Grade float64 
}
func main(){
	client,err:=rpc.Dial("tcp","localhost:1234")
	if err!=nil{
		log.Fatal(err)
	}
	defer client.Close()
	var reply string 
	var reply1 float64
	var reply2 Record
	args:=&Record{StudentID:101,Name:"jon snow"}
	if err:=client.Call("System.Add",args,&reply);err!=nil{
		log.Fatal(err)
	}
	fmt.Println(reply)
	args1:=&GradeArgs{StudentID:101,Grade:100.0}
	if err:=client.Call("System.AddGrade",args1,&reply);err!=nil{
		log.Fatal(err)
	}
	fmt.Println(reply)
	args2:=&GradeArgs{StudentID:101,Grade:80.0}
	if err:=client.Call("System.AddGrade",args2,&reply);err!=nil{
		log.Fatal(err)
	}
	fmt.Println(reply)
	args3:=&StudArgs{StudentID:101}
	if err:=client.Call("System.Calc",args3,&reply1);err!=nil{
		log.Fatal(err)
	}
	fmt.Println(reply1)
	if err:=client.Call("System.Get",args3,&reply2);err!=nil{
		log.Fatal(err)
	}
	fmt.Println(reply2)
}