package schema

import "time"

type Readings struct {
    ID      int       `db:"id"`
    Time    time.Time `db:"time"`
    Reading float32   `db:"reading"`
}
