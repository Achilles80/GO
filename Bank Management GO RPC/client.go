package main 
import(
	"log"
	"fmt"
	"net/rpc"
)
type Account struct{
	AccountNumber int 
	AccountHolderName string 
	Balance int
}
type AccountArgs struct{
	AccountNumber int 
}
type AmountArgs struct{
	AccountNumber int 
	Amount int 
}
func main(){
	client,err:=rpc.Dial("tcp","localhost:1234")
	if err!=nil{
		log.Fatal(err)
	}
	defer client.Close()
	var reply string 
	var reply1 int 
	args:=&Account{AccountNumber:101,AccountHolderName:"Adithya"}
	if err:=client.Call("Bank.Create",args,&reply);err!=nil{
		fmt.Println(err)
	}
	fmt.Println(reply)
	args1:=&AmountArgs{AccountNumber:101,Amount:100}
	if err:=client.Call("Bank.Deposit",args1,&reply);err!=nil{
		fmt.Println(err)
	}
	fmt.Println(reply)
	args2:=&AmountArgs{AccountNumber:101,Amount:50}
	if err:=client.Call("Bank.Withdraw",args2,&reply);err!=nil{
		fmt.Println(err)
	}
	fmt.Println(reply)
	args3:=&AmountArgs{AccountNumber:101,Amount:100}
	if err:=client.Call("Bank.Withdraw",args3,&reply);err!=nil{
		fmt.Println(err)
	}else{
	fmt.Println(reply)
	}
	args4:=&AccountArgs{AccountNumber:101}
	if err:=client.Call("Bank.CheckBalance",args4,&reply1);err!=nil{
		fmt.Println(err)
	}
	fmt.Println(reply1)
}