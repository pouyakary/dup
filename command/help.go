package command

import "fmt"

func DisplayHelp() {
	fmt.Println("dup: a tool to find duplicate files, by dragon's lover")
	fmt.Println("https://github.com/pouyakary/dup - kary@gnu.org\n")
	fmt.Println("usage: dup [options] <directory>")
	fmt.Println("  -h, help		displays this help message")
	fmt.Println("  -q, quite    stops the software from displaying the results")
	fmt.Println("  -r, remove   removes the duplicates from the directory")
	fmt.Println("")
}
