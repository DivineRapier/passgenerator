package passgenerator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type storage interface {
	read()
	write()
	find()
}

var user map[string]string

func init() {
	user = make(map[string]string)
	read()
}

func read() {
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		if err.Error()[:4] != "open" {
			panic(err)
		}
	}
	data = decrypt(data)
	if err := json.Unmarshal(data, &user); err != nil {
		fmt.Println(err)
		user = make(map[string]string)
	}
}

func write() {
	data, err := json.Marshal(&user)
	if err != nil {
		panic(err)
	}
	data = encrypt(data)
	err = ioutil.WriteFile("data.json", data, os.ModeAppend)
	if err != nil {
		panic(err)
	}
}

func find(name string) string {
	if pwd, exsist := user[name]; exsist {
		return pwd
	}
	return fmt.Sprintln("password for", name, "not found")
}
