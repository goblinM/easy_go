package main

import "fmt"

/**
结构体
结构体转map
https://zhuanlan.zhihu.com/p/163000449
*/

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func printBook(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
func main() {
	var book_obj Books
	book_obj.author = "Jane"
	book_obj.title = "old man and sea"
	book_obj.subject = "a classics story"
	book_obj.book_id = 1
	printBook(&book_obj)
	book2 := Books{title: "abc", author: "kk", subject: "sdasca", book_id: 12314}
	printBook(&book2)
}
