package main

import (
	"gorm.io/gen"
	"gorm.io/gen/examples/dal"
	"one-gorm-gen/cmd"
)

func init() {
	dal.DB = dal.ConnectDB(cmd.MySQLDSN).Debug()

	prepare(dal.DB) // prepare table for generate
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/onlyModel",
	})

	g.UseDB(dal.DB)

	g.GenerateAllTable()

	g.Execute()
}
