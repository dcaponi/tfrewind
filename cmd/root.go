/*
Copyright © 2020 OneLogin Inc dominick.caponi@onelogin.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/dcaponi/tfrewind/parsables"
	"github.com/dcaponi/tfrewind/stateparser"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tfrewind",
	Short: "A tool for converting a .tfstate file to useable HCL .tf file",
	Long: `Given a .tfstate from somewhere or from a run of terraform import,
  this tool will back-convert the .tfstate into useable HCL so you can start using Terraform quckly`,
	Run: func(cmd *cobra.Command, args []string) {
		planFile, err := os.OpenFile(filepath.Join("main.tf"), os.O_RDWR|os.O_CREATE, 0600)
		if err != nil {
			log.Fatalln("Unable to open main.tf ", err)
		}
		parsables := parsables.New()
		// grab the state from tfstate
		state := stateparser.State{}
		log.Println("Collecting State from tfstate File")
		data, err := ioutil.ReadFile(filepath.Join("terraform.tfstate"))
		if err != nil {
			planFile.Close()
			log.Fatalln("Unable to Read tfstate", err)
		}
		if err := json.Unmarshal(data, &state); err != nil {
			planFile.Close()
			log.Fatalln("Unable to Translate tfstate in Memory", err)
		}

		buffer := stateparser.ConvertTFStateToHCL(state, parsables)

		// go to the start of main.tf and overwrite whole file
		planFile.Seek(0, 0)
		_, err = planFile.Write(buffer)
		if err != nil {
			planFile.Close()
			fmt.Println("ERROR Writing Final main.tf", err)
		}

		if err := planFile.Close(); err != nil {
			fmt.Println("Problem writing file", err)
		}
	},
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
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.onelogin.json)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".onelogin" (without extension).
		viper.AddConfigPath(fmt.Sprintf("%s/.onelogin", home))
		viper.SetConfigName("profiles")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		p := filepath.Join(home, ".onelogin")
		os.Mkdir(p, 0750)
		p = filepath.Join(p, "profiles.json")
		_, err := os.Create(p)
		if err != nil {
			log.Fatalln("Unable to create config file!")
		}
		viper.ReadInConfig()
	}
}
