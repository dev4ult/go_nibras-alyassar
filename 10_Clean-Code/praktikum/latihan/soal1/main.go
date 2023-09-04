package main

type user struct {
	id int

	username int

	password int
}

// tidak camel case, sulit untuk dibaca
type userservice struct {
	// variable t mendeskripsikan apa? tidak mudah difahami
	t []user
}

// fungsi tidak menggunakan camel case, sulit dibaca
func (u userservice) getallusers() []user {
	// variable t sebagai properti objek tidak diketahui
	return u.t
}

// fungsi tidak menggunakan camel case, sulit dibaca
func (u userservice) getuserbyid(id int) user {

	// variable r mengreferensikan apa?
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}

	return user{}

}