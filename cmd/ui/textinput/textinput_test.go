package textinput

import "testing"

func TestInputSanitization(t *testing.T) {
	passTestCases := []string{
		"chi",
		"new_project",
		"NEW_PROJECT",
		"new-project",
	}
	for _, testCase := range passTestCases {
		if err := sanitizeInput(testCase); err != nil {
			t.Errorf("expected no error, got: %v", err)
		}
	}
	failTestCases := []string{
		"new project",
		"NEW\\PROJECT",
		"new%project",
		" ",
		"  ",
		"#",
		"@",
	}
	for _, testCase := range failTestCases {
		if err := sanitizeInput(testCase); err == nil {
			t.Errorf("expected error for input %v, got nil", testCase)
		}
	}
}
