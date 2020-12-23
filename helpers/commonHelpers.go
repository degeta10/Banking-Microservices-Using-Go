package helpers

import (
	"time"
)

// Location ..
var Location, _ = time.LoadLocation("Asia/Kolkata")

// CurrentDateTime ..
var CurrentDateTime = time.Now().In(Location).Format(time.RFC3339)
