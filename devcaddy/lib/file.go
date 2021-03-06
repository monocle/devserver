package lib

import (
	"io/ioutil"
	"log"
	"strings"
)

const FILE_PATH_SPLITTER = "__SERVER_FILE_PATH__="
const (
	CREATE FileOp = 1 << iota
	WRITE
	REMOVE
	RENAME
	CHMOD
	LOG
	ERROR
)

type FileOp uint32

func NewFile(e *Event) *File {
	f := File{Name: e.Name(), Op: e.Op, Error: e.Error}
	if f.IsDeleted() || f.IsError() {
		return &f
	}

	b, err := ioutil.ReadFile(f.Name)
	if err != nil {
		log.Fatalln(f.Name, err)
	}

	f.Content = string(b)
	return &f
}

func NewFileWithContent(name, content string, op FileOp) *File {
	return &File{Name: name, Content: content, Op: op}
}

func NewFileFromCommand(oFile *File, output []byte, err error, pluginName string) *File {
	f := &File{
		Name:       oFile.Name,
		Op:         oFile.Op,
		PluginName: pluginName,
	}

	if err != nil {
		f.Error = err
		f.Op = ERROR
	}

	split := strings.Split(string(output), FILE_PATH_SPLITTER)

	if len(split) > 1 {
		f.Name = strings.TrimSpace(split[1])
	}

	f.Content = split[0]
	return f
}

type FileConfig struct {
	Dir, Ext    string
	Files       []string
	PluginNames []string `json:"plugins"`
}

type File struct {
	FileConfig
	Name       string
	Content    string
	Type       string
	Error      error
	Op         FileOp
	PluginName string
}

func (f *File) IsDeleted() bool {
	return f.Op == REMOVE || f.Op == RENAME
}

func (f *File) IsError() bool {
	return f.Op == ERROR
}
