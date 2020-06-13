package main

/*
#cgo linux CFLAGS: -fplugin=./poc.so

void echo() {
  printf("测试创建文件夹/tmp/go-test，文件success.txt");
}

*/
import "C"

func main() {
	C.echo()
	return
}
