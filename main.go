package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/color"
)

var (
	// Valeur qui stockera le nom du fichier à tester
	executable = flag.String("executable", "executables", "path to the executable folder")

	// Valeur qui stockera le nom du dictionnaire à ouvrir
	dictionary = flag.String("dictionary", "dict.txt", "path to the dictionary to use")

	// Valeur de la solution (DEBUG)
	word = flag.String("word", "default", "[debug] word to put as a solution")

	// Valeur qui stockera le nombre de lettre dans le mot
	size = flag.Int("size", 5, "size of the word to find")

	// Valeur qui stockera le nombre de lettre dans le mot
	games = flag.Int("games", 1, "how many games to play")
)

func main() {
	fmt.Println(color.RedString("[INITIALISATION]"))

	// On récupére depuis l'appel du programme la valeur de l'argument -executable
	// Exemple : testeur -executable=wordchamp_solver
	flag.Parse()

	// On vérifie que le dossier existe bien
	_, err := os.Stat(*executable)
	if err != nil {
		// Il existe pas là donc on panique
		log.Panic("[DOSSIER EXECUTABLE MANQUANT, ARGUMENT -executable=...]")
	}

	fmt.Println("	Dossier executable trouvé")

	// On vérifie que le fichier du dictionnaire existe bien
	_, err = os.Stat(*dictionary)
	if err != nil {
		// Il existe pas là donc on panique
		log.Panic("[DICTIONNAIRE MANQUANT, ARGUMENT -dictionary=...]")
	}

	fmt.Println("	Dictionnaire trouvé")

	ioutil.WriteFile(fmt.Sprintf("%s/wsolf.txt", *executable), []byte(fmt.Sprintf("%d", *size)), 0644)

	fmt.Println("	Taille du mot :", *size)

	files, err := ioutil.ReadDir(*executable)
	if err != nil {
		panic(err)
	}

	fmt.Println("	Liste des fichiers trouvé")
	fmt.Println()
	fmt.Println(color.RedString("[LANCEMENT DES TESTS]"))

	for _, file := range files {
		if file.Name() == "wsolf.txt" {
			continue
		}
		exec := fmt.Sprintf("%s/%s", *executable, file.Name())

		fmt.Println(color.GreenString("	[TEST] " + exec))

		runGame(exec, *games)
	}

	fmt.Println(color.GreenString("	[TOP]"))

	for _, c := range OrderCountersByAverage() {
		fmt.Printf("		%s : %dW / %dL (%f)\n", c.Name, c.WinCount, c.Total-c.WinCount, c.AverageTry)
	}
}
