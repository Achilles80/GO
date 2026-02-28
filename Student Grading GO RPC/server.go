package main 
import(
	"errors"
	"net"
	"log"
	"net/rpc"
)
type Record struct{
	StudentID int 
	Name string 
	Grades []float64
}
type StudArgs struct{
	StudentID int 
}
type GradeArgs struct{
	StudentID int 
	Grade float64 
}
type System struct{
	Class map[int]Record 
}
func (s *System) Add(r *Record,reply *string)error{
	r.Grades=[]float64{}
	s.Class[r.StudentID]=*r
	*reply="Student Added Successfully"
	return nil
}
func (s *System) AddGrade(r *GradeArgs,reply *string)error{
	r1,ok:=s.Class[r.StudentID]
	if !ok{
		return errors.New("Student not found")
	}
	r1.Grades=append(r1.Grades,r.Grade)
	s.Class[r.StudentID]=r1
	*reply="Grade Added Successfully"
	return nil
}
func (s *System) Calc(r *StudArgs,reply *float64)error{
	r1,ok:=s.Class[r.StudentID]
	if !ok{
		return errors.New("Student not found")
	}
	cnt:=0.0
	for i:=0;i<len(r1.Grades);i++{
		cnt+=r1.Grades[i]
	}
	cnt=cnt/float64(len(r1.Grades))
	*reply=cnt 
	return nil
}
func (s *System) Get(r *StudArgs,reply *Record)error{
	r1,ok:=s.Class[r.StudentID]
	if !ok{
		return errors.New("Student not found")
	}
	*reply=r1
	return nil
}
func main(){
	system:=&System{Class:make(map[int]Record)}
	if err:=rpc.Register(system);err!=nil{
		log.Fatal(err)
	}
	ln,err:=net.Listen("tcp",":1234")
	if err!=nil{
		log.Fatal(err)
	}
	log.Println("System listening at port :1234")
	for{
		conn,err:=ln.Accept()
		if err!=nil{
			continue
		}
		go rpc.ServeConn(conn)
	}
}