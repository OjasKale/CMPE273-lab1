package main

import "testing"


func TestFibc(t *testing.T)  {
  var val uint
  val = Fibc(3)
  if(val!=2){
    t.Error("Expected 2,got ",val)
  }
}
