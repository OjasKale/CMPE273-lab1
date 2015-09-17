package main


  import "testing"

  

  func TestPerimeter(t *testing.T){
    var peri float64
    c := Cir{3}
    peri = c.perimeter()
    if(peri!=18.84955592153876){
      t.Error("Expected 18.84955592153876, got ", peri)
    }
    r := Rect{5,10}
    peri = r.perimeter()
    if(peri!=30){
      t.Error("Expected 30 , got ",peri)
    }

    peri = totalPerimeter(&c,&r)
    if(peri != 48.84955592153876){
      t.Error("Expected  48.849, got ",peri)
    }
  }
