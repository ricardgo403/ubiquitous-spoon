package main

import (
	"container/list"
	"fmt"
	"net"
	"net/rpc"

	"./args"
)

type Server struct{}

type Student struct {
	Name    string
	Subject map[string]uint
}

var l = list.New()

func validate(args *args.Args) bool {
	for e := l.Front(); e != nil; e = e.Next() {
		myElement := e.Value.(Student)
		if myElement.Name == args.Name {
			for key := range myElement.Subject {
				if key == args.Subject {
					return true
				}
			}
		}
	}
	return false
}

func (this *Server) AddGrade(args *args.Args, reply *string) error {
	var result string
	if validate(args) {
		result = "La calificación ya estaba asignada y no ha sido modificada..."
	} else {
		myStudent := Student{Name: args.Name, Subject: map[string]uint{}}
		myStudent.Subject[args.Subject] = args.Grade

		l.PushBack(myStudent)
		result = "La calificación ha sido asignada!"
	}
	*reply = result
	return nil
}

func (this *Server) StudentAverage(args string, reply *uint) error {
	total := uint(0)
	cont := uint(0)
	for e := l.Front(); e != nil; e = e.Next() {
		myElement := e.Value.(Student)
		if myElement.Name == args {
			for _, value := range myElement.Subject {
				total += value
				cont++
			}
		}
	}
	*reply = total / cont
	return nil
}

func (this *Server) GeneralAverage(args string, reply *uint) error {
	total := uint(0)
	cont := uint(0)
	for e := l.Front(); e != nil; e = e.Next() {
		myElement := e.Value.(Student)
		for _, value := range myElement.Subject {
			total += value
			cont++
		}
	}
	*reply = total / cont
	return nil
}

func (this *Server) SubjectAverage(args string, reply *uint) error {
	total := uint(0)
	cont := uint(0)
	for e := l.Front(); e != nil; e = e.Next() {
		myElement := e.Value.(Student)
		for key, value := range myElement.Subject {
			if key == args {
				total += value
				cont++
			}
		}
	}
	*reply = total / cont
	return nil
}

func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main() {
	go server()

	var input string
	fmt.Scanln(&input)
}
