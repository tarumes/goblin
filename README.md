# Documentation

## Usage Examples

```Go
func main {

	var testData struct {
		Hello []string
	} = struct{ Hello []string }{
		[]string{"World", "world"},
	}

	err := GOBWrite("test.go", testData)
	if err != nil {
		log.Error(err)
	}

	err = GOBRead("test.go", &testData)
	if err != nil {
		log.Error(err)
	}
}
```
