package models

var counterID uint = 0

func GetID() uint {
	counterID++
	return counterID
}
