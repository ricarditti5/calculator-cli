package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"os/exec"
	"runtime"

	"github.com/charmbracelet/lipgloss"
)

var (
	corPrimaria = lipgloss.Color("#7D56F4") // Roxo vibrante
	corSucesso  = lipgloss.Color("#04B575") // Verde

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
		"░█▀▀░█░░░▀█▀░░░░░█▀▀░█▀█░░░░░█▀▀░█▀█░█░░░█▀▀░█░█░█░░░█▀█░▀█▀░█▀█░█▀▄\n░█░░░█░░░░█░░▄▄▄░█░█░█░█░▄▄▄░█░░░█▀█░█░░░█░░░█░█░█░░░█▀█░░█░░█░█░█▀▄\n░▀▀▀░▀▀▀░▀▀▀░░░░░▀▀▀░▀▀▀░░░░░▀▀▀░▀░▀░▀▀▀░▀▀▀░▀▀▀░▀▀▀░▀░▀░░▀░░▀▀▀░▀░▀" +
			"\nWhat Operation you want to start?")

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

func calculator(value float64) (
	add func(...float64) float64,
	sub func(...float64) float64,
) {
	res := value
	add = func(x ...float64) float64 {
		if math.Mod(res, 1.0) == 0 {
			for _, num := range x {
				n := num
				res += n
				math.Round(res)
			}
			return res
		}
		for _, num := range x {
			n := num
			res += n
		}
		return res

	}

	sub = func(x ...float64) float64 {
		if math.Mod(res, 1.0) == 0 {
			for _, num := range x {
				n := num
				res -= n
				math.Round(res)
			}
			return res
		}
		for _, num := range x {
			n := num
			res -= n
		}
		return res
	}

	//-----------------------
	return
}

func main() {
InfiniteLoopAPP:
	for {
		var (
			mainChoice int
		)

		MenuCalculator()
		fmt.Print(" > Choice your option: ")
		_, err := fmt.Scan(&mainChoice)
		if err != nil {
			fmt.Println(errors.New("something got wrong"), err)
		}

		switch mainChoice {
		case 1:
			CalculatorMenu()
		case 2:
			CleanInterface()
			fmt.Println("This app was built by Ricardo Kassoma, an beginner go developer" +
				"\n© 2026 - RICARDO KASSOMA" +
				"\nClick enter to return...")
			_, _ = fmt.Scanln()
			_, _ = fmt.Scanln()
		case 3:
			fmt.Println("Exiting... Goodbye!")
			break InfiniteLoopAPP
		}
	}
}

/*
- calcMenu(escolher qual operacao)
- print para saber quanto numeros quer calcular
- SWITCH para chamar as func add e sub
- print para mostrar calculo
*/
