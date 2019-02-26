package database

type StringValue interface {
	Set(value string) error
	Get() string
}

