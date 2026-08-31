package cmd

import (
	"cleanSource/functions"
	"cleanSource/globalvar"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/charmbracelet/lipgloss"
)


var ignore []string
var startFile string
var includeComment bool
var listDependency bool
var (
	green = lipgloss.NewStyle().Foreground(lipgloss.Color("2"))   
	red   = lipgloss.NewStyle().Foreground(lipgloss.Color("1"))   
)

var root = &cobra.Command{
	Use: "cleansource",
	Short: "Clean every file or folder that are not needed in the project ",
	Args: cobra.ExactArgs(1),
	Version: "1.0.0",
	Run: func(cmd *cobra.Command, args []string) {

		globalvar.BasePath = args[0]
		ignore = append(ignore, ".env","*.git", ".csignore")
		data, err := os.ReadFile(filepath.Join(args[0], ".csignore"))

		if err == nil {
			for _,line:= range strings.Split(string(data),"\n"){
				if strings.HasPrefix(line,"#"){
					continue
				}
										line = strings.TrimSpace(line)

				line = strings.TrimPrefix(line,"/")
				line = strings.TrimSuffix(line,"/")
							


				ignore = append(ignore, line,)
			}
		}
		
		timeStart := time.Now()

		 functions.CheckDir(args[0],startFile,ignore,includeComment)
		 
	

			elapsedMS := time.Since(timeStart).Round(time.Millisecond).String()
			elapsedS := time.Since(timeStart).Round(time.Second).String()

			println(green.Render("Got dependencies under", elapsedMS, "or", elapsedS))
			println(len(globalvar.Dependencies), "in total")

			numberDeleted := functions.CleanSource(args[0], ignore, listDependency)

			println(red.Render("Number Deleted:", strconv.Itoa(numberDeleted)))
		 		println(green.Render("Completed under",time.Since(timeStart).Round(time.Millisecond).String(),"Milliseconds",time.Since(timeStart).Round(time.Second).String(),"Seconds"))

		
		
		

	},

}

func init(){
		root.Flags().StringSliceVar(&ignore,"ignore",[]string{},"Ignore a given list of folders or files path e.g [package.js,folder]")
				root.Flags().StringVar(&startFile,"start","","Intial file to start with e.g main.go")
				root.Flags().BoolVar(&includeComment,"includecomment",false,"Include commented imports")
				root.Flags().BoolVar(&listDependency,"list",false,"List dependencies intead delecting. This will just print the dependencies")


}

func Execute(){
	if err := root.Execute(); err != nil{
		panic(err)
	}
}