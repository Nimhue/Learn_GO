package main

import "fmt"

func main() {
  var publisher, writer, artist, title, genre string
  var year, pageNumber int
  var grade float32 

  title ="Mr. GoToSleep"
  writer = "Tracey Hatchet"
  artist = "Jewel Tampson"
  publisher = "DizzyBooks Publishing Inc."
  year = 1997
  pageNumber = 14
  grade = 6.5
  genre = "Horror"

  fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "in", year, "with", pageNumber, "pages and with a grade of", grade, ".", "This comic is in the", genre, "genre.")

  fmt.Println("***************************")

  title, writer, artist, publisher, genre = "Epic Vol. 1", "Ryan N. Shawn", "Phoebe Paperclips", "DizzyBooks Publishing Inc.", "Adventure"

  year, pageNumber, grade = 2013, 160, 9.0

  fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "in", year, "with", pageNumber, "pages and with a grade of", grade, ".", "This comic is in the", genre, "genre.")

  fmt.Println("***************************")

    title, writer, artist, publisher, genre = "Lightning Strike", "Peter Parker", "Pitty Party", "DizzyBooks Publishing Inc.", "Action"

  year, pageNumber, grade = 2019, 50, 7.6

  fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "in", year, "with", pageNumber, "pages and with a grade of", grade, ".", "This comic is in the", genre, "genre.")
}
