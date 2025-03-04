package ch13

import (
	"fmt"
	"os"
)

func Writing() {

	err := os.WriteFile("out.txt",[]byte("Hello kumar reddy"),0644)
	if err!=nil{
		fmt.Println("error:",err)
	}

}

/**
### **Explanation of Writing Data to a File in Go**
The selected code:
```go
func main() {
    err := os.WriteFile("output.txt", []byte("Hello, Go!"), 0644)
    if err != nil {
        fmt.Println("Error writing file:", err)
    }
}
```
This code **writes data to a file** named `output.txt` using the `os.WriteFile` function. Here's how it works:

#### **🔹 Breakdown:**
1. **`os.WriteFile("output.txt", []byte("Hello, Go!"), 0644)`**  
   - `"output.txt"` → The name of the file to be created or overwritten.  
   - `[]byte("Hello, Go!")` → Converts the string `"Hello, Go!"` into a slice of bytes (`[]byte`), as files work with byte data.  
   - `0644` → **File permissions** in octal notation:  
     - `6` → Owner can **read & write**.  
     - `4` → Group can **read**.  
     - `4` → Others can **read**.

2. **`if err != nil { fmt.Println("Error writing file:", err) }`**  
   - Checks if an error occurs while writing the file and prints it.

#### **🔹 Why Use `os.WriteFile`?**
- It's **simple** and **convenient** for writing small files.
- No need to **open/close** the file manually (Go handles it internally).
- Efficient for **quick file writes**.

#### **🔹 Example Output (File Content)**
After running this program, `output.txt` will contain:
```
Hello, Go!
```
Let me know if you want an **advanced example** with **file appending or streaming writes**! 🚀
**/