package storage

import "api/types"

type Storage interface {
	Get(int) *types.User
}
