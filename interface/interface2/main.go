package main

func main() {
	var t1 ITom = &Tom{1, 2}
	t1.Read()
	// t1.Write()

	var t3 = new(Tom)
	t3.X = 3
	t3.Y = 4

	var t4 IBase = t3
	t4.Read()
	t4.Write()

	//只要接口中定义了响应的方法，就可以

	// var t2 IBase = Tom{1, 2}
	// t2.Write()
}
