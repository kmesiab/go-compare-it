package types

import "github.com/google/uuid"

type Product struct {
	id          uuid.UUID
	title       string
	description string
}

type Store struct {
	id          uuid.UUID
	title       string
	description string
}

type Review struct {
	id          uuid.UUID
	title       string
	description string
}

type User struct {
	id        uuid.UUID
	firstname string
	lastname  string
	email     string
}
