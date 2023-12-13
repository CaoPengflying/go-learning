package main

import (
	"context"
	"gorm.io/gen"
	"gorm.io/gen/examples/dal"
	"one-gorm-gen/cmd"
	"one-gorm-gen/method"
)

func init() {
	dal.DB = dal.ConnectDB(cmd.MySQLDSN).Debug()
}

// generate code
func main() {
	db := dal.DB

	cfg := gen.Config{
		OutPath:       "./internal/WithDefaultQuery/query",
		ModelPkgPath:  "./internal/WithDefaultQuery/model",
		Mode:          gen.WithDefaultQuery,
		FieldNullable: true,
	}

	g := gen.NewGenerator(cfg)

	g.UseDB(db.WithContext(context.Background()))

	//g.GenerateAllTable()

	/*	g.ApplyBasic(
			g.GenerateModel("mytables", gen.FieldJSONTagWithNS(camelColumn)),
		)
		g.ApplyInterface(func(testIF method.TrimTest) {
		}, g.GenerateModel("mytables"))*/

	g.ApplyBasic(g.GenerateAllTable()...)
	g.ApplyInterface(func(testTrim method.TrimTest, selectMethod method.SelectMethod) {
	}, g.GenerateModel("mytables"))

	g.Execute()
}

// camelColumn 将下划线转为驼峰
func camelColumn(columnName string) (tagContent string) {
	data := make([]byte, 0, len(columnName))

	columnLen := len(columnName) - 1
	transFlag := false
	for i := 0; i <= columnLen; i++ {
		d := columnName[i]
		if d == '_' {
			transFlag = true
			continue
		}
		if d >= 'a' && d <= 'z' && transFlag {
			d = d - 32
			transFlag = false
		}
		data = append(data, d)
	}
	return string(data)
}
