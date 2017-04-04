package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/DivineRapier/passgenerator/crypto"
	"github.com/apcera/termtables"
)

type FileStorage struct {
	data map[string]string
	crypto.Crypto
}

const defaultFile = "data"

func NewFileStorage(key string) Storage {
	ret := new(FileStorage)
	ret.Crypto = crypto.NewAES(key)
	ret.Read()
	return ret
}

func (s *FileStorage) Read() {
	data, err := ioutil.ReadFile(defaultFile)
	if err != nil {
		if err.Error()[:4] != "open" {
			panic(err)
		}
	}
	data = s.Decrypt(data)
	if err := json.Unmarshal(data, &s.data); err != nil {
		fmt.Println(err)
		s.data = make(map[string]string)
	}
}

func (s *FileStorage) Write() {
	data, err := json.Marshal(&s.data)
	if err != nil {
		panic(err)
	}
	data = s.Encrypt(data)
	err = ioutil.WriteFile(defaultFile, data, os.ModeAppend)
	if err != nil {
		panic(err)
	}
}

func (s *FileStorage) Find(name string) string {
	if pwd, exsist := s.data[name]; exsist {
		return pwd
	}
	return fmt.Sprintln("password for", name, "not found")
}

func (s *FileStorage) Remove(names ...string) {
	size := len(names)
	if size == 0 {
		fmt.Println("nothing to remove")
		return
	}

	for i := 0; i < size; i++ {
		delete(s.data, names[i])
	}
}

func (s *FileStorage) Add(objs ...string) {
	size := len(objs)
	if size == 0 {
		fmt.Println("nothing to add")
		return
	}

	if size&0x01 == 1 {
		fmt.Println("count of input might be mul of 2")
		return
	}
	for i := 0; i < size; i += 2 {
		s.data[objs[i]] = objs[i+1]
	}
}

func (s *FileStorage) List() string {
	table := termtables.CreateTable()

	table.AddHeaders("username", "password")
	for username, password := range s.data {
		table.AddRow(username, password)
	}

	return table.Render()
}
