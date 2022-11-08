package domain

type User struct {
	id     int
	userId string
	name   string
}

func NewUser(id int, userId string, name string) *User {
	return &User{id: id, userId: userId, name: name}
}

func (u *User) Id() int {
	return u.id
}

func (u *User) SetId(id int) {
	u.id = id
}

func (u *User) UserId() string {
	return u.userId
}

func (u *User) SetUserId(userId string) {
	u.userId = userId
}

func (u *User) Name() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}
