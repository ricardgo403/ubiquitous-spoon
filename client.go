package main

import (
	"fmt"
	"net/rpc"
	"strings"

	"./args"
)

// type Args struct {
// 	Name    string
// 	Subject string
// 	Grade   uint
// }

func client() {
	c, err := rpc.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var opc uint64
	for {
		fmt.Println("1.- Capturar calificación de materia por alumno")
		fmt.Println("2.- Mostrar promedio de un alumno")
		fmt.Println("3.- Mostrar promedio general de todos los alumnos")
		fmt.Println("4.- Mostrar promedio de una materia")
		fmt.Println("0.- Exit")
		fmt.Scanln(&opc)

		switch opc {
		case 1:
			var name, lastname, subject string
			fmt.Print("Nombre: ")
			fmt.Scan(&name)
			fmt.Scanln(&lastname)
			name = strings.ToLower(name)
			lastname = strings.ToLower(lastname)

			fmt.Print("Materia: ")
			fmt.Scanln(&subject)
			subject = strings.ToLower(subject)

			var grade uint
			fmt.Print("Calificación: ")
			fmt.Scanln(&grade)

			args := &args.Args{name + " " + lastname, subject, grade}

			var result string
			err = c.Call("Server.AddGrade", args, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.AddGrade =", result)
			}
		case 2:
			var name, lastname string
			fmt.Print("Nombre: ")
			fmt.Scan(&name)
			fmt.Scanln(&lastname)
			name = strings.ToLower(name)
			lastname = strings.ToLower(lastname)

			args := name + " " + lastname

			var result uint
			err = c.Call("Server.StudentAverage", args, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.StudentAverage =", result)
			}
		case 3:
			var result uint
			err = c.Call("Server.GeneralAverage", "", &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.GeneralAverage =", result)
			}
		case 4:
			var subject string
			fmt.Print("Materia: ")
			fmt.Scanln(&subject)
			subject = strings.ToLower(subject)

			var result uint
			err = c.Call("Server.SubjectAverage", subject, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.SubjectAverage =", result)
			}
		case 0:
			return
		}
	}
}

func main() {
	client()
}
