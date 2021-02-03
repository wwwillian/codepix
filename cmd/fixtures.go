/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"github.com/wwwillian/codepix-go/domain/model"
	"github.com/wwwillian/codepix-go/infrastructure/db"
	"github.com/wwwillian/codepix-go/infrastructure/repository"
	"os"

	"github.com/spf13/cobra"
)

// fixturesCmd represents the fixtures command
var fixturesCmd = &cobra.Command{
	Use:   "fixtures",
	Short: "Run fixture for fake data generation",
	Run: func(cmd *cobra.Command, args []string) {

		database := db.ConnectDB(os.Getenv("env"))
		defer database.Close()
		pixRepository := repository.PixKeyRepositoryDb{Db: database}

		bankNubank, _ := model.NewBank("111", "Nubank")
		pixRepository.AddBank(bankNubank)

		bankItau, _ := model.NewBank("222", "Itau")
		pixRepository.AddBank(bankItau)

		accountA, _ := model.NewAccount(bankNubank, "1010", "Willian A")
		accountA.ID = "6e4635ce-88d1-4e58-9597-d13fc446ee47"
		pixRepository.AddAccount(accountA)

		accountB, _ := model.NewAccount(bankNubank, "2020", "Maria B")
		accountB.ID = "51a720b2-5144-4d7f-921d-57023b1e24c1"
		pixRepository.AddAccount(accountB)

		accountC, _ := model.NewAccount(bankItau, "3030", "User CTER 1")
		accountC.ID = "103cc632-78e7-4476-ab63-d5ad3a562d26"
		pixRepository.AddAccount(accountC)

		accountD, _ := model.NewAccount(bankItau, "4040", "User CTER 2")
		accountD.ID = "463b1b2a-b5fa-4b88-9c31-e5c894a20ae3"
		pixRepository.AddAccount(accountD)

	},
}

func init() {
	rootCmd.AddCommand(fixturesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fixturesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fixturesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
