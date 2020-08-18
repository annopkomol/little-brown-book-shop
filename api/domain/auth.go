package domain

type AuthService interface {
	Login(username string, password string) (success bool, err error)
	GetPosID(username string) (posID int, err error)
}

type AuthRepository interface {
	GetPassword(username string) (hashed string, err error)
	GetPosID(username string) (posID int, err error)
}
