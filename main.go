package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"runtime"

	"github.com/charmbracelet/lipgloss"
)

var (
	corPrimaria = lipgloss.Color("#7D56F4") // Roxo vibrante
	//corSucesso  = lipgloss.Color("#04B575") // Verde

	estiloMenu = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(corPrimaria).
			Padding(1, 3).
			Margin(1)

	estiloTitulo = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(corPrimaria).
			Padding(0, 1).
			Bold(true)

	estiloOpcao = lipgloss.NewStyle().Foreground(lipgloss.Color("#B3B3B3"))
)

func CleanInterface() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

func MenuCalculator() {
	CleanInterface()

	tittle := estiloTitulo.Render(
		"░█▀▀░█░░░▀█▀░░░░░█▀▀░█▀█░░░░░█▀▀░█▀█░█░░░█▀▀░█░█░█░░░█▀█░▀█▀░█▀█░█▀▄\n░█░░░█░░░░█░░▄▄▄░█░█░█░█░▄▄▄░█░░░█▀█░█░░░█░░░█░█░█░░░█▀█░░█░░█░█░█▀▄\n░▀▀▀░▀▀▀░▀▀▀░░░░░▀▀▀░▀▀▀░░░░░▀▀▀░▀░▀░▀▀▀░▀▀▀░▀▀▀░▀▀▀░▀░▀░░▀░░▀▀▀░▀░▀")

	opcoes := lipgloss.JoinVertical(lipgloss.Left,
		estiloOpcao.Render("1. [+] Open Calculator"),
		estiloOpcao.Render("2. [i] INFO"),
		estiloOpcao.Render("3. [+] quit/exit"),
	)

	menuFinal := estiloMenu.Render(
		lipgloss.JoinVertical(lipgloss.Center, tittle, "\n", opcoes),
	)

	fmt.Println(menuFinal)
}

func CalculatorMenu() {
	CleanInterface()

	tittle := estiloTitulo.Render(
		"░█▀▀░█░░░▀█▀░░░░░█▀▀░█▀█░░░░░█▀▀░█▀█░█░░░█▀▀░█░█░█░░░█▀█░▀█▀░█▀█░█▀▄\n░█░░░█░░░░█░░▄▄▄░█░█░█░█░▄▄▄░█░░░█▀█░█░░░█░░░█░█░█░░░█▀█░░█░░█░█░█▀▄\n░▀▀▀░▀▀▀░▀▀▀░░░░░▀▀▀░▀▀▀░░░░░▀▀▀░▀░▀░▀▀▀░▀▀▀░▀▀▀░▀▀▀░▀░▀░░▀░░▀▀▀░▀░▀")

	opcoes := lipgloss.JoinVertical(lipgloss.Left,
		estiloOpcao.Render(">"),
		estiloOpcao.Render("1.SUM[+]"),
		estiloOpcao.Render("2.SUB[-]"),
		estiloOpcao.Render("3.MULTIPLY[x]"),
		estiloOpcao.Render("4.DIVISION[/]"),
		estiloOpcao.Render("q. [!] quit/exit"),
	)

	menuFinal := estiloMenu.Render(
		lipgloss.JoinVertical(lipgloss.Center, tittle, "\n", opcoes),
	)

	fmt.Println(menuFinal)
}

func calculator(value float64) (
	add func(...float64) float64,
	sub func(...float64) float64,
) {
	res := value
	add = func(x ...float64) float64 {
		for _, num := range x {
			res += num
		}
		if math.Mod(res, 1.0) > 0 {
			res = math.Round(res)
		}
		return res

	}

	sub = func(x ...float64) float64 {
		for _, num := range x {
			res -= num
		}
		if math.Mod(res, 1.0) > 0 {
			res = math.Round(res)
		}
		return res
	}

	//-----------------------
	return
}

func main() {

	const limitValuesInput = 10

InfiniteLoopAPP:
	for {
		var mainChoice string

		MenuCalculator()
		fmt.Print(" > Choice your option: ")
		_, _ = fmt.Scanln(&mainChoice)

		switch mainChoice {

		case "1":
		OperationMenu:
			for {
				var (
					valuesOp      []float64 //the values to operate
					opChoice      string    //value to switch the operators
					howManyValues int       //length the user want to operate
				)
				CalculatorMenu()
				//operation choice
				fmt.Println("\nWhat Operation you want to start?" +
					"\n>:")
				_, _ = fmt.Scanln(&opChoice)

				if opChoice == "q" {
					break OperationMenu
				}
				fmt.Printf("How many value you want operate? Limit is %d"+
					"\n>:", limitValuesInput)
				_, _ = fmt.Scanln(&howManyValues)

				//only to verify if the user put a value to this verification
				if howManyValues <= 0 || howManyValues > limitValuesInput {
					fmt.Println("Invalid quantity")
					continue
				}
				for i := 0; i < howManyValues; i++ {
					var temp float64
					fmt.Println("\n>Write the ", i+1, "º value")
					_, _ = fmt.Scan(&temp)
					valuesOp = append(valuesOp, temp) //store the values in array valuesOp
				}

				_, _ = fmt.Scanln() //to clean stdin in keyboard

				firstValue := valuesOp[0]
				otherValues := valuesOp[1:]
				add, sub := calculator(firstValue)

				switch opChoice {
				case "1":
					res := add(otherValues...)
					fmt.Printf("The result of sum is %.2f", res)

				case "2":
					res := sub(otherValues...)
					fmt.Printf("The result of sum is %.2f", res)

				default:
					fmt.Println("Invalid operation")
				}
				fmt.Println("\n\nPress Enter to return...")
				var pause string
				_, _ = fmt.Scanln(&pause)
			} //----------main for loop
		case "2":
			CleanInterface()
			fmt.Println("This app was built by Ricardo Kassoma, an beginner go developer" +
				"\n© 2026 - RICARDO KASSOMA")
			fmt.Println("\n\nPress Enter to return...")
			var pause string
			_, _ = fmt.Scanln(&pause)
		case "q":
			fmt.Println("Exiting... Goodbye!")
			break InfiniteLoopAPP
		}
	}
} //----end of project--------//
