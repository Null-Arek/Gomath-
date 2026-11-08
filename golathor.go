package main

import (
        "fmt"
        "math"
)

func aritme() {
        var num1, num2 float64
        var option string

        reset := "\033[0m"
        red := "\033[31m"
        green := "\033[32m"
        yellow := "\033[33m"
        blue := "\033[34m"
        purple := "\033[35m"
        cyan := "\033[36m"
        for {
                fmt.Println(red+"Ingrese el nombre de su operacion deseada:", reset)
                fmt.Println(green+"Sumar", reset)
                fmt.Println(yellow+"Restar", reset)
                fmt.Println(blue+"Multiplicar", reset)
                fmt.Println(purple+"Division", reset)
                fmt.Println(cyan+"Potencia", reset)
                fmt.Println("Raiz")
                fmt.Scanln(&option)

                fmt.Println("Ingresa el primer valor:")
                fmt.Scanln(&num1)
                fmt.Println("Ingresa el segundo valor:")
                fmt.Scanln(&num2)

                switch option {

                case "Sumar":
                        suma := num1 + num2
                        fmt.Println(green+"El resultado es:", suma, reset)
                case "Restar":
                        resta := num1 - num2
                        fmt.Println(yellow+"El resultado es:", resta, reset)
                case "Multiplicar":
                        multi := num1 * num2
                        fmt.Println(blue+"El resultado es:", multi, reset)
                case "Division":
                        divi := num1 / num2
                        fmt.Println(purple+"El resultado es:", divi, reset)
                case "Potencia":
                        poten := math.Pow(num1, num2)
                        fmt.Println(cyan+"El resultado es:", poten, reset)
                case "Raiz":
                        raiz := math.Sqrt(num1)
                        fmt.Printf("El resultado es: %.2f\n", raiz)

                }
        }
}
func main() {
        aritme()
}












