package main 
import(
	"log"
	"errors"
	"net"
	"net/rpc"
	"strings"
	"fmt"
)
type Ship struct{
	ShipID string
	Captain string 
	Crew []string 
	IsDocked bool
}
type ShipArgs struct{
	ShipID string
}
type CrewArgs struct{
	ShipID string 
	CrewMember string
}
type System struct{
	Ships map[string]Ship
}
func (l *System) RegisterShip(s *Ship,reply *string) error{
	rev:=s.ShipID 
	rev=strings.Replace(rev," ","",-1)
	rev=strings.ToUpper(rev)
	s.ShipID=rev
	s.Crew=[]string{}
	s.IsDocked=true
	l.Ships[rev]=*s 
	*reply="Ship Creation SUccessfull"
	return nil
}
func (l *System) AddCrew(c *CrewArgs,reply *string) error{
	rev:=c.ShipID 
	rev=strings.Replace(rev," ","",-1)
	rev=strings.ToUpper(rev)
	s,ok:=l.Ships[rev]
	if !ok{
		return errors.New("Ship Name is Wrong or Ship Doesnt Exist")
	}	
	s.Crew=append(s.Crew,c.CrewMember)
	l.Ships[rev]=s
	*reply="Crew Member Added Successfully"
	return nil
}
func (l *System) LaunchShip(s *ShipArgs,reply *string) error{
	rev:=s.ShipID 
	rev=strings.Replace(rev," ","",-1)
	rev=strings.ToUpper(rev)
	s1,ok:=l.Ships[rev]
	if !ok{
		return errors.New("Ship Name is Wrong or Ship Doesnt Exist")
	}
	if len(s1.Crew)==0{
		return errors.New("Cannot Launch Empty ship")
	}
	s1.IsDocked=false
	l.Ships[rev]=s1
	*reply="Ship Undocked Successfull"
	return nil
}
func (l *System) GetStatus(s *ShipArgs,reply *string) error{
	rev:=s.ShipID 
	rev=strings.Replace(rev," ","",-1)
	rev=strings.ToUpper(rev)
	s1,ok:=l.Ships[rev]
	if !ok{
		return errors.New("Ship Name is Wrong or Ship Doesnt Exist")
	}
	rev=fmt.Sprintf(" ShipID: %s | Captain: %s | Docked: %t | Total: %d ",s1.ShipID,s1.Captain,s1.IsDocked,len(s1.Crew))
	*reply=rev
	return nil
}
func main(){
	system:=&System{Ships:make(map[string]Ship)}
	if err:=rpc.Register(system);err!=nil{
		log.Fatal(err)
	}
	ln,err:=net.Listen("tcp",":1234")
	if err!=nil{
		log.Fatal(err)
	}
	log.Println("SCS GO RPC Listening on :1234")
	for{
		conn,err:=ln.Accept()
		if err!=nil{
			continue
		}
		go rpc.ServeConn(conn)
	}
}