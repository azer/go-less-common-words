## go-less-common-words

Filter [most common english words](http://github.com/azer/go-most-common-words) from given string or list of strings.

```go
import (
  "github.com/azer/go-less-common-words"
)

lessCommonWords.Keywords("I'm from Iceland and I make goat cheese. How about you? Do you work?")
// => ["iceland", "goat", "cheese"]
```

See `less_common_words_test.go` for more info.
