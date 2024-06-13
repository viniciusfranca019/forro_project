package Seeders

import (
	"fmt"
	Location "forro_project/api/Location/Domain/Model"
	"forro_project/packages/Database/Config"
	"forro_project/packages/Util"
	"gorm.io/gorm"
)

func LocationSeeds() {
	migrateBrazil()
}

func migrateBrazil() {
	sourceData := "packages/Database/Seeds/Storage/brazil_cities.csv"
	cities := Util.HandleOpenCsvFile(&sourceData)

	db := Config.StartDbConnection()

	brazil := createCountry(db, "Brazil", "BR")
	stateMap := make(map[string]*Location.State)

	for index, cityData := range cities {
		if index == 0 {
			continue
		}

		stateCode := cityData[4]
		stateName := cityData[5]
		cityName := cityData[2]

		state, ok := stateMap[stateCode]
		if !ok {
			state = createState(db, *brazil, stateName, stateCode)
			stateMap[stateCode] = state
		}

		createCity(db, state, cityName)
	}

	fmt.Println("Data import completed.")
}

func createCountry(db *gorm.DB, countryName string, abbreviation string) *Location.Country {
	var country Location.Country
	db.FirstOrCreate(&country, Location.Country{Name: countryName, Abbreviation: abbreviation})
	if country.Name == "" {
		country = Location.Country{Name: countryName, Abbreviation: abbreviation}
		db.Create(&country)
	}
	return &country
}

func createState(db *gorm.DB, country Location.Country, name, abbreviation string) *Location.State {
	var state Location.State
	db.FirstOrCreate(&state, Location.State{Name: name, Abbreviation: abbreviation, CountryID: country.ID})
	if state.Abbreviation == "" {
		state = Location.State{Name: name, Abbreviation: abbreviation, CountryID: country.ID}
		db.Create(&state)
	}
	return &state
}

func createCity(db *gorm.DB, state *Location.State, name string) *Location.City {
	var city Location.City
	db.FirstOrCreate(&city, Location.City{Name: name, StateID: state.ID})
	if city.Name == "" {
		city := Location.City{Name: name, StateID: state.ID}
		db.Create(&city)
	}
	return &city
}
