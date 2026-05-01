# 🧮 calculator-cli

A terminal-based calculator built with **Go**, featuring a stylish TUI (Terminal User Interface) powered by [Charm's lipgloss](https://github.com/charmbracelet/lipgloss).

> Built by **Ricardo Kassoma** — a beginner Go developer.
> © 2026 - RICARDO KASSOMA

---

## ✨ Features

- Interactive terminal menus with styled UI (rounded borders, colors)
- Support for up to **10 values** per operation
- Available operations:
    - `[1]` **Sum** — Add multiple values
    - `[2]` **Subtraction** — Subtract multiple values
    - `[3]` **Multiplication** — Multiply multiple values
    - `[4]` **Division** — Divide multiple values
    - `[5]` **Power** — Raise a value to a given exponent
    - `[6]` **Modulo (Rest)** — Get the remainder of a division
- Cross-platform screen clearing (Windows & Unix/Linux/macOS)
- Automatic rounding for floating-point results

---

## 📋 Requirements

- [Go](https://golang.org/dl/) `1.18+`

---

## 🚀 Installation

```bash
# Clone the repository
git clone https://github.com/ricarditti5/calculator-cli.git

# Navigate into the project directory
cd calculator-cli

# Install dependencies
go mod tidy

# Run the app
go run main.go
```

Or build a binary:

```bash
go build -o calculator-cli
./calculator-cli
```

---

## 🖥️ Usage

When you launch the app, you'll be presented with the **main menu**:

```
1. [+] Open Calculator
2. [i] INFO
q.     quit/exit
```

After selecting `1`, the **calculator menu** appears:

```
1. SUM       [+]
2. SUB       [-]
3. MULTIPLY  [x]
4. DIVISION  [/]
5. POWER     [^]
6. MODULE    [%]
q. quit/exit
```

**Example flow — summing 3 numbers:**

```
What Operation you want to start?
>: 1

How many values you want to operate? Limit is 10
>: 3

> Write the 1º value: 10
> Write the 2º value: 20
> Write the 3º value: 5

The result of sum is 35.00
```

> **Note:** For `Power` and `Modulo`, only the **first two values** entered are used.

---

## 📁 Project Structure

```
calculator-cli/
└── main.go       # Entry point — menus, UI styles, calculator logic
```

---

## 📦 Dependencies

| Package | Description |
|---|---|
| [`github.com/charmbracelet/lipgloss`](https://github.com/charmbracelet/lipgloss) | Terminal styling (borders, colors, layout) |

---

## 🤝 Contributing

This is a learning project, but contributions and suggestions are welcome!

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/my-feature`)
3. Commit your changes (`git commit -m 'Add my feature'`)
4. Push to the branch (`git push origin feature/my-feature`)
5. Open a Pull Request

---

## 📄 License

This project is open source and available under the [MIT License](LICENSE).