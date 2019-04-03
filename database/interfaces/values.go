package interfaces

//TODO: Some kind of queue for the music module

type StringValue interface {
	Set(value string) error
	Get() string
}

type MapValue interface {
	Get(field string) StringValue
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
	Map(key string) MapValue
	Set(key string) SetValue
	Decorate(key string) string
	GetContainer(key string) Container
}

type Config interface {
	OptionalValue(name string) (value StringValue, exists bool)
	RequiredValue(name string, defaultValue string) StringValue
}