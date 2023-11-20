# Cesar Cipher

### Usage

```go
package main

import (
    "fmt"
    cipher "github.com/1garo/cesar_cipher"
)

func main() {
    input := "testing"
    shift := 5
    encryptedText, err := cipher.Cesar(input, cipher.ENCRYPT, shift)

    if err != nil {
      panic("something went wrong!")
    }
    decrypted, err := cipher.Cesar(encryptedText, cipher.DECRYPT, shift)

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
