package main
import(
	"log"
	"errors"
	"net"
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
type Bank struct{
	Accounts map[int]Account
}
func (b *Bank) Create(a *Account,reply *string) error {
	a.Balance=0.0 
	b.Accounts[a.AccountID]=*a 
	*reply="Account created successfully"
	return nil
}
func (b *Bank) Deposit(args *DepositArgs,reply *string) error {
	a,ok:=b.Accounts[args.AccountID]
	if !ok{
		return errors.New("Account doesnt exist")
	}
	a.Balance=a.Balance+args.Amt 
	b.Accounts[args.AccountID]=a 
	*reply="Amount deposited successfully"
	return nil
}
func (b *Bank) Withdraw(args *DepositArgs,reply *string) error{
	a,ok:=b.Accounts[args.AccountID]
	if !ok{
		return errors.New("Account doesnt exist")
	}
	if a.Balance<args.Amt{
		return errors.New("Insufficient funds")
	}
	a.Balance=a.Balance-args.Amt 
	b.Accounts[args.AccountID]=a
	*reply="Amount withdrawn successfully"
	return nil
}
func (b *Bank) CheckBalance(args *AccountArgs,reply *float64) error{
	a,ok:=b.Accounts[args.AccountID]
	if !ok{
		return errors.New("Account doesnt exist")
	}
	*reply=a.Balance
	return nil
}
func main(){
	bank:=&Bank{Accounts:make(map[int]Account)}
	if err:=rpc.Register(bank); err!=nil{
		log.Fatal(err)
	}
	ln,err:=net.Listen("tcp",":1234")
	if err!=nil{
		log.Fatal(err)
	}
	log.Println("Bank RPC Listening on :1234")
	for{
		conn,err:=ln.Accept()
		if err!=nil{
			continue 
		}
		go rpc.ServeConn(conn)
	}
}