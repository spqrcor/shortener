package storage

import (
	"context"
	"fmt"
	"shortener/internal/config"
)

func ExampleMemoryStorage_Add() {
	memoryStorage := CreateMemoryStorage(config.NewConfig())
	url := "https://ya.ru"

	_, err := memoryStorage.Add(context.Background(), url)
	if err == nil {
		fmt.Println("Success add")
	} else {
		fmt.Println("Error add")
	}

	// Output:
	// Success add
}
