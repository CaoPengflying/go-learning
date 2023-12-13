package main

import (
	"context"
	"git.100tal.com/wangxiao_tech/one/frame"
	"gorm.io/gen"
	"gorm.io/gorm"
	"one-gorm-gen/method"
)

// generate code
func main() {
	db := GetGenerateDB()

	/*	cfg := gen.Config{
		OutPath: "./internal/onlyModel/model",
	}*/

	/*	cfg := gen.Config{
		OutPath:       "./internal/WithDefaultQuery/query",
		ModelPkgPath:  "./internal/WithDefaultQuery/model",
		Mode:          gen.WithDefaultQuery,
		FieldNullable: true,
	}*/

	/*	cfg := gen.Config{
		OutPath:       "./internal/WithoutContext/gen",
		ModelPkgPath:  "./internal/WithoutContext/model",
		Mode:          gen.WithoutContext,
		FieldNullable: true,
	}*/

	cfg := gen.Config{
		OutPath:       "./internal/withQueryInterface/gen",
		ModelPkgPath:  "./internal/withQueryInterface/model",
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	}

	g := gen.NewGenerator(cfg)

	g.UseDB(db.WithContext(context.Background()))

	//g.GenerateAllTable()

	/*	g.ApplyBasic(
			g.GenerateModel("task", gen.FieldJSONTagWithNS(camelColumn)),
			g.GenerateModel("step", gen.FieldJSONTagWithNS(camelColumn)),
			g.GenerateModel("plugin", gen.FieldJSONTagWithNS(camelColumn)),
			g.GenerateModel("task_step", gen.FieldJSONTagWithNS(camelColumn)),
			g.GenerateModel("step_plugin", gen.FieldJSONTagWithNS(camelColumn)),
			g.GenerateModel("step_log", gen.FieldJSONTagWithNS(camelColumn)),
			g.GenerateModel("task_log", gen.FieldJSONTagWithNS(camelColumn)),
		)
		g.ApplyInterface(func(testIF method.TrimTest) {
		}, g.GenerateModel("task"))*/

	g.ApplyBasic(g.GenerateAllTable()...)
	g.ApplyInterface(func(testTrim method.TrimTest, selectMethod method.SelectMethod) {
	}, g.GenerateModel("task"))

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

const HuiyinDb = "huiyin_writer"

// 获取生成文件的db
func GetGenerateDB() *gorm.DB {
	bootstrap, err := frame.CreateBootstrap()
	if err != nil {
		panic(err)
	}
	err = bootstrap.LightweightStart(frame.RegisterGormRouter())
	if err != nil {
		panic(err)
	}
	router, err := bootstrap.GetGormRouter()
	if err != nil {
		return nil
	}
	db, exist := router.RouteDB(HuiyinDb)
	if !exist {
		panic("not exist route db")
	}
	return db.WithContext(context.Background())
}
