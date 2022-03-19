package detector

import "sort"

type Memory map[string][]string

func (memory *Memory) Add(file string, hash string) {
	if _, ok := (*memory)[hash]; !ok {
		(*memory)[hash] = []string{}
	}
	(*memory)[hash] = append((*memory)[hash], file)
}

func (memory *Memory) ListDouplicateFiles() []string {
	results := []string{}
	for _, files := range *memory {
		size := len(files)
		if size > 1 {
			for i := 1; i < size; i++ {
				results = append(results, files[i])
			}
		}
	}
	sort.Strings(results)
	return results
}
