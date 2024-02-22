package models

type Config struct{
	Host string
	Port string
	User string
	Password string
	DBname string
	SSLmode string
}

type User struct {
    ID uint64 `gorm:"primaryKey"`
}

// type Blockchain struct {
//     Index      uint   `gorm:"primaryKey"`
//     PrevHash   string `gorm:"not null"`
//     Timestamp  string `gorm:"not null"`
//     Data       string `gorm:"not null"`
//     Hash       string `gorm:"not null;unique"`
//     OwnerID    uint   `gorm:"not null"`
//     Owner      User   `gorm:"foreignKey:OwnerID"`
// }

type Poll struct {
    Block string `gorm:"primaryKey"`
    Title string `gorm:"not null"`
}

type Option struct {
    Block     string `gorm:"primaryKey"`
    Text      string `gorm:"not null"`
    PollBlock string `gorm:"not null;foreignKey:PollBlock"`
}

type Vote struct {
    Block       string `gorm:"primaryKey"`
    PollBlock   string `gorm:"not null"`
    OptionBlock string `gorm:"not null"`
}