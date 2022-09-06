package filedemo

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func FileTest() {
	// openFile()
	// readFile()
	// readFileByByte()
	createFile()
}

func openFile() {
	// Open函数只能打开,不能读取
	fp, err := os.Open("E://temp//test.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fp)
	}
	defer func() {
		err := fp.Close()
		if err != nil {
			fmt.Println("close", err)
		}
	}()
}

func readFile() {
	fp, _ := os.Open("E://temp//test.txt")
	buf := make([]byte, 20)
	for {
		// 不带缓冲区
		count, err := fp.Read(buf)
		fmt.Print(string(buf[:count]))
		if err == io.EOF {
			break
		}
	}
	defer func() {
		err := fp.Close()
		if err != nil {
			fmt.Println("close", err)
		}
	}()
}
func readFileByByte() {
	fmt.Println()
	fp, _ := os.Open("E://temp//test.txt")
	r := bufio.NewReader(fp)
	for {
		// 带缓冲区
		buf, err := r.ReadString('\n')
		fmt.Println(string(buf))
		if err == io.EOF {
			break
		}
	}
	defer func() {
		err := fp.Close()
		if err != nil {
			fmt.Println("close", err)
		}
	}()
}

func createFile() {
	info, err := os.Stat("E://temp//test2.txt")
	if err == nil {
		fmt.Println("file", info)
	}

	fp, _ := os.Create("E://temp//test2.txt")
	// 2.关闭打开的文件
	defer func() {
		err := fp.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	// 2.往文件中写入数据
	// 注意: Windows换行是\r\n
	bytes := []byte{'l', 'n', 'j', '\r', '\n'}
	fp.Write(bytes)
	fp.WriteString("嘎嘎嘎\r\n")
	fp.WriteString("kkk\r\n")
}
