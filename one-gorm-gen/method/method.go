// Package method
// @Description method
// @Author caopengfei 2023/12/13 18:41
package method

import "gorm.io/gen"

type SelectMethod interface {
	// FindByID
	//
	// select * from users where id=@id
	FindByID(id int) gen.T

	// LikeSearch
	//
	// SELECT * FROM @@table where name LIKE concat('%',@name,'%')
	LikeSearch(name string) *gen.T

	// InSearch
	//
	// select * from @@table where name in @names
	InSearch(names []string) []*gen.T

	// ColumnSearch
	//
	// select * from @@table where @@name in @names
	ColumnSearch(name string, names []string) []*gen.T
}

type TrimTest interface {
	// TestTrim
	//
	// select * from @@table where
	// {{trim}}
	// {{for _,name :=range list}}
	// name = @name or
	// {{end}}
	// {{end}}
	TestTrim(list []string) []gen.T

	// TestTrimInWhere
	//
	// select * from @@table
	// {{where}}
	// {{trim}}
	// {{for _,name :=range list}}
	// name = @name or
	// {{end}}
	// {{end}}
	// {{end}}
	TestTrimInWhere(list []string) []gen.T

	// TestInsert
	//
	// insert into users (name,age) values
	// {{trim}}
	// {{for key,value :=range data}}
	// (@key,@value),
	// {{end}}
	// {{end}}
	TestInsert(data gen.M) error
}
