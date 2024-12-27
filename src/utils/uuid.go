package utils

import (
	"fmt"
	"math/rand/v2"
)

func Uuidv4() string {
	return generateUuidChunk(8) +
		"-" +
		generateUuidChunk(4) +
		"-4" +
		generateUuidChunk(3) +
		"-" +
		generateUuidVariant() +
		generateUuidChunk(3) +
		"-" +
		generateUuidChunk(12)
}

func GenerateRandomHexString(size int) string {
	return generateUuidChunk(size)
}

func generateUuidVariant() string {
	options := []string{"8", "9", "A", "B"}
	return options[rand.IntN(4)]
}

func generateUuidChunk(size int) string {
	result := ""
	for range size {
		digit := rand.IntN(16)
		result += fmt.Sprintf("%x", digit)
	}

	return result
}
