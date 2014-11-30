package caseformat

import (
	"testing"
)

func TestLowerCamel(t *testing.T) {

	variations := []string{"this_is_a_test", "ThisIsATest", "THIS_IS_A_TEST", "thisIsATest"}
	expected := "thisIsATest"

	for _, in := range variations {
		if res := ToLowerCamel(in); res != expected {
			t.Errorf("toLowerCamel(%v) = %v, want %v", in, res, expected)
		}
	}

}

func TestUpperCamel(t *testing.T) {

	variations := []string{"this_is_a_test", "ThisIsATest", "THIS_IS_A_TEST", "thisIsATest"}
	expected := "ThisIsATest"

	for _, in := range variations {
		if res := ToUpperCamel(in); res != expected {
			t.Errorf("ToUpperCamel(%v) = %v, want %v", in, res, expected)
		}
	}

}

func TestUpperUnderscore(t *testing.T) {

	variations := []string{"this_is_a_test", "ThisIsATest", "THIS_IS_A_TEST", "thisIsATest"}
	expected := "THIS_IS_A_TEST"

	for _, in := range variations {
		if res := ToUpperUnderscore(in); res != expected {
			t.Errorf("ToUpperUnderscore(%v) = %v, want %v", in, res, expected)
		}
	}

}

func TestLowerUnderscore(t *testing.T) {

	variations := []string{"this_is_a_test", "ThisIsATest", "THIS_IS_A_TEST", "thisIsATest"}
	expected := "this_is_a_test"

	for _, in := range variations {
		if res := ToLowerUnderscore(in); res != expected {
			t.Errorf("ToLowerUnderscore(%v) = %v, want %v", in, res, expected)
		}
	}

}
