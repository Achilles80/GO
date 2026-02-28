package main 
import(
	"fmt"
	"log"
	"net/rpc"
)
type Account struct{
	AccountID int 
	AccountHolder string 
	Balance float64
}
type AccountArgs struct{
	AccountID int
}
type DepositArgs struct{
	AccountID int 
	Amt float64
}
func main(){
	client,err:=rpc.Dial("tcp","localhost:1234")
	if err!=nil{
		log.Fatal(err)
	}
	defer client.Close()
	acc:=&Account{AccountID:101,AccountHolder:"jon"}
	var msgReply string 
	var flReply float64
	if err:=client.Call("Bank.Create",acc,&msgReply);err!=nil{
		log.Fatal(err)
	}
	fmt.Println(msgReply)
	args:=&AccountArgs{AccountID:101}
	if err:=client.Call("Bank.CheckBalance",args,&flReply);err!=nil{
		log.Fatal(err)
	}
	fmt.Println(flReply)
	args1:=&DepositArgs{AccountID:101,Amt:100.0}
	if err:=client.Call("Bank.Deposit",args1,&msgReply);err!=nil{
		log.Fatal(err)
	}
	fmt.Println(msgReply)
	if err:=client.Call("Bank.CheckBalance",args,&flReply);err!=nil{
		log.Fatal(err)
	}
	fmt.Println(flReply)
	args2:=&DepositArgs{AccountID:101,Amt:50.0}
	if err:=client.Call("Bank.Withdraw",args2,&msgReply);err!=nil{
		log.Fatal(err)
	}
	fmt.Println(msgReply)
	if err:=client.Call("Bank.CheckBalance",args,&flReply);err!=nil{
		log.Fatal(err)
	}
	fmt.Println(flReply)
}