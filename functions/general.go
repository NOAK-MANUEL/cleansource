package functions

import (
	"bufio"
	"sync"

	"cleanSource/globalvar"

	"log"

	"os"
	"path/filepath"

	"slices"
	"strings"
)

func ReadFile(path string)([]byte){
	data, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}
		
		return  data
}
func mergeDeps(base []string, add []string) []string {
    for _, d := range add {
        if !slices.Contains(base, d) {
            base = append(base, d)
        }
    }
    return base
}

func CheckDir(
    path string,
    startPath string,
    ignoreList []string,
    includeComment bool,
)  {


    // If start file is provided
    if startPath != "" {
        OpenFile(
            filepath.Join(path, startPath),
            ignoreList,
            includeComment,
        )
        return
    }

    data, err := os.ReadDir(path)

    if err != nil {
        return 
    }

    for _, entry := range data {
                newPath := filepath.Join(path, entry.Name())


        if slices.Contains(ignoreList, newPath) {
            continue
        }


        if entry.IsDir() {
             CheckDir(
                newPath,
                startPath,
                ignoreList,
                includeComment,
            )
            continue
        }

        OpenFile(
            newPath,
            ignoreList,
            includeComment,
        )
    }

    
}

func OpenFile(
    path string,
    ignore []string,
    includeComment bool,
)  {
    var wg sync.WaitGroup



    if slices.Contains(ignore, path) {
        return 
    }

    file, err := os.Open(path)
    if err != nil {
        println("couldn't open file", err)
        return 
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        ext := strings.ToLower(filepath.Ext(path))
        line := scanner.Text()

        isComment := false

        if strings.HasPrefix(line, "/*") ||
            strings.HasPrefix(line, "//") ||
            strings.HasSuffix(line, "*/") {
            isComment = true
        }

        if includeComment {
            line = removeComments(line, &isComment)
        } else if isComment {
            continue
        }

        line = CheckImport(ext, line)

        if line == "" {
            continue
        }

        var dependencyChanged bool
        if strings.HasPrefix(line,"."){
                                     line = strings.TrimPrefix(line,".")

                dotSlice := strings.Split(line,".")
                                 path =filepath.Dir(path)

                for i:=1; i<len(dotSlice);i++{
                                                         line = strings.TrimPrefix(line,".")

                    path = filepath.Dir(path)
                }

        }
        line, dependencyChanged = ParseImport(path, line)

        if line == "" {
            continue
        }

        if slices.Contains(globalvar.Dependencies, line) {
            continue
        }

        var dependencyPath string

        if dependencyChanged {
            dependencyPath = filepath.Join(
                globalvar.BasePath,
                line,
            )
        } else {
            dependencyPath = filepath.Join(
                filepath.Dir(path),
                line,
            )
        }

        subFile, err := os.Stat(dependencyPath)
       if err != nil {
            continue
        }



        globalvar.Dependencies = append(globalvar.Dependencies, dependencyPath)

        wg.Add(1)

        go func(
            newFile os.FileInfo,
            dependencyPath string,
            line string,
        ) {
            defer wg.Done()


            if subFile.IsDir() {
                 CheckDir(
                    dependencyPath,
                    "",
                    ignore,
                    includeComment,
                )
            } else {
                // println(path,dependencyPath)
                OpenFile(
                    dependencyPath,
                    ignore,
                    includeComment,
                )
            }

           

        }(
            subFile,
            dependencyPath,
            line,
        )
        
    }

    if err := scanner.Err(); err != nil {
        println("Error scanning file:", err.Error())
    }

    wg.Wait()

}
func ParseImport(path, line string) (string,bool) {
    ext := strings.ToLower(filepath.Ext(path))
     
    switch ext {
    case ".js", ".jsx", ".ts", ".tsx":
                    baseChanged :=false


        if strings.ToLower(filepath.Ext(line)) == "" {


        if strings.HasPrefix(line,"@"){
                                                     line = strings.TrimPrefix(line,"@")


                                      baseChanged = true

        }
            // sliceLine := strings.Split(line,"/")
            // nameOfFile :=  sliceLine[len(sliceLine)-1]
            fileDir := filepath.Dir(line)
            var entries []os.DirEntry; var err error;


                if baseChanged{
                    entries , err=os.ReadDir(filepath.Join(globalvar.BasePath, fileDir))
                }else{
                                        entries , err=os.ReadDir(filepath.Join(path, fileDir))

                }

            if err != nil {
                return line,baseChanged
            }

            for _,entry:= range entries{
                if entry.IsDir(){
                                  continue
 
                }
                if strings.HasPrefix(entry.Name(), filepath.Base(line)){

                            // println(entry.Name())

                                   line= filepath.Join(fileDir,entry.Name())
                                   break;

                }


            }

        }
        // if strings.ToLower(filepath.Ext(line)) == ".js" && ext != ".js" {
        //     line = strings.Split(line, ".js")[0]
        //     line = line+ext
        // }



                return line,baseChanged
    case ".dart":
        line = strings.TrimPrefix(line, "package:")
        line = strings.TrimPrefix(line, "dart:")
        return line,false

    case ".go":

        return GetGoModule(path, line),true

    case ".rs":
        line = strings.TrimPrefix(line, "crate::")
        line = strings.TrimPrefix(line, "self::")
        line = strings.TrimPrefix(line, "super::")
        return strings.ReplaceAll(line, "::", "/"),false

    default:
        return line,false
    }
}

// CheckDir scans the project directory, or starts from startPath.
// It returns the dependency files found.
