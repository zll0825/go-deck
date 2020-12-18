package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-deck/app/model"
	"go-deck/app/service"
)

var (
	dbName string
	tableName string
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "process sync notify worker",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate code running...")

		if len(tableName) == 0 {
			fmt.Println("tableName is required")
		}

		g := service.NewGenerate(model.SystemDB())

		table, err := g.GetTables(dbName)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("%+v", table)
		column, err := g.GetColumns(tableName, dbName)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Printf("%+v", column)

		g.BuildData()

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	rootCmd.PersistentFlags().StringVarP(&dbName, "dbName", "d", "", "choose database")
	rootCmd.PersistentFlags().StringVarP(&tableName, "tableName", "t", "", "choose table")
}