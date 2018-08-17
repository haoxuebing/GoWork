package main

func main() {
	var file1 IFile = new(File)
	var file2 IReader = new(File)
	var file3 IWriter = new(File)
	var file4 ICloser = new(File)

	file1.Read(nil)
	file2.Read(nil)
	file3.Write(nil)
	file4.Close()
}
