package entropy

import "os"

func WriteToFile(filename string, line string) {
    // os.Mkdir("test", 0777)
	f, createErr := os.Create(filename)
    check(createErr)
    defer f.Close()
    _, writeErr := f.WriteString(line + "\n")
    check(writeErr)
    f.Sync()
    f.Close()
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
