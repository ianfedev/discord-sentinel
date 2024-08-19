package punishment

import (
	"discord-sentinel/core/database"
	"time"
)

type Type int

const (
	Warn Type = 1
	Kick Type = 2
	Ban  Type = 3
	Mute Type = 4
)

// Punishment struct defines the record of a user punishment
// when is applied automatically or manually.
type Punishment struct {
	database.Model

	// Issuer of the punishment
	// Can be another Discord user identifier or "SENTINEL",
	// indicating any automatic punishment.
	Issuer string `json:"issuer" gorm:"type:varchar(255)"`

	// Target is the user Discord id which receives the punishment.
	Target string `json:"target" gorm:"type:varchar(255)"`

	// Expire date indicates the time of the expiration if present
	ExpireDate time.Time `json:"expire_date" gorm:"type:timestamp"`

	// Reason of the punishment
	Reason string `json:"reason" gorm:"type:varchar(255)"`

	// Type of punishment (Warn, Kick, Ban, Mute)
	Type Type `json:"type" gorm:"type:int"`

	// Automatic if the punishment was issued by Sentinel itself.
	Automatic bool `json:"automatic" gorm:"type:bool;default:false"`
}
