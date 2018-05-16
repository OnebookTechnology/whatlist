package _interface

type ServerDB interface {
	InitialDB(confPath string, tagName string) error
}
