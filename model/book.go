package model
type Book struct {
  Title, Author string
  Pages int
}

func (book Book) CategoryByLength() string {
  if book.Pages >= 300 {
    return "NOVEL"
  }else{
    return "ARTICLE"
  }
}
