package mydict

import "errors"

//Dictionary type
type Dictionary map[string]string

//return에 적는 것보다 이렇게 변수로 선언하는 것이 좋음
var (
	errNotFound   = errors.New("Not Found")
	errWordExists = errors.New("That word already exists")
	errCantUpdate = errors.New("Cant update non-existing word")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word] //map은 두 가지값을 반환하는데 하나는 키에 해당하는 value 그리고 값이 존재하는지 boolean값
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word string, def string) error {
	_, err := d.Search(word)

	// 방법 1
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil

	// // 방법 2
	// if err == errNtFound { //dictionary안에 추가하려는 단어가 없을 때
	// 	d[word] = def
	// } else if err == nil {
	// 	return errWordExists
	// }
}

// Update a word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {
	// delete함수를 사용할 수 있는데 아무 값도 return하지 않고 첫번째 인자에는 삭제하려는 맵이름, 두번째 인자에는 key값
	delete(d, word)

}
