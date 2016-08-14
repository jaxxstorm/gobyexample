
package main // means this is an executable thing

import "fmt"
import "math" // new package!

const s string = "constant" // declare a constant

func main() {

  fmt.Println(s)

  const n = 50000000

  const d = 3e20 / n

  fmt.Println(d)

  fmt.Println(int64(d))

  fmt.Println(math.Sin(n))

}
