package reflect

import (
	"errors"
	"reflect"
	"testing"
)

func TestBasicReflect(t *testing.T) {
	basic := 1
	CheckType(&basic)
	t.Log(reflect.TypeOf(basic), reflect.ValueOf(basic))
}

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)

	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		println("float")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		println("integer")
	default:
		println("unKnown")
	}
}

type Employee struct {
	Name string `format:normal`

	Age int
}

func (e *Employee) UpdateEmployeeAge(age int) {
	e.Age = age
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{"cpf", 27}

	t.Logf("Name: value(%[1]v), Type(%[1]T)", reflect.ValueOf(*e).FieldByName("name"))

	if nameFiled, ok := reflect.TypeOf(*e).FieldByName("name"); !ok {
		t.Error("failed to get name field")
	} else {
		t.Log("Tag:format", nameFiled.Tag.Get("format"))
	}
	reflect.ValueOf(e).MethodByName("UpdateEmployeeAge").Call([]reflect.Value{reflect.ValueOf(1)})
	t.Log("update age:", e)
}

func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 3: "three"}
	//t.Log(a == b)
	t.Log(reflect.DeepEqual(a, b))

}

func fillBySetting(st interface{}, settings map[string]interface{}) error {
	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
			return errors.New("the first param should be a pointer to the struct type")
		}
	}
	if settings == nil {
		return errors.New("settings can't be nil")
	}

	var (
		field reflect.StructField
		ok    bool
	)
	for k, v := range settings {
		if field, ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok {
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(st)
			vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}
	return nil
}

type Consumer struct {
	Name string
	Age  int
}
//使用 go的反射时，字段名必须大写，大写表示公有的

func TestFillSetting(t *testing.T) {
	settings := map[string]interface{}{"Name": "zzc", "Age": 40}
	consumer := Consumer{}
	err := fillBySetting(&consumer, settings)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(consumer)
	}
	e := new(Employee)
	if err = fillBySetting(e, settings); err != nil {
		t.Error(err)
	}else {
		t.Log(e)
	}
}
