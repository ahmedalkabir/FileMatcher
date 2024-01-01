# FileMatcher
A Simple File Detector to infer binary file types based on some criteria (magic numbers, etc...)

# Example 
```go
package main

import (
	"fmt"
	"os"
	" github.com/ahmedalkabir/filematcher"
)

func main() {
	file, err := os.Open("./file.docx")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	buff := make([]byte, 550)
	_, err = file.Read(buff)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	filetype, _ := filematcher.Match(buff)
	fmt.Println(filetype.Type)

}

```
