package main 
import(
	"log"
	"fmt"
	"os"
	"bufio"
	"strings"
	"net/rpc"
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
func main(){
	client,err:=rpc.Dial("tcp","localhost:1234")
	if err!=nil{
		log.Fatal(err)
	}
	defer client.Close()
	var reply string 
	var opt int
	reader:=bufio.NewReader(os.Stdin)
		for{
			fmt.Println("=== Galactic Spaceport Control ===")
			fmt.Println("1. Register a ship")
			fmt.Println("2. Add Crew")
			fmt.Println("3. Launch Ship")
			fmt.Println("4. View Ship Status")
			fmt.Println("5. Exit")
			fmt.Println("Enter the Option: ")
			fmt.Scanln(&opt)

		if opt==1{
			fmt.Println("Enter Ship ID: ")
			id,_:=reader.ReadString('\n')
			id=strings.TrimSpace(id)

			fmt.Println("Enter Captain Name: ")
			name,_:=reader.ReadString('\n')
			name=strings.TrimSpace(name)
			args:=&Ship{ShipID:id,Captain:name}
			if err:=client.Call("System.RegisterShip",args,&reply);err!=nil{
				fmt.Println(err)
				continue
			}
			fmt.Println(reply)
		}else if opt==2{
			fmt.Println("Enter Ship ID: ")
			id,_:=reader.ReadString('\n')
			id=strings.TrimSpace(id)

			fmt.Println("Enter Crew Member Name: ")
			name,_:=reader.ReadString('\n')
			name=strings.TrimSpace(name)

			args1:=&CrewArgs{ShipID:id,CrewMember:name}
			if err:=client.Call("System.AddCrew",args1,&reply);err!=nil{
				fmt.Println(err)
				continue
			}
			fmt.Println(reply)
		}else if opt==3{
			fmt.Println("Enter Ship ID: ")
			id,_:=reader.ReadString('\n')
			id=strings.TrimSpace(id)

			args2:=&ShipArgs{ShipID:id}
			if err:=client.Call("System.LaunchShip",args2,&reply);err!=nil{
				fmt.Println(err)
				continue
			}
			fmt.Println(reply)
		}else if opt==4{
			fmt.Println("Enter Ship ID: ")
			id,_:=reader.ReadString('\n')
			id=strings.TrimSpace(id)

			args2:=&ShipArgs{ShipID:id}
			if err:=client.Call("System.GetStatus",args2,&reply);err!=nil{
				fmt.Println(err)
				continue
			}
			fmt.Println(reply)
		}else{
			fmt.Println("exiting....")
			break
		}
	}
}