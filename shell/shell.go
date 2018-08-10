package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"goprojects/shell/jobs"
	"strconv"
)

type command struct {	
	bin string
	args []string
	command string
	empty bool
	background bool
	exists bool
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var l = jobs.NewList()
	for {
		fmt.Print("go-shell: ")
		in, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("go-shell: Unable to process your input")
			continue
		}
		parse := parseCMD(in)
		if parse.empty {
			continue
		} else if(isBuiltIn(parse.bin)) {
			execBuiltIn(parse, l)
		} else if(!parse.exists) {
			fmt.Printf("go-shell: binary not found: %s\n", parse.bin)
			continue
		} else {
			execBin(parse, l)
		}
	}
}

func parseCMD(str string) (*command) {
	str = str[:len(str)-1]
	var tokens = strings.Split(str, " ")
	/*remove end of line*/
	var c = command{}
	c.command = str
	c.empty = len(str) < 1
	/*verify binary*/
	var bin = tokens[0]
	c.bin = bin
	_, err := exec.LookPath(bin)
	if err == nil {
		c.exists = true
	}
	if !c.empty && tokens[len(tokens) - 1] == "&" {
		c.background = true
	}
	if !c.empty {
		c.args = tokens[1:]
	}
	return &c
}

func isBuiltIn(in string) bool {
	var m = map[string]bool{"help": true,
		"cd":   true,
		"jobs": true,
		"fg":   true,
		"bg":   true,
		"kill": true,
	}
	_, ok := m[in]
	return ok
}

func execBuiltIn(parseCMD *command, l *jobs.List) {
	switch(parseCMD.bin) {
	case "jobs":
		fmt.Print(l.Print())
		break;
	case "fg":
		handleFG(l)
	default:
		fmt.Println("unsupported builtin")
		break;
	}
}

func handleFG(l *jobs.List) {
	fmt.Printf("go-shell: please provide fg pid ")
	var reader = bufio.NewReader(os.Stdin)
	in, _ := reader.ReadString('\n')
	pid, _ := strconv.Atoi(in[:len(in) - 1])
	l.Get(pid).CmdObj.Wait()
}

func execBin(parse *command, l *jobs.List) {
	executable := exec.Command(parse.bin, parse.args...)
	executable.Stdout = os.Stdout
	executable.Stdin = os.Stdin
	executable.Start()
	if parse.background {
		l.Add(executable.Process.Pid, parse.command, executable)
	} else {
		err := executable.Wait()
		if err != nil {
			fmt.Println("go-shell: unable to run", err)
		}
	}
}

