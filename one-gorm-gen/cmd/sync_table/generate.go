package main

import (
	"one-gorm-gen/cmd"
	"strings"

	"gorm.io/gen"
	"gorm.io/gen/examples/dal"
	"gorm.io/gorm"
)

func init() {
	dal.DB = dal.ConnectDB(cmd.MySQLDSN).Debug()

	prepare(dal.DB) // prepare table for generate
}

// dataMap mapping relationship
var dataMap = map[string]func(gorm.ColumnType) (dataType string){
	// int mapping
	"int": func(columnType gorm.ColumnType) (dataType string) { return "int32" },

	// bool mapping
	"tinyint": func(columnType gorm.ColumnType) (dataType string) {
		ct, _ := columnType.ColumnType()
		if strings.HasPrefix(ct, "tinyint(1)") {
			return "bool"
		}
		return "byte"
	},
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./internal/syncTable",
		ModelPkgPath: "./internal/syncTable/model",

		// generate model global configuration
		FieldNullable:     true, // generate pointer when field is nullable
		FieldCoverable:    true, // generate pointer when field has default value
		FieldWithIndexTag: true, // generate with gorm index tag
		FieldWithTypeTag:  true, // generate with gorm column type tag
	})

	g.UseDB(dal.DB)

	// specify diy mapping relationship
	g.WithDataTypeMap(dataMap)

	// generate all field with json tag end with "_example"
	g.WithJSONTagNameStrategy(func(c string) string { return c + "_example" })

	mytable := g.GenerateModel("mytables")
	g.ApplyBasic(mytable)
	// g.ApplyBasic(g.GenerateAllTable()...) // generate all table in db server

	g.Execute()
}
