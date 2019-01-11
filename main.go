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

func obtainVector(atts map[string]float64) []float64 {

	values := make([]float64, 0, len(atts))

	for _, v := range atts {
		values = append(values, v)
	}

	return values
}

func cosine(a []float64, b []float64) (cosine float64, err error) {
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

func compareCharacter(database []Character, myCharacter *Character) *Character {

	var mostSimilar Character
	var maxSimilarity float64

	for _, current := range database {

		cosine, err := cosine(
			obtainVector(myCharacter.Attributes),
			obtainVector(current.Attributes))

		if err == nil && cosine > maxSimilarity {
			mostSimilar = current
			maxSimilarity = cosine
		}
	}

	return &mostSimilar
}

func menuCompareCharacter(database []Character) *Character {

	myCharacter := Character{}

	myCharacter.Attributes = make(map[string]float64)
	myCharacter.Attributes["Amigou"] = 2
	myCharacter.Attributes["Joder"] = 2

	character := compareCharacter(database, &myCharacter)

	fmt.Printf("%+v\n", character)

	return character
}

func storeDatabase(database []Character) error {

	dbBytes, _ := json.MarshalIndent(database, "", "  ")

	return ioutil.WriteFile("database.json", dbBytes, 0644)
}

func menu() uint8 {

	fmt.Println("--- MENU ---")
	fmt.Println("1.\tCreate character")
	fmt.Println("2.\tCompare character")
	fmt.Println("3.\tModify character")
	fmt.Println("4.\tDelete character")
	fmt.Println("0.\tExit")

	var option uint8
	_, err := fmt.Scan(&option)

	if err != nil {
		return 100
	}

	return option
}

func menuLoop(database []Character) {

Menu:
	for {
		option := menu()

		switch option {
		case 1:
			fmt.Println("CREATE")
		case 2:
			menuCompareCharacter(database)
		case 3:
			fmt.Println("MODIFY")
		case 4:
			fmt.Println("DELETE")
		case 0:
			break Menu
		default:
			fmt.Println("WRONG OPTION")
		}
	}
}

func main() {

	database := loadDatabase()

	menuLoop(database)

	storeDatabase(database)
}
