package main

import (
	"fmt"
	"time"
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
func ProcessResults(student Student,resultChan chan string){
	fmt.Println("Checking result for :" ,student.name)
	time.Sleep(2 *time.Second)


	resultChan <- student.name
}
func main(){
	students :=[]Student{}
	nextId:=IdGenerator()
	students = append(students, 
	Student{nextId(),"Atharva",85,"passed"},
	Student{nextId(),"demo",70,"passed"},
	Student{nextId(),"check",50,"failed"},
)
for _,student := range students{
	student.Display()
}
var p Person

for _,person :=range students{
	p=person
	p.Show()
}
PrintAll(students)
resultChan := make(chan string)
for _,student :=range students{
	go ProcessResults(student,resultChan)
}
for i := 0; i < len(students); i++ {
	name:= <- resultChan
	fmt.Println("Result Processed for :" , name)
}

}