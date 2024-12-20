package services

import (
	"golang.org/x/sys/windows"
	"os"
	"os/user"
)

type UserService struct {
}

func (*UserService) GetAllUsers() (*user.User, error) {
	_, err := os.OpenFile("q", 1, windows.FILE_PIPE_BYTE_STREAM_MODE)
	return &user.User{}, err
}
