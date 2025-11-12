package generator

import (
	"math/rand"
)

var (
	shuffle_iter_count int    = 5
	digits             string = "0123456789"
	lowercase          string = "abcdefghijklmnopqrstuvwxyz"
	uppercase          string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialcase        string = `!@#$%^&*()_-+=[]{}|;:'",<>/?`
)

func fill_password_with_given_character_set(password *[]byte, start_index *int, character_set *string, count int) {
	for i := *start_index; i < *start_index+count; i++ {
		(*password)[i] = (*character_set)[rand.Intn(len(*character_set))]
	}
	*start_index += count
}

func Generate_password(length, digit_count, lowercase_count, uppercase_count, specialcase_count int) string {
	password := make([]byte, length)
	all_chars := ""

	start_index := 0
	fill_password_with_given_character_set(&password, &start_index, &digits, digit_count)
	fill_password_with_given_character_set(&password, &start_index, &lowercase, lowercase_count)
	fill_password_with_given_character_set(&password, &start_index, &uppercase, uppercase_count)
	fill_password_with_given_character_set(&password, &start_index, &specialcase, specialcase_count)

	if digit_count > 0 {
		all_chars += digits
	}

	if lowercase_count > 0 {
		all_chars += lowercase
	}

	if uppercase_count > 0 {
		all_chars += uppercase
	}

	if specialcase_count > 0 {
		all_chars += specialcase
	}

	for i := start_index; i < length; i++ {
		password[i] = all_chars[rand.Intn(len(all_chars))]
	}

	// shuffle the characters
	for i := 0; i < shuffle_iter_count; i++ {
		for j := 0; j < length; j++ {
			switch_index := rand.Intn(length)
			temp := password[switch_index]
			password[switch_index] = password[j]
			password[j] = temp
		}
	}

	return string(password)
}
