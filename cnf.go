package main

import (
	"fmt"
	"os"
)

func ParseCnf(filepath string) []Pro {

	pArr := []Pro{}
	p := Pro{}

	newReader, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	for {
		var Name,Path,GitPath string
		n, err := fmt.Fscanf(newReader, "%s %s %s", &Name, &GitPath, &Path)

		if err != nil || n == 0 {
			break
		}

		p.Name = Name
		p.Path = Path
		p.GitPath = GitPath

		pArr = append(pArr, p)
	}

	return pArr
}
