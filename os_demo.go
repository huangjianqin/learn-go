package main

import (
	"fmt"
	"os"
)

func main() {
	var err error
	//fn := "data/file.txt"

	//err = os.Mkdir("data", os.ModePerm)
	//if err != nil{
	//	fmt.Printf("%v\n", err)
	//}
	//
	//_, err = os.Create(fn)
	//if err != nil{
	//	fmt.Printf("%v\n", err)
	//}

	wd, err := os.Getwd()
	fmt.Printf("%v\n", wd)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	//err = os.WriteFile(fn, []byte("abcseqfdasda"), os.ModePerm)
	//if err != nil {
	//	fmt.Printf("%v\n", err)
	//} else{
	//	println("write finished")
	//}
	//
	//bytes, err := os.ReadFile(fn)
	//if err != nil {
	//	fmt.Printf("%v\n", err)
	//} else{
	//	println(string(bytes[:]))
	//}

	//file, err := os. OpenFile(fn, os.O_RDWR | os.O_APPEND, 0777)
	//if err != nil {
	//	fmt.Printf("%v\n", err)
	//} else{
	//	fmt.Printf("%v\n", file.Name())
	//	n, err := file.WriteString("1235412351")
	//	if err != nil {
	//		fmt.Printf("%v\n", err)
	//	} else{
	//		fmt.Printf("已写%v\n", n)
	//	}
	//}
	//defer file.Close()

	//file, err := os. OpenFile(fn, os.O_RDONLY, 0777)
	//if err != nil {
	//	fmt.Printf("%v\n", err)
	//} else{
	//	bs := make([]byte, 2)
	//	buffer := bytes.Buffer{}
	//	for{
	//		_, err := file.Read(bs)
	//		if err != nil {
	//			fmt.Printf("%v\n", err)
	//			break
	//		}
	//		//fmt.Printf("%v\n", string(bs))
	//		buffer.Write(bs)
	//	}
	//	fmt.Printf("%v\n", buffer.String())
	//}
	//defer file.Close()

	o1("data")
}

//遍历指定目录下所有文件或目录
func o1(root string) {
	dirs, err := os.ReadDir(root)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		for _, dir := range dirs {
			o2(root, dir)
		}
	}
}

func o2(root string, dir os.DirEntry) {
	fmt.Printf("%v\n", dir.Name())
	if dir.IsDir() {
		o1(root + string(os.PathSeparator) + dir.Name())
	}
}
