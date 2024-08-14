package punishment

import "time"

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

	// ID of the punishment
	Id int `gorm:"type:int;primary_key"`

	// Issuer of the punishment
	// Can be another Discord user identifier or "SENTINEL",
	// indicating any automatic punishment.
	Issuer string `gorm:"type:varchar(255)"`

	// Target is the user Discord id which receives the punishment.
	Target string `gorm:"type:varchar(255)"`

	// Expire date indicates the time of the expiration if present
	ExpireDate time.Time `gorm:"type:timestamp"`

	// Reason of the punishment
	Reason string `gorm:"type:varchar(255)"`

	// Type of punishment (Warn, Kick, Ban, Mute)
	Type Type `gorm:"type:int"`

	// Automatic if the punishment was issued by Sentinel itself.
	Automatic bool `gorm:"type:bool;default:false"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
