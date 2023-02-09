package main

import (
    
    "fmt"
    "os"
)

type Book struct {
    ID    int
    Title string
    Author string
}

var books []Book

func main() {
    for {
        fmt.Println("1. Agregar libro")
        fmt.Println("2. Editar libro")
        fmt.Println("3. Eliminar libro")
        fmt.Println("4. Ver libros")
        fmt.Println("5. Salir")

        var option int
        fmt.Scanln(&option)

        switch option {
        case 1:
            addBook()
        case 2:
            editBook()
        case 3:
            deleteBook()
        case 4:
            viewBooks()
        case 5:
            fmt.Println("Gracias por usar la aplicación")
            os.Exit(0)
        default:
            fmt.Println("Opción inválida, intente de nuevo")
        }
    }
}

func addBook() {
    var id int
    var title, author string

    fmt.Print("ID: ")
    fmt.Scanln(&id)
    fmt.Print("Título: ")
    fmt.Scanln(&title)
    fmt.Print("Autor: ")
    fmt.Scanln(&author)

    books = append(books, Book{ID: id, Title: title, Author: author})
    fmt.Println("Libro agregado exitosamente")
}

func editBook() {
    var id int
    fmt.Print("ID: ")
    fmt.Scanln(&id)

    for i, book := range books {
        if book.ID == id {
            fmt.Print("Nuevo título: ")
            fmt.Scanln(&books[i].Title)
            fmt.Print("Nuevo autor: ")
            fmt.Scanln(&books[i].Author)
            fmt.Println("Libro editado exitosamente")
            return
        }
    }

    fmt.Println("Libro no encontrado")
}

func deleteBook() {
    var id int
    fmt.Print("ID: ")
    fmt.Scanln(&id)

    for i, book := range books {
        if book.ID == id {
            books = append(books[:i], books[i+1:]...)
            fmt.Println("Libro eliminado exitosamente")
            return
        }
    }

    fmt.Println("Libro no encontrado")
}

func viewBooks() {
    fmt.Println("ID\tTítulo\tAutor")
    for _, book := range books {
        fmt.Printf("%d\t%s\t%s\n", book.ID, book.Title, book.Author)
    }
}