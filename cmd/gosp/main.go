package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/0xErwin1/gosp/internal/lexer"
	"github.com/chzyer/readline"
)

func run(source *bytes.Buffer) {
	reader := strings.NewReader(source.String())
	lexer := lexer.NewLexer(reader, source)
	tokens, errs := lexer.Lex()

	for _, err := range errs {
		fmt.Println(err)
	}

	for _, token := range tokens {
		fmt.Println(token)
	}
}

func relp() {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:          " |> ",
		HistoryFile:     "/tmp/gosp_history",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	defer rl.Close()

	var cmds []string
	var parenCount int
	var inString bool

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		cmds = append(cmds, line)

		for _, c := range line {
			switch c {
			case '(':
				parenCount++
			case ')':
				parenCount--
			case '"':
				inString = !inString
			}
		}

		if parenCount > 0 || inString {
			rl.SetPrompt(" |>> ")
			continue
		}

		cmd := strings.Join(cmds, " ")
		cmds = cmds[:0]
		rl.SetPrompt(" |> ")

		err = rl.SaveHistory(cmd)
		if err != nil {
			log.Fatal(err)
		}

		if cmd == "exit" {
			fmt.Println("Bye!")
			os.Exit(0)
			return
		}

		run(bytes.NewBufferString(cmd))
	}
}

func main() {
	relp()
}
