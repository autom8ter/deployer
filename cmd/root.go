// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var inCluster bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "deployer",
	Long: `
ooooooooo                          o888                                              
 888    88o  ooooooooo8 ooooooooo   888   ooooooo  oooo   oooo ooooooooo8 oo oooooo  
 888    888 888oooooo8   888    888 888 888     888 888   888 888oooooo8   888    888
 888    888 888          888    888 888 888     888  888 888  888          888       
o888ooo88     88oooo888  888ooo88  o888o  88ooo88      8888     88oooo888 o888o      
                        o888                        o8o888                           

`,
	Version: "v1.0",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&inCluster, "in-cluster", "i", false, "initialize inside a kubernetes a cluster")
}
