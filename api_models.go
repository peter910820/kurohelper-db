package kurohelperdb

import "time"

type UpsertUserGameErogsTXInput struct {
	UserID         string
	UserName       string
	ErogsBrandID   int
	ErogsBrandName string
	ErogsGameID    int
	ErogsGamename  string

	HasPlayed    bool
	InWish       bool
	CompleteDate time.Time
}

type BrandCountOuput struct {
	BrandID   int
	BrandName string
	Count     int
}
