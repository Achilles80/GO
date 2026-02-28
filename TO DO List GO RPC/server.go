package main 
import(
	"log"
	"errors"
	"net"
	"net/rpc"
)
type Task struct{
	TaskID int
	Description string
	Completed bool
}
type TaskArgs struct{
	TaskID int 
}
type TODO struct{
	List map[int]Task 
}
func (l *TODO) Add(t *Task,reply *string) error{
	t.Completed = false
	l.List[t.TaskID]=*t
	*reply="Task Added Successfully"
	return nil
}
func (l *TODO) Complete(t *TaskArgs,reply *string) error{
	t1,ok:=l.List[t.TaskID]
	if !ok{
		return errors.New("Task Doesnt Exist")
	}
	t1.Completed=true
	l.List[t.TaskID]=t1 
	*reply="Task Completed Successfully"
	return nil
}
func (l *TODO) Delete(t *TaskArgs,reply *string) error{
	_,ok:=l.List[t.TaskID]
	if !ok{
		return errors.New("Task Doesnt Exist")
	}

	delete(l.List,t.TaskID)
	*reply="Task Deleted Successfully"
	return nil
}
func (l *TODO) ListTasks(t *TaskArgs,reply *[]Task) error{
	var tasks []Task
	for _,v := range l.List{
		tasks=append(tasks,v)
	}
	*reply=tasks 
	return nil
}
func main(){
	todo:=&TODO{List:make(map[int]Task)}
	if err:=rpc.Register(todo); err!=nil{
		log.Fatal(err)
	}
	ln,err:=net.Listen("tcp",":1234")
	if err!=nil{
		log.Fatal(err)
	}
	log.Println("TODO List RPC Listening on :1234")
	for{
		conn,err:=ln.Accept()
		if err!=nil{
			continue 
		}
		go rpc.ServeConn(conn)
	}
}