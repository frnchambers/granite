package physics

import "os"

func PrintToFile(system *System_t, filename string) {

	file, err := os.Create(filename)
	check(err)
	defer file.Close()

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
