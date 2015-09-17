package main

import "fmt"
  import "math"

type Shape interface {
perimeter() float64
}

func totalPerimeter(shapes ...Shape) float64 {
var peri float64
for _, p := range shapes {
peri += p.perimeter()
}
return peri
}


type Cir struct {
r float64
}

func (c *Cir) perimeter() float64 {
return  2 * math.Pi * c.r
}

type Rect struct {
l, w float64
}

func (r *Rect) perimeter() float64 {
return 2 * (r.l + r.w)

}

func main() {

  var c1,r1,r2 float64
  fmt.Println("Enter value for the redius of circle")
  fmt.Scanf("%f",&c1)
    c := Cir{c1}
    fmt.Println("Enter side1 of rectangle:")
    fmt.Scanf("%f",&r1)
    fmt.Println("Enter side2 of rectangle:")
    fmt.Scanf("%f",&r2)
    r := Rect{r1,r2}
    fmt.Println("Perimeter of circle:")
    fmt.Println(c.perimeter())
    fmt.Println("Perimeter of Rectangle")
    fmt.Println(r.perimeter())
    fmt.Println("Total Perimeter")
    fmt.Println(totalPerimeter(&c, &r))
}
