package main

import "fmt"

type validationError struct {
	Message string
}

func (e *validationError) Error() string {
	return e.Message
}

type notFoundError struct {
	Message string
}

func (e *notFoundError) Error() string {
	return e.Message
}

func SaveData (id string, data any) error {
	if id == "" {
		return &validationError{"Id tidak boleh kosong, validation error"}
	} 

	if id != "Eko" {
		return &notFoundError{"Data tidak ditemukan, not found error"}
	}

	// ok

	return nil
}

func main47()  {
	err := SaveData("Eko", nil)
	if err != nil {
		// terjadi error
		// Menggunakan if
		// if validationErr, ok := err.(*validationError); ok {
		// 	fmt.Println("Validation Error", validationErr.Message)
		// } else if notFoundErr, ok := err.(*notFoundError); ok {
		// 	fmt.Println("Not Found Error", notFoundErr.Message)
		// } else {
		// 	fmt.Println("Unknow error: ", err.Error())
		// }

		// Menggunakan switch case
		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("Validation Error", finalError.Message)
		case *notFoundError:
			fmt.Println("Not Found Error", finalError.Message)
		default:
			fmt.Println("Unknow error: ", err.Error())
		}
	} else {
		// berhasil
		fmt.Println("Sukses")
	}
}