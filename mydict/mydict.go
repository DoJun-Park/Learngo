package mydict

import "errors"

//Dictionary type
type Dictionary map[string]string

//return에 적는 것보다 이렇게 변수로 선언하는 것이 좋음
var errNotFound = errors.New("not found")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word] //map은 두 가지값을 반환하는데 하나는 키에 해당하는 value 그리고 값이 존재하는지 boolean값
	if exists {
		return value, nil
	}
	return "", errNotFound
}
