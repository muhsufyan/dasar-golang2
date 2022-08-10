package simple

import "fmt"

type File struct {
	Name string
}

func (f *File) Close() {
	fmt.Println("Close file", f.Name)
}

// provider func dg cleanup func dimana provider NewFile akan return File & closure func dimana closure func yg diinjek adalah func Close()
// jika ingin ada tambahan error maka deklarasi func providernya tinggal tambah error stlh closure func. ex
// func NewFile(name string) (*File, func(), error) dan return nya juga tambah errror
func NewFile(name string) (*File, func()) {
	file := &File{Name: name}
	return file, func() {
		file.Close()
	}
}
