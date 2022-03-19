package command

import (
	"fmt"
	"os"
)

func RemoveFiles(directory string, names []string) {
	for _, name := range names {
		var (
			path = directory + "/" + name
			err  = os.Remove(path)
		)
		if err != nil {
			fmt.Printf("failed to remove %v\n", path)
		}
	}
}
