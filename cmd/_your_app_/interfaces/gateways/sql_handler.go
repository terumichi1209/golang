package gateways

// SQLHandler test
type SQLHandler interface {
	Create(object interface{})
	FindAll(object interface{})
	DeleteByID(object interface{}, id string)
}
