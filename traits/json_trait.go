package traits

type JSONSerializable interface {
	ToJson() string
}

type JSONDeserializable interface {
	FromJson(string)
}

type JSONTrait interface {
	JSONDeserializable
	JSONSerializable
}
