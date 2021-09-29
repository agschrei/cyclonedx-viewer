/*
Copyright © 2021 Alexander Gschrei

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

	"github.com/agschrei/cyclonedx-viewer/models"
	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var (
	cfgFile string
	bomPath string
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "cyclonedx-viewer",
		Short: "a small web server that allows to visually inspect CycloneDx BOMs",
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) { models.ParseCycloneDxBom(bomPath) },
	}
)

func PrintBanner() {
	banner := `
|***************************************************************************************|
|     ______           __                 ____          _    ___                        |
|    / ____/_  _______/ /___  ____  ___  / __ \_  __   | |  / (_)__ _      _____  _____ |
|   / /   / / / / ___/ / __ \/ __ \/ _ \/ / / / |/_/   | | / / / _ \ | /| / / _ \/ ___/ |
|  / /___/ /_/ / /__/ / /_/ / / / /  __/ /_/ />  <     | |/ / /  __/ |/ |/ /  __/ /     |
|  \____/\__, /\___/_/\____/_/ /_/\___/_____/_/|_|     |___/_/\___/|__/|__/\___/_/      |
|       /____/                                                                          |
|                                                                                       |
*****************************************************************************************                                                                                                                                                                                      
 `
	color.Green(banner)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	PrintBanner()
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cyclonedx-viewer.yaml)")
	rootCmd.PersistentFlags().StringVarP(&bomPath, "bom-file", "f", "", "provide the path to the BOM to operate on")
	rootCmd.MarkPersistentFlagRequired("bom-file")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cyclonedx-viewer" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cyclonedx-viewer")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
