package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type TransactionType int

const (
	Credit TransactionType = iota
	Debit
)

func (d TransactionType) String() string {
	return [...]string{"Credit", "Debit"}[d]
}

type User struct {
	gorm.Model
	Name         string `gorm:"unique"`
	Balance      uint
	Transactions []Transaction
}

type Transaction struct {
	gorm.Model
	Amount      uint
	Type        string
	Description string
	UserID      uint
}

var db *gorm.DB

func init() {
	logFile, err := os.Create("hello.log")
	if err != nil {
		log.Fatal(err)
	}
	logger := logger.New(
		log.New(logFile, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  false,
		},
	)
	db, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{Logger: logger})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&User{}, &Transaction{})
}

func CreateUser(name string, balance uint) {
	user := User{Name: name, Balance: balance}
	if result := db.Create(&user); result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Printf("User %v created\n", name)
}

func GetUser(name string) (user User) {
	if result := db.Preload("Transactions").First(&user, "Name = ?", name); result.Error != nil {
		log.Fatal(result.Error)
	}
	return
}

func AddTransaction(name string, transactionType TransactionType, amount uint, description string) {
	var user User
	if result := db.First(&user, "Name = ?", name); result.Error != nil {
		log.Fatal(result.Error)
	}
	switch transactionType {
	case Debit:
		if user.Balance < amount {
			fmt.Println("Balance is too small for this transaction")
			return
		}
		user.Balance -= amount
	case Credit:
		user.Balance += amount
	}
	user.Transactions = append(user.Transactions, Transaction{Amount: amount, Type: transactionType.String(), Description: description})
	if result := db.Save(&user); result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Printf("Added %v transaction to %v. Balance: %v\n", transactionType, name, user.Balance)
}
