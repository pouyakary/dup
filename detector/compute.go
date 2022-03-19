package detector

import (
	"dup/terminal"
	"fmt"
	"io/ioutil"
	"math"
	"runtime"
)

type resultHash struct {
	hash string
	name string
}

func detectListOfDuplicates(directory string) []string {
	list := []string{}
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		list = append(list, file.Name())
	}
	return list
}

func FindDouplicateFiles(directory string, term *terminal.Terminal) []string {
	var (
		filesList        = detectListOfDuplicates(directory)
		size             = len(filesList)
		memory           = &Memory{}
		availableCores   = runtime.NumCPU()
		finishedRoutines = 0
		portionSize      = int(math.Floor(float64(size) / float64(availableCores)))
		post             = make(chan *resultHash)
		resultCount      = 0
	)
	for portionIndex := 0; portionIndex < availableCores; portionIndex++ {
		var (
			currentIndex       = portionIndex
			nextIndex          = portionIndex + 1
			arrayStartingIndex = currentIndex * portionSize
			arrayStopIndex     = nextIndex * portionSize
		)
		if (portionIndex + 1) == availableCores {
			arrayStopIndex = size
		}
		go func() {
			for index := arrayStartingIndex; index < arrayStopIndex; index++ {
				var (
					name     = filesList[index]
					path     = directory + "/" + name
					md5, err = ComputeFileMD5(path)
				)
				if err != nil {
					fmt.Printf("failed to compute md5 for: %s\n", name)
				} else if md5 != "" {
					post <- &resultHash{name: name, hash: md5}
				}
			}
			post <- nil
		}()
	}

	for {
		result := <-post
		if result == nil {
			finishedRoutines++
			if finishedRoutines == availableCores {
				break
			}
		} else {
			memory.Add(result.name, result.hash)

			percent := int(math.Floor(float64(resultCount) / float64(size) * 100))
			term.MoveToX(0)
			term.Print(fmt.Sprintf("%v%%", percent))
			term.MoveToX(100000)
			resultCount++
		}
	}

	term.MoveToX(0)

	return memory.ListDouplicateFiles()
}
