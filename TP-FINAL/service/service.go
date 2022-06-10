package service

import "time"

type User struct {
	ID       int64  `json:"id"`
	Nickname string `json:"nickname"`
}

type UserService interface {
	Save(User) (User, error)
	Find() []User
}
type userService struct {
	data map[int64]User
}

func NewUserSer() UserService {
	d := map[int64]User{}
	return &userService{d}
}

func (s *userService) Save(u User) (User, error) {
	u.ID = time.Now().Unix()
	s.data[u.ID] = u
	return u, nil
}
func (s *userService) Find() []User {
	list := []User{}
	for _, v := range s.data {
		list = append(list, v)
	}
	return list
}
