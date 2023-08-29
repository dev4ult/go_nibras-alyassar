package main

import (
	"fmt"
	"sync"
)

type form struct {
	name        string
	email       string
	phoneNumber string
	m           sync.Mutex
}

func (f *form) setForm(name, email, phoneNumber string) {
	f.m.Lock()

	f.name = name
	f.email = email
	f.phoneNumber = phoneNumber

	f.m.Unlock()
}

func (f *form) getFormData() map[string]string {
	f.m.Lock()

	defer f.m.Unlock()
	
	return map[string]string {"nama": f.name, "email": f.email, "nomor telepon": f.phoneNumber}
}

func main() {
	userForm := new(form)
	userForm.setForm("Nibras", "nibras@example.com", "08XXXXXXXXXX")
	fmt.Println(userForm.getFormData())
}