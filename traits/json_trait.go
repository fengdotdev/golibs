package traits

type JSONSerializable interface {
	ToJSON() (string, error)
}

type JSONDeserializable interface {
	FromJSON(string) error
}

type JSONTrait interface {
	JSONDeserializable
	JSONSerializable
}
