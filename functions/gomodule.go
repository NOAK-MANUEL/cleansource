package functions

import (
	"os"
	"path/filepath"
	"strings"
)

var goBase string

func GetGoModule(projectRoot,importLine string,) string {

	if goBase != "" {

			if strings.Contains(importLine,goBase){
							return strings.Split(importLine,goBase)[1]	

			}
			return ""
	}
	data, err := os.ReadFile(filepath.Join(filepath.Dir(projectRoot), "go.mod"))

	if err != nil {
		return ""
	}

	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "module ") {
			line= strings.TrimSpace(
				strings.TrimPrefix(line, "module "),
			)
			if strings.Contains(importLine,line){
				goBase = line
							
							
						
						
						return strings.Split(importLine,line)[1]
			}
			return ""
		}
	}

	return ""
}