package main 
import(
	"log"
	"errors"
	"net"
	"net/rpc"
	"strings"
	"fmt"
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
type System struct{
	Carts map[int]Cart
}
func (s *System) CreateCart(u *UserArgs,reply *string) error{
	c:=Cart{UserID:u.UserID,Items:[]Item{},DiscountApplied:false}
	s.Carts[u.UserID]=c 
	*reply="Cart created successfully"
	return nil
}
func (s *System) AddItem(args *AddArgs,reply *string) error{
	c,ok:=s.Carts[args.UserID]
	if !ok{
		return errors.New("User Doesnt Exist")
	}
	c.Items=append(c.Items,args.Item)
	s.Carts[args.UserID]=c
	*reply="item successfully added to cart"
	return nil
}
func (s *System) ApplyPromoCode(args *DiscArgs,reply *string)error{
	c,ok:=s.Carts[args.UserID]
	if !ok{
		return errors.New("User Doesnt Exist")
	}
	parts:=strings.Split(args.Code," ")
	rev:=""
	for i:=0;i<len(parts);i++{
		rev+=parts[i]
	}
	rev=strings.ToUpper(rev)
	if rev!="SAVE10"{
		return errors.New("Invalid promo code")
	}
	c.DiscountApplied=true
	s.Carts[args.UserID]=c
	*reply="DIscount Applied Successfully"
	return nil
}
func (s *System) GetCart(args *UserArgs,reply *Cart)error{
	c,ok:=s.Carts[args.UserID]
	if !ok{
		return errors.New("User Doesnt Exist")
	}
	*reply=c
	return nil
}
func (s *System) GenerateReceipt(args *UserArgs,reply *string)error{
	c,ok:=s.Carts[args.UserID]
	if !ok{
		return errors.New("User Doesnt Exist")
	}
	cnt:=0.0
	for i:=0;i<len(c.Items);i++{
		cnt+=c.Items[i].Price
	}
	if c.DiscountApplied{
		cnt=cnt*0.9
	}
	rev := fmt.Sprintf(
    "User: %d | Total Items: %d | Final Price: %.2f",
    c.UserID,
    len(c.Items),
    cnt,
)
	*reply=rev
	return nil
}
func main(){
	system:=&System{Carts:make(map [int]Cart)}
	if err:=rpc.Register(system);err!=nil{
		log.Fatal(err)
	}
	ln,err:=net.Listen("tcp",":1234")
	if err!=nil{
		log.Fatal(err)
	}
	log.Println("System litening on :1234")
	for {
		conn,err:=ln.Accept()
		if err!=nil{
			continue
		}
		go rpc.ServeConn(conn)
	}
}
