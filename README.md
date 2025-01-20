# NoOb
NoOb (No Obscenity) is a blazingly fast profanity filtering library which supports filtering in multiple languages.

I wanted to make a profanity filter that is not as "sanitized" as other filters but 

# Usage
```go
import (
    "fmt"
	"github.com/Hurka5/noob"
    _ "github.com/Hurka5/noob/hard/en" // Hard english filter 
    _ "github.com/Hurka5/noob/soft/hu" // Soft hungarian filter 

)

func main() {
    
    // IsProfane
    word := "frick"
    if noob.IsProfane(word) {
        fmt.Printf("%v is a bad word :(", word)
    }

}
```

