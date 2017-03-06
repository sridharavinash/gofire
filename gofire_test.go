package gofire

import (
  "testing"
  "reflect"
)

type Calculator struct {
  n1 int
}

func (c *Calculator) Double() int{
  return c.n1 * 2
}

func (c *Calculator) Square() int{
  return c.n1 * c.n1
}

func (c *Calculator) Add(n2 int) int{
  return c.n1 + n2
}

func TestSimpleCalc(t *testing.T){
  f := &fire {
    iface: &Calculator{},
  }
  methods := f.Members()
  expected := []string{"Add", "Double", "Square"}
  if !reflect.DeepEqual(methods,expected)  {
    t.Error("Expected [Add Double Square], got ", methods)
  }
}

func TestSimpleCalcDouble(t *testing.T){
  f := &fire {
    iface: &Calculator{n1: 10,},
  }

  result := f.CallMethod("Square")
  if result != 100  {
    t.Error("Expected 100, got ", result)
  }
}

func TestSimpleCalcAdd(t *testing.T){
  f := &fire {
    iface: &Calculator{n1: 10,},
  }

  result := f.CallMethod("Add",20)
  if result != 30  {
    t.Error("Expected 30, got ", result)
  }
}
