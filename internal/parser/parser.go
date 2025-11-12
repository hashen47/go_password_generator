package parser

import (
	"fmt"
	"github.com/hashen47/go_password_generator/internal/generator"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

const (
	flag_length_long_form       string = "len"
	flag_digit_short_form       string = "D"
	flag_digit_long_form        string = "digit"
	flag_uppercase_short_form   string = "U"
	flag_uppercase_long_form    string = "upper"
	flag_lowercase_short_form   string = "L"
	flag_lowercase_long_form    string = "lower"
	flag_specialcase_short_form string = "S"
	flag_specialcase_long_form  string = "special"
)

var (
	flag_length      string = ""
	flag_digit       string = ""
	flag_uppercase   string = ""
	flag_lowercase   string = ""
	flag_specialcase string = ""
)

const (
	flag_length_default_value      int = 8
	flag_digit_default_value       int = 1
	flag_uppercase_default_value   int = 1
	flag_lowercase_default_value   int = 1
	flag_specialcase_default_value int = 1

	flag_length_minimum_value      int = 4
	flag_digit_minimum_value       int = 0
	flag_uppercase_minimum_value   int = 0
	flag_lowercase_minimum_value   int = 0
	flag_specialcase_minimum_value int = 0
)

type validate_Argument_And_Get_Value_Result struct {
	is_default_value_set bool
	value                int
	err                  error
}

var Root_Cmd = &cobra.Command{
	Use:   "passwd_gen [len] [-D digit] [-L lower] [-U upper] [-S special]",
	Short: "generate custom password according to the given options",
	RunE:  command_handler,
}

func init() {
	Root_Cmd.Flags().StringVar(&flag_length, flag_length_long_form, flag_length, "password length")
	Root_Cmd.Flags().StringVarP(&flag_digit, flag_digit_long_form, flag_digit_short_form, flag_digit, "minimum digit character count")
	Root_Cmd.Flags().StringVarP(&flag_uppercase, flag_uppercase_long_form, flag_uppercase_short_form, flag_uppercase, "minimum uppercase character count")
	Root_Cmd.Flags().StringVarP(&flag_lowercase, flag_lowercase_long_form, flag_lowercase_short_form, flag_lowercase, "minimum lowercase character count")
	Root_Cmd.Flags().StringVarP(&flag_specialcase, flag_specialcase_long_form, flag_specialcase_short_form, flag_specialcase, "minimum specialcase character count")
}

func validate_argument_and_get_value(minimum_value, default_value int, value, short_form, long_form, error_msg string) validate_Argument_And_Get_Value_Result {
	result := validate_Argument_And_Get_Value_Result{}

	if strings.Compare(value, "") == 0 {
		result.is_default_value_set = true
		result.value = default_value
		return result
	}

	var argument_err error

	if strings.Compare(short_form, "") == 0 {
		argument_err = fmt.Errorf(`invalid argument "%s" for "--%s" flag: %s: minimum value should be %d`, value, long_form, error_msg, minimum_value)
	} else {
		argument_err = fmt.Errorf(`invalid argument "%s" for "-%s --%s" flag: %s: minimum value should be %d`, value, short_form, long_form, error_msg, minimum_value)
	}

	n, err := strconv.Atoi(value)
	if err != nil {
		result.err = argument_err
		return result
	}

	if n < minimum_value {
		result.err = argument_err
		return result
	}

	result.value = n
	return result
}

func command_handler(cmd *cobra.Command, args []string) error {
	result_length := validate_argument_and_get_value(flag_length_minimum_value, flag_length_default_value, flag_length, "", flag_length_long_form, "should be an integer")
	if result_length.err != nil {
		return result_length.err
	}

	result_digit_count := validate_argument_and_get_value(flag_digit_minimum_value, flag_digit_default_value, flag_digit, flag_digit_short_form, flag_digit_long_form, "should be an integer")
	if result_digit_count.err != nil {
		return result_digit_count.err
	}

	result_lowercase_count := validate_argument_and_get_value(flag_lowercase_minimum_value, flag_lowercase_default_value, flag_lowercase, flag_lowercase_short_form, flag_lowercase_long_form, "should be an integer")
	if result_lowercase_count.err != nil {
		return result_lowercase_count.err
	}

	result_uppercase_count := validate_argument_and_get_value(flag_uppercase_minimum_value, flag_uppercase_default_value, flag_uppercase, flag_uppercase_short_form, flag_uppercase_long_form, "should be an integer")
	if result_uppercase_count.err != nil {
		return result_uppercase_count.err
	}

	result_specialcase_count := validate_argument_and_get_value(flag_specialcase_minimum_value, flag_specialcase_default_value, flag_specialcase, flag_specialcase_short_form, flag_specialcase_long_form, "should be an integer")
	if result_specialcase_count.err != nil {
		return result_specialcase_count.err
	}

	constraints_length := result_digit_count.value + result_lowercase_count.value + result_uppercase_count.value + result_specialcase_count.value
	if constraints_length == 0 && !result_digit_count.is_default_value_set && !result_lowercase_count.is_default_value_set && !result_uppercase_count.is_default_value_set && !result_specialcase_count.is_default_value_set {
		return fmt.Errorf("constraints sum cannot be zero, at least one should non zero")
	}

	if constraints_length > result_length.value {
		if !result_length.is_default_value_set {
			return fmt.Errorf("constraints sum cannot be exceed password length (length :%d, digit: %d, upper: %d, lower: %d, special: %d)", result_length.value, result_digit_count.value, result_uppercase_count.value, result_lowercase_count.value, result_specialcase_count.value)
		}
	}

	fmt.Printf("%s\n", generator.Generate_password(result_length.value, result_digit_count.value, result_lowercase_count.value, result_uppercase_count.value, result_specialcase_count.value))

	return nil
}
