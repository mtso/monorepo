# GET

`GET` retrieves the value for a given key.

<div class="snippet-group">

```tcp
GET foo
bar
```

```js
const tackdb = require('tackdb');
tackdb.connect(
  process.env.TACKDB_URL,
  function(err) { throw new Error(err); }
);

const foo = tackdb.get('foo');
console.log(foo);

const foo = tackdb.json.get('foo');
console.log(foo);
```

```go
package main

import (
    "fmt"
    "github.com/tackdb/tackdb-go"
    "os"
)

func main() {
    store, err := tackdb.Connect(os.Getenv("TACKDB_URL"))
    if err != nil {
        panic(err)
    }

    foo := store.Get("foo")
    fmt.Println("foo:", foo)
}
```

</div>
