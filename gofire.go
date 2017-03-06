package gofire

import(
    "reflect"
)


type fire struct{
  Name string
  intf interface{}
}

func (f *fire) Members() []string{
  tType := reflect.TypeOf(f.intf)
  numOfExportedMethods := tType.NumMethod()
  ret := make([]string, numOfExportedMethods)
	for i := 0; i < tType.NumMethod(); i++ {
		method := tType.Method(i)
		ret[i] = method.Name
	}
  return ret
}
