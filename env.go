package env

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"path"
	"strings"
)

type Env struct {
	folder string
	data   map[string]string
}

func New(folder string) *Env {
	env := &Env{
		folder: folder,
		data:   map[string]string{},
	}
	file := path.Join(folder, ".env")
	fd, err := os.Open(file)
	if err == nil {
		defer fd.Close()
		buffer := bufio.NewReader(fd)
		for {
			line, _, err := buffer.ReadLine()
			if err == io.EOF {
				break
			}
			s := bytes.SplitN(line, []byte{'='}, 2)
			if len(s) < 2 {
				continue
			}
			key := string(s[0])
			value := string(s[1])
			env.data[key] = value
		}
	}
	for _, e := range os.Environ() {
		s := strings.SplitN(e, "=", 2)
		if len(s) < 2 {
			continue
		}
		env.data[s[0]] = s[1]
	}
	return env
}

func (e *Env) IsExist(key string) bool {
	_, ok := e.data[key]
	return ok
}

func (e *Env) Get(key string) string {
	if value, ok := e.data[key]; ok {
		return value
	}
	return ""
}

func (e *Env) All() map[string]string {
	return e.data
}
