package main

import (
	"context"
	"fmt"
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var ctx context.Context

func main() {
	fmt.Println("Welcome to the Sandwich Shop.")
}
