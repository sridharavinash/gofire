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

func (c *Calculator) AddTwo(n2 int, n3 int) int{
  return c.n1 + n2 + n3
}

func TestSimpleCalc(t *testing.T){
  f := Fire(&Calculator{})
  methods := f.Members()
  expected := []string{"Add", "AddTwo", "Double", "Square"}
  if !reflect.DeepEqual(methods,expected)  {
    t.Error("Expected [Add AddTwo Double Square], got ", methods)
  }
}

func TestSimpleCalcDouble(t *testing.T){
  f := Fire(&Calculator{n1: 10,})

  result, _ := f.CallMethod("Square")
  if result != 100  {
    t.Error("Expected 100, got ", result)
  }
}

func TestSimpleCalcAdd(t *testing.T){
  f := Fire(&Calculator{n1: 10,},)

  result, _ := f.CallMethod("Add",20)
  if result != 30  {
    t.Error("Expected 30, got ", result)
  }
}

func TestSimpleCalcAddTwo(t *testing.T){
  f := Fire(&Calculator{n1: 10,},)

  result, _ := f.CallMethod("AddTwo",20,40)
  if result != 70  {
    t.Error("Expected 70, got ", result)
  }
}

func TestSimpleCalcInvalidMethod(t *testing.T){
  f := Fire(&Calculator{})

  result, err := f.CallMethod("Subtract",20)
  if err == nil {
    t.Error("Expected error, got ", result)
  }
}
