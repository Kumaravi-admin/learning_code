// Same folder = same package. Both files (math.go and int.go) belong to package math.
package math

// Files in the same package share everything, even unexported names.
// So subtract() from math.go can be used here WITHOUT any import.
func something() {
	subtract(4, 2)
}