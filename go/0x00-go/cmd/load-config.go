package main

import (
	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load()	
	if err != nil {
		panic("Error loading .env file")
	}
}