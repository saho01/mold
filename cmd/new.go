/*
Copyright © 2020 Seth Aho <saho01@hotmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/saho01/mold/pkg/parse"
	"github.com/saho01/mold/pkg/project"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var (
	ignore  bool
	read_me bool
	newCmd  = &cobra.Command{
		Use:   "new",
		Short: "Create new project structure",
		Long:  `Create new projct structure. Will prompt for project name and location`,
		Run: func(cmd *cobra.Command, args []string) {
			location, err := parse.Location()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			err = project.CreateDir(location)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if ignore {
				fmt.Println("Adding gitignore")
				err = project.CreateFile(location, ".gitignore")
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}
			if read_me {
				fmt.Println("Adding README")
				err = project.CreateFile(location, "README.md")
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}
			fmt.Println("Done")
		},
	}
)

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	newCmd.Flags().BoolVarP(&ignore, "ignore", "i", false, "Include gitignore")
	newCmd.Flags().BoolVarP(&read_me, "me", "m", false, "Include readme")
}
