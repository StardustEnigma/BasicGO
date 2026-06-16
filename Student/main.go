package main

import (
	"fmt"
	"time"
	"sync"
)

type Student struct {
	studentId int
	name string
	marks int
	status string
}

func (s Student) Display(){
	fmt.Println("ID : ",s.studentId,"Name :",s.name,"Marks :",s.marks)
}

type Person interface{
	Show()
}
func(s Student) Show(){
	fmt.Println("Student :",s.name)
}
func PrintAll[T any](items []T){
	for _,item :=range items{
		fmt.Println(item)
	}
}
func IdGenerator() func() int{
	id:=0
	return func() int {
		id++
		return id
	}
}
func ProcessResults(student Student,passChan chan string,failChan chan string,mu *sync.Mutex,stats map[string]int){
	fmt.Println("Checking result for :" ,student.name)
	time.Sleep(2 *time.Second)	
	if student.marks <= 40 {
		failChan <- student.name
		mu.Lock()
		stats["fail"]++
		mu.Unlock()
		return
	}
	passChan <- student.name
	mu.Lock()
	stats["pass"]++
	mu.Unlock()
}
func main(){
	students :=[]Student{}
	nextId:=IdGenerator()
	students = append(students, 
	Student{nextId(),"Atharva",85,"passed"},
	Student{nextId(),"demo",70,"passed"},
	Student{nextId(),"check",40,"failed"},
)
for _,student := range students{
	student.Display()
}
var p Person

for _,person :=range students{
	p=person
	p.Show()
}
var mu sync.Mutex
stats := map[string]int{
	"pass":0,
	"fail":0,
}
PrintAll(students)
passChan := make(chan string)
failChan:= make(chan string)
for _,student :=range students{
	go ProcessResults(student,passChan,failChan,&mu,stats)
}
for i := 0; i < len(students); i++ {
	select{
	case name := <-passChan :
		fmt.Println("Pass :",name)

	case name := <-failChan:
		fmt.Println("fail :",name)
}
}
fmt.Println(stats)

}