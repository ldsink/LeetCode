package main

func squareIsWhite(coordinates string) bool {
	return int(int(coordinates[0]-'a')%2+int(coordinates[1]-'1')%2)%2 == 1
}
