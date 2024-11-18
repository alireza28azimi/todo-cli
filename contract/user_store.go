package contract

import "my-go-project/entity"

type UserWriteStore interface {
	Save(u entity.User)
}

type UserReadStore interface {
	Load() []entity.User
}
