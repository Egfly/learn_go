package models

import "errors"

type Book struct {
	Name       string
	Total      int
	Author     string
	CreateTime string
}

var (
	ErrBorrowNumIllegal = errors.New("borrow num illegal")
	ErrStockNotEnough   = errors.New("stock is not enough")
)

/**
创建书籍
*/
func CreateBook(name string, total int, author string, createTime string) (b *Book) {
	b = &Book{
		Name:       name,
		Total:      total,
		Author:     author,
		CreateTime: createTime,
	}
	return
}

/**
是否能借
*/
func (b *Book) CanBorrow(num int) (bool, error) {
	if num <= 0 {
		return false, ErrBorrowNumIllegal
	}
	if b.Total-num < 0 {
		return false, ErrStockNotEnough
	}
	return true, nil
}

/**
借书
*/
func (b *Book) Borrow(num int) (bool, error) {
	_, err := b.CanBorrow(num)
	if err != nil {
		return true, err
	}
	b.Total -= num
	return true, nil
}

/**
还书
*/
func (b *Book) Back(num int) {
	b.Total += num
}
