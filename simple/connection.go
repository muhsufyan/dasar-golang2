package simple

import "fmt"

type Connection struct {
	*File
}

func (c *Connection) Close() {
	fmt.Println("Close connection", c.File.Name)
}

// provider func dibawah ini juga sama sprti provider NewFile di simple/file.go tp we tetap jelaskan konsepnya yg sama
// func provider dg cleanup func dimana provider NewConnection akan return Connection & closure func dimana closure func yg diinjek adalah func Close()
func NewConnection(file *File) (*Connection, func()) {
	connection := &Connection{File: file}
	return connection, func() {
		connection.Close()
	}
}
