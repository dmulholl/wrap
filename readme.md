# Wrap

A simple utility for wrapping text to within a specified maximum line-length, using a customizable set of breakpoints where possible.

```go
import "github.com/dmulholl/wrap"

output := wrap.Wrap(text, limit)
```

Line-length in the output is guaranteed to be within the specified `limit`.


## Caveats

Line-length is specified as a simple rune-count, i.e. as a count of Unicode code-points.

This library doesn't attempt to account for display-width issues with double-width code-points or multi-code-point grapheme clusters, e.g. emojis.
