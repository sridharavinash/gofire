package gofire

import(
    "reflect"
    "errors"
    "strings"
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
    inputs := make([]reflect.Value, len(args))
    for i, _ := range args {
      inputs[i] = reflect.ValueOf(args[i])
    }
    return callMethod.Call(inputs)[0].Interface(), nil
  }
  valid_methods := f.Members()
  return nil, errors.New("No such method: " + methodName + "\nValid Methods: " + strings.Join(valid_methods, ", ") )
}
