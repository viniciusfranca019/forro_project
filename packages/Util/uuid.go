package Util

import "github.com/google/uuid"

func GenarateUUid() string {
	id, err := uuid.NewV7()

	if err != nil {
		panic("error on new id")
	}

	return id.String()
}
