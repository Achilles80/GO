package main 
import(
	"log"
	"net"
	"net/rpc"
	"strings"
	"errors"
	"fmt"
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
type Ums struct{
	Courses map[string]Course
}
func (u *Ums) AddCourse(c* Course,reply *string) error {
	c.Students=[]string{}
	rev:=strings.Replace(c.Code," ","",-1)
	rev=strings.ToUpper(rev)
	c.Code=rev 
	u.Courses[c.Code]=*c 
	*reply="Course Created Successfully"
	return nil
}
func (u *Ums) EnrollStudent(e* EnrollArgs,reply *string) error {
	rev:=strings.Replace(e.Code," ","",-1)
	rev=strings.ToUpper(rev)
	c,ok:=u.Courses[rev]
	if !ok{
		return errors.New("Either Course code is wrong or course Doesnt Exist")
	}
	c.Students=append(c.Students,e.StudentName)
	u.Courses[rev]=c 
	*reply="Student Enrolled Successfully"
	return nil
}
func (u *Ums) GetCourseDetails(c* CourseArgs,reply *string) error{
	rev:=strings.Replace(c.Code," ","",-1)
	rev=strings.ToUpper(rev)
	c1,ok:=u.Courses[rev]
	if !ok{
		return errors.New("Either Course code is wrong or course Doesnt Exist")
	}
	r:=fmt.Sprintf(" Course Code: %s | Course Title: %s | ",c1.Code,c1.Title)
	for i:=0;i<len(c1.Students);i++{
		r+=c1.Students[i]+" | "
	}
	*reply=r 
	return nil
}
func main(){
	ums:=&Ums{Courses:make(map[string]Course)}
	if err:=rpc.Register(ums);err!=nil{
		log.Fatal(err)
	}
	ln,err:=net.Listen("tcp",":1234")
	if err!=nil{
		log.Fatal(err)
	}
	log.Println("UMS RPC Listening on :1234")
	for{
		conn,err:=ln.Accept()
		if err!=nil{
			continue
		}
		go rpc.ServeConn(conn)
	}
}