package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	var buf bytes.Buffer // {[] 0 0} *bytes.Buffer

	f(&buf) // <- передается ссылка на буфер

	p := &buf

	fmt.Println(*p) // {[72 101 108 108 111 32 119 111 114 108 100 10] 0 0}

	if &buf != nil {
		fmt.Printf(buf.String())
	}

	fmt.Println("Main completed")
}

func f(out *bytes.Buffer) { // принимаем указатель на buf
	fmt.Println(out) // ""
	if out != nil {  // out не будет nil, так как хранит данные
		_, err := out.Write([]byte("Hello world\n"))
		if err != nil {
			log.Fatal(err)
		}
	}
}

//package main
//
//import (
//	"bytes"
//	"fmt"
//	"io"
//	"log"
//	"reflect"
//)
//
//func main() {
//	var buf *bytes.Buffer
//
//	fmt.Println(buf, reflect.TypeOf(buf)) // <nil> *bytes.Buffer
//
//	f(buf) // <- nil | panic
//
//	if buf != nil { // false
//		fmt.Printf(buf.String())
//	}
//
//	fmt.Println("Main completed")
//}
//
//func f(out io.Writer) {
//	fmt.Println(out, reflect.TypeOf(out)) // <nil> *bytes.Buffer
//	fmt.Println(out == nil) // false
//	if out != nil {
//		_, err := out.Write([]byte("Hello world\n")) // panic!
//		if err != nil {
//			log.Fatal(err)
//		}
//	}
//}
