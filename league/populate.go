package league

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

func PopulateDatabase(db *sql.DB, src string) error {
	log.Println("Creating tables if necessary")
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS teams (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			city TEXT,
			name TEXT,
			abbr TEXT
		);
	`
	_, err := db.Exec(createTableQuery)
	if err != nil {
		return err
	}
	createTableQuery = `
		CREATE TABLE IF NOT EXISTS players (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			firstName TEXT,
			lastName TEXT,
			pos TEXT,
			number INTEGER,
			team INTEGER,
			SPD INTEGER,
			STR INTEGER,
			ACC INTEGER,
			THP INTEGER,
			THA INTEGER,
			CAR INTEGER,
			CAT INTEGER,
			BLK INTEGER,
			TAC INTEGER,
			COV INTEGER,
			BLS INTEGER,
			PSR INTEGER,
			FOREIGN KEY (team) REFERENCES teams(id)
		);
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		return err
	}

	log.Printf("Reading source file %s\n", src)
	srcFile, err := os.OpenFile(src, os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer srcFile.Close()
	if err != nil {
		return err
	}

	csvReader := csv.NewReader(srcFile)
	header, err := csvReader.Read()

	index := make(map[string]int)
	for i, col := range header {
		index[col] = i
	}
	records, err := csvReader.ReadAll()

	players := []*Player{}
	for _, record := range records {
		firstName := record[index["Full Name"]]
		lastName := record[index["Full Name"]]
		pos := record[index["Position"]]
		switch record[index["Position"]] {
		case "LG":
			pos = "OG"
		case "RG":
			pos = "OG"
		case "LT":
			pos = "OT"
		case "RT":
			pos = "OT"
		case "HB":
			pos = "RB"
		case "FB":
			pos = "RB"
		case "LE":
			pos = "DE"
		case "RE":
			pos = "DE"
		case "LOLB":
			pos = "OLB"
		case "ROLB":
			pos = "OLB"
		case "MLB":
			pos = "ILB"
		case "FS":
			pos = "S"
		case "SS":
			pos = "S"
		default:
			continue
		}

		cur, err := NewPlayer(firstName, lastName, pos)
		if err != nil {
			log.Printf("Invalid player %s\n", record[index["Full Name"]])
			log.Print(err)
		}

		attrs := []Attribute{}
		for attr, _ := range attributes {
			srcAttrName := ""
			switch attr {
			case "SPD":
				srcAttrName = "Speed"
			case "STR":
				srcAttrName = "Strength"
			case "ACC":
				srcAttrName = "Acceleration"
			case "THP":
				srcAttrName = "Throw Power"
			case "CAR":
				srcAttrName = "Carrying"
			case "CAT":
				srcAttrName = "Catching"
			case "TAC":
				srcAttrName = "Tackle"
			case "BLS":
				srcAttrName = "Block Shedding"
			default:
				attrVal, err := getAvgAttr(attr, record, index)
				if err != nil {
					return err
				}
				attrs = append(attrs, Attribute{attr, attrVal})
				continue
			}
			attrVal, err := strconv.Atoi(record[index[srcAttrName]])
			if err != nil {
				return err
			}
			attrs = append(attrs, Attribute{attr, attrVal})
		}

		cur.UpdateAttrs(attrs)

		players = append(players, cur)
	}

	for _, plyr := range players {
		fmt.Println(plyr)
	}

	return nil
}

func getAvgAttr(attr string, record []string, index map[string]int) (int, error) {
	sum := 0
	counter := 0
	switch attr {
	case "THA":
		cur, err := strconv.Atoi(record[index["Throw Accuracy Short"]])
		if err != nil {
			return 0, err
		}
		sum += cur
		cur, err = strconv.Atoi(record[index["Throw Accuracy Mid"]])
		if err != nil {
			return 0, err
		}
		sum += cur
		cur, err = strconv.Atoi(record[index["Throw Accuracy Deep"]])
		if err != nil {
			return 0, err
		}
		sum += cur
		counter = 3
	case "BLK":
		cur, err := strconv.Atoi(record[index["Lead Block"]])
		if err != nil {
			return 0, err
		}
		sum += cur
		cur, err = strconv.Atoi(record[index["Run Block"]])
		if err != nil {
			return 0, err
		}
		sum += cur
		cur, err = strconv.Atoi(record[index["Pass Block"]])
		if err != nil {
			return 0, err
		}
		sum += cur
		counter = 3
	case "COV":
		cur, err := strconv.Atoi(record[index["Man Coverage"]])
		if err != nil {
			return 0, err
		}
		sum += cur
		cur, err = strconv.Atoi(record[index["Zone Coverage"]])
		if err != nil {
			return 0, err
		}
		sum += cur
		counter = 2
	case "PSR":
		cur, err := strconv.Atoi(record[index["Finesse Moves"]])
		if err != nil {
			return 0, err
		}
		sum += cur
		cur, err = strconv.Atoi(record[index["Power Moves"]])
		if err != nil {
			return 0, err
		}
		sum += cur
		counter = 2

	}
	return (sum / counter), nil
}
