# Caesar's Cipher

### Usage

```go
package main

import (
    "fmt"
    cipher "github.com/1garo/caesar_cipher"
)

func main() {
    input := "testing"
    shift := 5
    encryptedText, err := cipher.Caesar(input, cipher.ENCRYPT, shift)

    if err != nil {
      panic("something went wrong!")
    }
    decrypted, err := cipher.Caesar(encryptedText, cipher.DECRYPT, shift)

    if err != nil {
      panic("something went wrong!")
    }

    if decrypted == input {
      fmt.Println("YAY!")
    } else {
      fmt.Println("NAY!")
    }
}
```
