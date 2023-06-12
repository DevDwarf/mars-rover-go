package terminal

import (
	"bufio"
	"fmt"
	"os"
)

type Terminal struct {
	Reader *bufio.Reader
}

func New() *Terminal {
	reader := bufio.NewReader(os.Stdin)

	return &Terminal{
		Reader: reader,
	}
}

func (ts *Terminal) ReadString() string {
	fmt.Print(">> ")
	cmdStr, _ := ts.Reader.ReadString('\n')

	cmdStr = cmdStr[:len(cmdStr)-1]

	return cmdStr
}

func (ts *Terminal) RenderError(err string) {
	if err == "" {
		return
	}

	fmt.Println("!! " + err + " !!")
	fmt.Println("")
}

func (ts *Terminal) RenderMultiError(errs []error) {
	if len(errs) == 0 {
		return
	}

	for k, v := range errs {
		fmt.Println("!! " + v.Error() + " !!")
		if k == len(errs)-1 {
			fmt.Println("")
		}
	}
}
