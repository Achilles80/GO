package main
import(
	"fmt"
	"log"
	"net/rpc"
)
type Item struct{
	Name string 
	Price float64
}
type Cart struct{
	UserID int
	Items []Item 
	DiscountApplied bool	
}
type UserArgs struct{
	UserID int
}
type AddArgs struct{
	UserID int
	Item Item 
}
type DiscArgs struct{
	UserID int 
	Code string 
}
func main(){
	client,err:=rpc.Dial("tcp","localhost:1234")
	if err!=nil{
		log.Fatal(err)
	}
	defer client.Close()
	var reply string 
	var reply1 Cart
    args:=&UserArgs{UserID:101}
	if err:=client.Call("System.CreateCart",args,&reply);err!=nil{
		log.Fatal(err)
	}
	fmt.Println(reply)
	args1:=&AddArgs{UserID:101,Item:Item{Name:"Phone",Price:120.00}}
	if err:=client.Call("System.AddItem",args1,&reply);err!=nil{
		log.Fatal(err)
	}
	fmt.Println(reply)
	args2:=&DiscArgs{UserID:101,Code:"sa Ve 10"}
	if err:=client.Call("System.ApplyPromoCode",args2,&reply);err!=nil{
		log.Fatal(err)
	}
	fmt.Println(reply)
	args3:=&UserArgs{UserID:101}
	if err:=client.Call("System.GetCart",args3,&reply1);err!=nil{
		log.Fatal(err)
	}
	fmt.Println(reply1)
	if err:=client.Call("System.GenerateReceipt",args3,&reply);err!=nil{
		log.Fatal(err)
	}
	fmt.Println(reply)
}

