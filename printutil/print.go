package printutil

import (
	"fmt"
	"os"
)

func PrintErrorAndExit(err error) {
	fmt.Println(err)
	os.Exit(2)
}
