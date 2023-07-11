package utils

import "github.com/google/uuid"

var UUID *uuid.UUID

func InitUUID() {
	uid, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	UUID = &uid
}
