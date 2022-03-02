
```golang
import (
	"github.com/clivern/peacock/core/driver"
)

db := driver.NewDatabase(fmt.Sprintf("%s/.peacock/peacock.db", HOME))

err := db.Migrate()

err := db.Insert("fs", `{"id": 1, "name": "joe"}`)

value, err := db.FindByKey("fs")

fmt.Println(value)
```