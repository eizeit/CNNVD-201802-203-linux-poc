package main

/*
#cgo linux CFLAGS: -fplugin=./poc.so

void echo() {
  printf("创建文件夹/tmp/go-test\n");
}

*/
import "C"

func main() {
	C.echo()
	return
}
