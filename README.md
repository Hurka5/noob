# NoOb
NoOb (No Obscenity) is a blazingly fast profanity filtering library which supports filtering in multiple languages.

# Usage
```go
import (
    "fmt"
	"github.com/Hurka5/noob"
	_"github.com/Hurka5/noob/langs/en" // English language support
	_"github.com/Hurka5/noob/langs/hu" // Hungarian language support
)

func main() {
    
    // IsProfane
    word := "fuck"
    if noob.IsProfane(word) {
        fmt.Printf("%v is a bad word :(", word)
    }

}
```

