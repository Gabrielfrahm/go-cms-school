package adapters

type Hash interface {
	Hashed(value string) (string, error)
	CompareHashed(value, valueHashed string) (bool, error)
}
