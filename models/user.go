package models

import (
	"fmt"
	"strconv"
	"time"
)

type UserRole int

const (
	Doctor UserRole = iota
	Receptionist
)

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`               // Explicit autoincrement
	Name      string    `gorm:"type:varchar(100);not null"`             // Define length for portability
	Email     string    `gorm:"type:varchar(100);uniqueIndex;not null"` // uniqueIndex preferred over unique
	Password  string    `gorm:"type:varchar(255);not null"`             // Hashes can be long; define length
	Role      UserRole  `gorm:"type:int;not null"`                      // Case-sensitive ENUM
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// String method to convert UserRole to string
func (r UserRole) String() string {
	switch r {
	case Doctor:
		return "doctor"
	case Receptionist:
		return "receptionist"
	default:
		return "unknown"
	}
}

// ParseRole function to convert string to UserRole
func ParseRole(role string) (UserRole, error) {
	switch role {
	case "doctor":
		return Doctor, nil
	case "receptionist":
		return Receptionist, nil
	default:
		return -1, fmt.Errorf("invalid role: %s", role)
	}
}

// Scan method for converting the database value to UserRole

func (r *UserRole) Scan(value interface{}) error {
	switch v := value.(type) {
	case int64:
		*r = UserRole(int(v)) // safe conversion
	case int:
		*r = UserRole(v)
	default:
		return fmt.Errorf("cannot scan %T into UserRole", value)
	}
	return nil
}

// Value method for converting UserRole to database value (int)
func (r UserRole) Value() (interface{}, error) {
	return int(r), nil
}

func (u *User) IDString() string {
	return strconv.Itoa(int(u.ID))
}
