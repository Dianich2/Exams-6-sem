package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(n int, mas []int) string {
	t := false
	for i := 0; i < n; i++ {
		if mas[i] >= 3 {
			return "YES"
		}
		if mas[i] == 2 && t {
			return "YES"
		}
		if mas[i] == 2 {
			t = true
		}
	}
	return "NO"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		mas := make([]int, n)
		for j := 0; j < n; j++ {
			scanner.Scan()
			mas[j], _ = strconv.Atoi(scanner.Text())
		}
		ans := solve(n, mas)
		fmt.Fprintln(out, ans)
	}
}
