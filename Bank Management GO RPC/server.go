package main 
import(
	"log"
	"errors"
	"net"
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
type Bank struct{
	Accounts map[int]Account
}
func (b *Bank) Create(a *Account,reply *string) error{
	_, ok := b.Accounts[a.AccountNumber]
    if ok {
        return errors.New("Account already exists! Cannot overwrite.")
    }
	a.Balance=0 
	b.Accounts[a.AccountNumber]=*a
	*reply="Account Created Successfully"
	return nil
}
func (b *Bank) Deposit(a *AmountArgs,reply *string) error{
	a1,ok:=b.Accounts[a.AccountNumber]
	if !ok{
		return errors.New("Account not Found")
	}
	if a.Amount<0{
		return errors.New("Amount cannot be negative")
	}
	a1.Balance+=a.Amount
	b.Accounts[a.AccountNumber]=a1
	*reply="Amount Deposited Successfully"
	return nil
}
func (b *Bank) Withdraw(a *AmountArgs,reply *string) error{
	a1,ok:=b.Accounts[a.AccountNumber]
	if !ok{
		return errors.New("Account not Found")
	}
	if a.Amount<0{
		return errors.New("Amount cannot be negative")
	}
	if a1.Balance<a.Amount{
		return errors.New("Insufficient Balance")
	}
	a1.Balance-=a.Amount
	b.Accounts[a.AccountNumber]=a1
	*reply="Amount Withdrawn Successfully"
	return nil
}
func (b *Bank) CheckBalance(a *AccountArgs,reply *int) error{
	a1,ok:=b.Accounts[a.AccountNumber]
	if !ok{
		return errors.New("Account not Found")
	}
	*reply=a1.Balance 
	return nil
}
func main(){
	bank:=&Bank{Accounts:make(map[int]Account)}
	if err:=rpc.Register(bank);err!=nil{
		log.Fatal(err)
	}
	ln,err:=net.Listen("tcp",":1234")
	if err!=nil{
		log.Fatal(err)
	}
	log.Println("Bank GO RPC listening on :1234")
	for{
		conn,err:=ln.Accept()
		if err!=nil{
			continue
		}
		go rpc.ServeConn(conn)
	}
}