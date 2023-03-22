package main

import (
    "fmt"
    "strconv"
)

func main() {
    var num1, num2 float64
    var operator string

    fmt.Println("Calculadora básica em Go")

    fmt.Print("Digite o primeiro número: ")
    fmt.Scanln(&num1)

    fmt.Print("Digite o segundo número: ")
    fmt.Scanln(&num2)

    fmt.Print("Digite a operação (+, -, *, /): ")
    fmt.Scanln(&operator)

    result := 0.0
    switch operator {
    case "+":
        result = num1 + num2
    case "-":
        result = num1 - num2
    case "*":
        result = num1 * num2
    case "/":
        result = num1 / num2
    default:
        fmt.Println("Operador inválido!")
        return
    }

    fmt.Println("Resultado: " + strconv.FormatFloat(result, 'f', 2, 64))
}
