package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	empData := [][]string{
		{"UNIX", "SYMBOL", "OPEN", "HIGH", "LOW", "CLOSE"},
	}
	milliseconds := 1293843600000
	// Generate 10 Million records
	for i := 1; i < 100000; i++ { // 10000000
		milliseconds += 1000
		rand.Seed(time.Now().UnixNano())
		empData = append(empData, []string{strconv.Itoa(milliseconds), "BTCUSDT", fmt.Sprintf("%.8f", rand.Float64()*50000), fmt.Sprintf("%.8f", rand.Float64()*50000), fmt.Sprintf("%.8f", rand.Float64()*50000), fmt.Sprintf("%.8f", rand.Float64()*50000)})
	}
	// Generate 100 Million records data
	csvFile, err := os.Create("ohlc.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	csvWriter := csv.NewWriter(csvFile)
	// 10 Million rows writer
	for _, empRow := range empData {
		_ = csvWriter.Write(empRow)
	}
	csvWriter.Flush()
	csvFile.Close()
	log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")
	os.Exit(exitVal)
}

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func Hello() string {
	return "Hello, world."
}
