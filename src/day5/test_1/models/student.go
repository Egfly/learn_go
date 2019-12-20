package models

import "errors"

type Student struct {
	id    int
	Name  string
	Age   int
	Grade int
	books []*BorrowItem
}

var (
	ErrBackBookNumTooLarge = errors.New("you hadn't borrow enough book")
	ErrNotBorrowBook       = errors.New("you hadn't borrow this book")
)

/**
借书记录
*/
type BorrowItem struct {
	book *Book
	num  int
}

/**
创建学员
*/
func createStudent(id int, name string, age int, grade int) *Student {
	stu := &Student{
		id:    id,
		Name:  name,
		Age:   age,
		Grade: grade,
		books: nil,
	}
	return stu
}

/**
借书
*/
func (s *Student) BorrowBook(b *BorrowItem) {
	s.books = append(s.books, b)
}

/**
还书
*/
func (s *Student) BackBook(b *BorrowItem) (err error) {
	isset := false
	for i := 0; i < len(s.books); i++ {
		if s.books[i].book.Name == b.book.Name {
			if b.num > s.books[i].num {
				return ErrBackBookNumTooLarge
			}
			if b.num == s.books[i].num {
				left := s.books[:i-1]
				right := s.books[i+1:]
				s.books = append(left, right...)
				return
			}
			s.books[i].num -= b.num
			isset = true
		}
	}

	if isset == false {
		return ErrNotBorrowBook
	}
	return
}
