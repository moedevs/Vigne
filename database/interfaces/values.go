package interfaces

//TODO: Some kind of queue for the music module

type StringValue interface {
	Set(value string) error
	Get() string
}

type IntegerValue interface {
	Set(value int) error
	Get() (int, error)
	Add(amount int) error
}

type MapValue interface {
	Get(field string) StringValue
	Contains(field string) bool
	//TODO: GetAll
}

type IntegerMapValue interface {
	Get(field string) IntegerValue
	Contains(field string) bool
	//TODO: GetAll
}

type SetValue interface {
	Add(member string) error
	Contains(member string) bool
	Remove(member string) error
}


type Container interface{
	Value(key string) StringValue
	Integer(key string) IntegerValue
	Map(key string) MapValue
	IntegerMap(key string) IntegerMapValue
	Set(key string) SetValue
	Decorate(key string) string
	GetContainer(key string) Container
}

type Config interface {
	OptionalValue(name string) (value StringValue, exists bool)
	RequiredValue(name string, defaultValue string) StringValue
}