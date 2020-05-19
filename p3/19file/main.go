package main

func main() {
	/*
		文件操作：
		1.路径：
			相对路径：relative
				ab.txt
				相对于当前工程
			绝对路径：absolute
				/Users/qiuwen/Nutstore Files/我的坚果云/go_src/q107580018/studyGo/p3/19file/aaa.txt
	*/
	/*	// 1.路径
		fileName1 := "/Users/qiuwen/Nutstore Files/我的坚果云/go_src/q107580018/studyGo/p3/19file/aaa.txt"
		fileName2 := "aaa.txt"
		fmt.Println(filepath.IsAbs(fileName1)) // true
		fmt.Println(filepath.IsAbs(fileName2)) // false
		fmt.Println(filepath.Abs(fileName1))
		fmt.Println(filepath.Abs(fileName2)) // /Users/qiuwen/Nutstore Files/我的坚果云/go_src/q107580018/studyGo/p3/19file/aaa.txt <nil>
		fmt.Println("获取父目录", path.Join(fileName1, ".."))
		// 2.创建目录
		err := os.Mkdir("/Users/qiuwen/Nutstore Files/我的坚果云/go_src/q107580018/studyGo/p3/19file/aaa", os.ModePerm)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("文件夹创建成功")
		}
		// 3.创建多层文件夹
		err = os.MkdirAll("/Users/qiuwen/Nutstore Files/我的坚果云/go_src/q107580018/studyGo/p3/19file/bbb/ccc/ddd", os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("多层文件夹创建成功")
	*/
}
