Thank you Mr. John Barton 

```bash
go get github.com/lucas-woo/godotenv
```

```go 
import "github.com/lucas-woo/godotenv"
```

To download load '.env' as default use 
```go
godotenv.LoadEnv()
```


Code snippet: 
```go
package main

import (
  "log"
  "os"
  "github.com/lucas-woo/godotenv"
)

func main() {
  err := godotenv.LoadEnv()
  if err != nil {
    log.Fatal(err.Error())
  }
  
  myPassword := os.Getenv("SUPER_SECRET_PASSWORD")
  ...
}

```