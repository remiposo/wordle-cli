package presentation

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func ReadAnswer(count int) string {
	fmt.Printf("%d - Submit a word:\n", count)
	scanner.Scan()
	return scanner.Text()
}
