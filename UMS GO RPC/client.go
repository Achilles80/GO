package main 
import(
	"fmt"
	"log"
	"net/rpc"
	"os"
	"bufio"
	"strings"
)
type Course struct{
	Code string 
	Title string 
	Students []string
}
type CourseArgs struct{
	Code string
}
type EnrollArgs struct{
	Code string 
	StudentName string
}
func main(){
	client,err:=rpc.Dial("tcp","localhost:1234")
	if err!=nil{
		log.Fatal(err)
	}
	defer client.Close()
	var reply string 
	var code string 
	var name string 
	var title string 
	var opt int 
	reader:=bufio.NewReader(os.Stdin)
	for{
		fmt.Println("-- University System --")
		fmt.Println("1. Add a New Course")
		fmt.Println("2. Enroll a Student")
		fmt.Println("3. View Course Details")
		fmt.Println("4. Exit")
		fmt.Scanln(&opt)
		if opt==1{
			fmt.Println("Enter Course Code: ")
			code,_=reader.ReadString('\n')
			code=strings.TrimSpace(code)

			fmt.Println("Enter Course Title: ")
			title,_=reader.ReadString('\n')
			title=strings.TrimSpace(title)

			args:=&Course{Code:code,Title:title}
			if err:=client.Call("Ums.AddCourse",args,&reply);err!=nil{
				log.Fatal(err)
			}
			fmt.Println(reply)
		}else if opt==2{
			fmt.Println("Enter Course Code: ")
			code,_=reader.ReadString('\n')
			code=strings.TrimSpace(code)

			fmt.Println("Enter Student Name: ")
			name,_=reader.ReadString('\n')
			name=strings.TrimSpace(name)

			args1:=&EnrollArgs{Code:code,StudentName:name}
			if err:=client.Call("Ums.EnrollStudent",args1,&reply);err!=nil{
				log.Fatal(err)
			}
			fmt.Println(reply)
		}else if opt==3{
			fmt.Println("Enter Course Code: ")
			code,_=reader.ReadString('\n')
			code=strings.TrimSpace(code)

			args2:=&CourseArgs{Code:code}
			if err:=client.Call("Ums.GetCourseDetails",args2,&reply);err!=nil{
				log.Fatal(err)
			}
			fmt.Println(reply)
		}else if opt==4{
			fmt.Println("exiting....")
			break 
		}
	}
}