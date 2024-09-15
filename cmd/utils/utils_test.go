package utils

import "testing"

func TestValidateModuleName(t *testing.T) {
	passTestCases := []string{
		"github.com/user/project",
		"github.com/user/projec1-hyphen",
		"github.com/user/projecT_under_Score",
		"github.com/user/project.hyphen3",
		"project",
		"ProJEct",
		"PRo_45-.4Jc",
		"PRo_/4J/c",
	}
	for _, testCase := range passTestCases {
		ok := ValidateModuleName(testCase)
		if !ok {
			t.Errorf("testing:%s expected:true got:%v", testCase, ok)
		}
	}

	failTestCases := []string{
		"",
		"/",
		".",
		"//",
		"/project",
		"ProJEct/",
		"PRo_$4Jc",
		"PRo_@J/c",
	}
	for _, testCase := range failTestCases {
		ok := ValidateModuleName(testCase)
		if ok {
			t.Errorf("testing:%s expected:false got:%v", testCase, ok)
		}
	}
}

func TestGeRootDir(t *testing.T) {
	testCases := map[string]string{
		"github.com/user/pro-ject": "pro-ject",
		"pro-ject":                 "pro-ject",
		"/":                        "",
		"":                         "",
		"//":                       "",
		"@":                        "@",
	}

	for intput, output := range testCases {
		rootDir := GetRootDir(intput)
		if rootDir != output {
			t.Errorf("testing:%s expected:%s got:%s", intput, output, rootDir)
		}
	}
}
