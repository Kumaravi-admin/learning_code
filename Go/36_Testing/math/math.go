/*
                  Testing


• The standard lib testing package and go test command provide all you need.
• Test file name should end with _test.go and are not compiled.
• Example: If you are testing a sum function in sum.go file, test file will be sum_test.go.
• Test functions are name TestXxx, example: TestSum.



                  *testing.T


• The test function accepts one argument which is a pointer to testing.T.
• T is a structure that helps orchestrate your tests.
• Methods like Error, Errorf, Fatal, Fatalf, Log, etc can be used to inform of testing outcomes and log results.



                  Error & Fatal


• Method Error and Errorf fail the tests but test execution continues.
• Method Fatal and Fatalf fail the tests and stop the execution of test.


                  Subtests & Cleanup


• *testing.T has method Run that is used to run subtests within the test function.
• *testing.T has method Cleanup that can be used to perform cleanups.




*/

// 1. Inisilize packege using go mode command
// Example : go.mod: module github.com/Kumaravi-admin/lesson_testing

package math

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}