// datatype: go-basic/datatype.go
// int datatype:
//    int8 : -128 to 127 
//    int16 : -32768 to 32767 alias rune
//    int32 : -2147483648 to 2147483647 alias int
//    int64 : -9223372036854775808 to 9223372036854775807
//
//    uint8 : 0 to 255 alias byte
//    uint16 : 0 to 65535 
//    uint32 : 0 to 4294967295 alias uint
//
//    uint64 : 0 to 18446744073709551615
//
//    floating point numbers:
//    float32 :1.18e-38 to 3.4e38
//    float64: 2.23e-308 to 1.8e308
//    complex64 : 32 bit real and imaginary numbers
//    complex128 : 64 bit real and imaginary numbers
//
//

package main

import "fmt"

func main(){
  var a int = 10
  var b int
  c:= 20
  d:= 3.14
  fmt.Println(a,b,c, d);


  var x bool = true
  y:=false
  fmt.Println(x,y);
}

