package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
)

// Character represents a character
type Character struct {
	Name       string
	Attributes map[string]float64
}

// ObtainVector obtains the vector from attribute map
func ObtainVector(atts map[string]float64) []float64 {

	values := make([]float64, 0, len(atts))

	for _, v := range atts {
		values = append(values, v)
	}

	return values
}

// Cosine calculates cosine
func Cosine(a []float64, b []float64) (cosine float64, err error) {
	count := 0
	lengthA := len(a)
	lengthB := len(b)
	if lengthA > lengthB {
		count = lengthA
	} else {
		count = lengthB
	}
	sumA := 0.0
	s1 := 0.0
	s2 := 0.0
	for k := 0; k < count; k++ {
		if k >= lengthA {
			s2 += math.Pow(b[k], 2)
			continue
		}
		if k >= lengthB {
			s1 += math.Pow(a[k], 2)
			continue
		}
		sumA += a[k] * b[k]
		s1 += math.Pow(a[k], 2)
		s2 += math.Pow(b[k], 2)
	}
	if s1 == 0 || s2 == 0 {
		return 0.0, errors.New("Vectors should not be null (all zeros)")
	}
	return sumA / (math.Sqrt(s1) * math.Sqrt(s2)), nil
}

func loadDatabase() []Character {

	var database []Character

	dbBytes, _ := ioutil.ReadFile("database.json")
	json.Unmarshal(dbBytes, &database)

	fmt.Printf("%+v\n", database)

	return database
}

func storeDatabase(database []Character) error {

	s3, _ := json.MarshalIndent(database, "", "  ")

	return ioutil.WriteFile("database.json", s3, 0644)
}

func menu() uint8 {

	fmt.Println("--- MENU ---")
	fmt.Println("1.\tCreate character")
	fmt.Println("2.\tCompare character")
	fmt.Println("0.\tExit")

	var option uint8
	_, err := fmt.Scan(&option)

	if err != nil {
		return 100
	}

	return option
}

func menuLoop() {

Menu:
	for {
		option := menu()

		switch option {
		case 1:
			fmt.Println("CREATE")
		case 2:
			fmt.Println("COMPARE")
		case 0:
			fmt.Println("WHATEVER")
			break Menu
		default:
			fmt.Println("WRONG OPTION")
		}
	}
}

func main() {

	database := loadDatabase()

	menuLoop()

	storeDatabase(database)
}
