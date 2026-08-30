package functions

import (
	"cleanSource/globalvar"
	"path/filepath"
	"regexp"
	"strings"
)


func jsConfig(line string) string{
	    if !strings.HasPrefix(line,".") && !strings.HasPrefix(line,"@")&& !strings.HasPrefix(line,"/") {
	         	return""
	         }
	         return line

}

func CheckImport( ext,line string) string {

	switch ext {

	case ".js", ".jsx",".ts",".tsx":
		importReg,err := regexp.Compile(
    `^\s*import(?:.*?\sfrom)?\s*["']([^"']+)["']`,
)

	if err == nil {
		substring := importReg.FindStringSubmatch(line)
		if len(substring)>0{
			return jsConfig(substring[1])
					}
									
}
			


sideEffectImportReg,err := regexp.Compile(
    `^\s*(?:.*?\s)from?\s*["']([^"']+)["']`,
)
if err == nil {
		substring := sideEffectImportReg.FindStringSubmatch(line)
		if len(substring)>0{
			return jsConfig(substring[1])
					}
									
}


dynamicImportReg,err := regexp.Compile(
    `\bimport\s*\(\s*["']([^"']+)["']\s*\)`,
)

if err == nil {
		substring := dynamicImportReg.FindStringSubmatch(line)
		if len(substring)>0{
			return jsConfig(substring[1])
					}
									
}

requireReg,err := regexp.Compile(
    `require\s*\(\s*["']([^"']+)["']`,
)
if err == nil {
		substring := requireReg.FindStringSubmatch(line)
		if len(substring)>0{
			return jsConfig(substring[1])
					}
									
}
	case ".go":

		
		dynamicImportReg,err := regexp.Compile(
    `\bimport\s*\(\s*["']([^"']+)["']\s*\)`,
)

		if err == nil {
			substring := dynamicImportReg.FindStringSubmatch(line)
			if len(substring)>0{
				return substring[1]
			}
									
		}

		line =GetGoModule(filepath.Join(globalvar.BasePath,ext),line)
		if line != ""{
			if strings.HasSuffix(line,")"){
				line = strings.TrimSuffix(line,")")
			}
			line = strings.TrimSuffix(line,"\"")
			return filepath.Join(globalvar.BasePath,line)
		}

	case ".py":
		importReg, err := regexp.Compile(`^\s*import\s+([a-zA-Z_][\w.]*)`)
		if err == nil {
			if m := importReg.FindStringSubmatch(line); len(m) > 1 {
				return m[1]
			}
		}

		fromReg, err := regexp.Compile(`^\s*from\s+([a-zA-Z_][\w.]*)\s+import\s+`)
		if err == nil {
			if m := fromReg.FindStringSubmatch(line); len(m) > 1 {
				return m[1]
			}
		}

	case ".java":
		importReg, err := regexp.Compile(
			`^\s*import\s+(?:static\s+)?([\w.]+)`,
		)
		if err == nil {
			if m := importReg.FindStringSubmatch(line); len(m) > 1 {
				return m[1]
			}
		}

	case ".kt", ".kts":
		importReg, err := regexp.Compile(
			`^\s*import\s+([\w.]+)`,
		)
		if err == nil {
			if m := importReg.FindStringSubmatch(line); len(m) > 1 {
				return m[1]
			}
		}

	case ".c", ".h", ".cpp", ".cc", ".cxx", ".hpp":
		includeReg, err := regexp.Compile(
			`^\s*#\s*include\s*[<"]([^>"]+)[>"]`,
		)
		if err == nil {
			if m := includeReg.FindStringSubmatch(line); len(m) > 1 {
				return m[1]
			}
		}

	case ".cs":
		usingReg, err := regexp.Compile(
			`^\s*using\s+(?:static\s+)?([\w.]+)`,
		)
		if err == nil {
			if m := usingReg.FindStringSubmatch(line); len(m) > 1 {
				return m[1]
			}
		}

	case ".rs":
		useReg, err := regexp.Compile(
			`^\s*use\s+([^;]+)`,
		)
		if err == nil {
			if m := useReg.FindStringSubmatch(line); len(m) > 1 {
				return strings.TrimSpace(m[1])
			}
		}

		externReg, err := regexp.Compile(
			`^\s*extern\s+crate\s+([\w_]+)`,
		)
		if err == nil {
			if m := externReg.FindStringSubmatch(line); len(m) > 1 {
				return m[1]
			}
		}

	case ".php":
		useReg, err := regexp.Compile(
			`^\s*use\s+([^;]+)`,
		)
		if err == nil {
			if m := useReg.FindStringSubmatch(line); len(m) > 1 {
				return strings.TrimSpace(m[1])
			}
		}

		requireReg, err := regexp.Compile(
			`^\s*(?:require|require_once|include|include_once)\s*[\(\s]*["']([^"']+)["']`,
		)
		if err == nil {
			if m := requireReg.FindStringSubmatch(line); len(m) > 1 {
				return m[1]
			}
		}

	case ".rb":
		requireReg, err := regexp.Compile(
			`^\s*require\s+["']([^"']+)["']`,
		)
		if err == nil {
			if m := requireReg.FindStringSubmatch(line); len(m) > 1 {
				return m[1]
			}
		}

		requireRelativeReg, err := regexp.Compile(
			`^\s*require_relative\s+["']([^"']+)["']`,
		)
		if err == nil {
			if m := requireRelativeReg.FindStringSubmatch(line); len(m) > 1 {
				return m[1]
			}
		}

	case ".swift":
		importReg, err := regexp.Compile(
			`^\s*import\s+(?:class|struct|enum|protocol|func|var|let|typealias)?\s*([\w.]+)`,
		)
		if err == nil {
			if m := importReg.FindStringSubmatch(line); len(m) > 1 {
				return m[1]
			}
		}

	case ".dart":

		importReg, err := regexp.Compile(
			`^\s*import\s+["']([^"']+)["']`,
		)

		if err == nil {
			if m := importReg.FindStringSubmatch(line); len(m) > 1 {
				return m[1]
			}
		}

	case ".scala":
		importReg, err := regexp.Compile(
			`^\s*import\s+([\w.]+)`,
		)
		if err == nil {
			if m := importReg.FindStringSubmatch(line); len(m) > 1 {
				return m[1]
			}
		}

	case ".lua":
		requireReg, err := regexp.Compile(
			`^\s*(?:local\s+\w+\s*=\s*)?require\s*\(?\s*["']([^"']+)["']`,
		)
		if err == nil {
			if m := requireReg.FindStringSubmatch(line); len(m) > 1 {
				return m[1]
			}
		}

	case ".pl", ".pm":
		useReg, err := regexp.Compile(
			`^\s*use\s+([\w:]+)`,
		)
		if err == nil {
			if m := useReg.FindStringSubmatch(line); len(m) > 1 {
				return m[1]
			}
		}

		requireReg, err := regexp.Compile(
			`^\s*require\s+["']([^"']+)["']`,
		)
		if err == nil {
			if m := requireReg.FindStringSubmatch(line); len(m) > 1 {
				return m[1]
			}
		}

	case ".ex", ".exs":
		elixirReg, err := regexp.Compile(
			`^\s*(?:alias|import|require)\s+([A-Za-z0-9_.]+)`,
		)
		if err == nil {
			if m := elixirReg.FindStringSubmatch(line); len(m) > 1 {
				return m[1]
			}
		}

	case ".hs":
		importReg, err := regexp.Compile(
			`^\s*import\s+(?:qualified\s+)?([A-Za-z0-9_.]+)`,
		)
		if err == nil {
			if m := importReg.FindStringSubmatch(line); len(m) > 1 {
				return m[1]
			}
		}
	}

	return ""
}