package models

import "time"

type DepositeTransaction struct {
    TransactionID   int       
    AccountID       int      
    TransactionDate time.Time
    Amount          float64  
    Description     string   
}