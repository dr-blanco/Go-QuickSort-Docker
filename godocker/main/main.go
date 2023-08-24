package main

import "fmt"

func main(){
  x := []int{1,2,2,2,3,2,1,5,7}
  x_sorted := []int{}
  for i := range x{
    if len(x_sorted) <= 0 {
      x_sorted = append(x_sorted, x[i])
      continue
    }
    found := false
    for s := range x_sorted{
      if x_sorted[s] > x[i] {
        // add x[i] to x_sorted[s-1]
        x_sorted = append(x_sorted[:s], append([]int{x[i]}, x_sorted[s:]...)...)
	found = true
	break
      }
    }
    if !found {
        x_sorted = append(x_sorted, x[i])
    }
  }
  fmt.Println(x_sorted)
}
