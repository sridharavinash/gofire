package gofire

import(
    "reflect"
    "errors"
    "strings"
    "strconv"
)


type fire struct{
  Name string
  iface interface{}
}

func Fire(i interface{}) *fire{
  return &fire{
    iface: i,
  }
}

func (f *fire) Members() []string{
  tType := reflect.TypeOf(f.iface)
  numOfExportedMethods := tType.NumMethod()
  ret := make([]string, numOfExportedMethods)
	for i := 0; i < tType.NumMethod(); i++ {
		method := tType.Method(i)
		ret[i] = method.Name
	}
  return ret
}

func (f *fire) CallMethod(methodName string, args ...interface{}) (interface{}, error){
  tType := reflect.ValueOf(f.iface)
  callMethod := tType.MethodByName(methodName)

  if callMethod.IsValid(){
    inputs := make([]reflect.Value, callMethod.Type().NumIn())
    for i, _ := range args {
      in_type := callMethod.Type().In(i)
      var z int
      switch in_type.Kind() {
      case reflect.Int:
        z ,_ = strconv.Atoi(reflect.ValueOf(args[i]).String())
      }
      inputs[i] = reflect.ValueOf(z)
    }

  return callMethod.Call(inputs)[0].Interface(), nil
}
  valid_methods := f.Members()
  return nil, errors.New("No such method: " + methodName + "\nValid Methods: " + strings.Join(valid_methods, ", ") )
}
