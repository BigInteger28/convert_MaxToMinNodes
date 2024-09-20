package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Globale variabele voor levels
var levelNodes = [17]uint64{0, 5, 15, 25, 35, 45, 55, 65, 75, 85, 95, 110, 125, 140, 160, 180, 200}

// Functie om minimumNodes te berekenen op basis van maxNodes en level
func getMinimum(maxNodes uint64, level uint64) uint64 {
	if level == 1 {
		return 1
	} else if level == 2 {
		return uint64(float64(maxNodes) * 0.8)
	}
	eightyPercent := uint64(float64(maxNodes) * 0.8)
	if eightyPercent < getNodesFromLevel(level-2) {
		return getNodesFromLevel(level - 2)
	} else {
		return eightyPercent
	}
}

// Functie om het aantal nodes op basis van het level te krijgen
func getNodesFromLevel(level uint64) uint64 {
	if level < 18 {
		return levelNodes[level-1]
	}
	if level < 38 {
		return 200 + ((level - 17) * 40)
	}
	return 1000 + ((level - 37) * 50)
}

func main() {
	// Open input en output bestanden
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Fout bij het openen van het inputbestand:", err)
		return
	}
	defer inputFile.Close()

	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Fout bij het maken van het outputbestand:", err)
		return
	}
	defer outputFile.Close()

	// Lees regel voor regel uit het inputbestand
	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)

	for scanner.Scan() {
		line := scanner.Text()

		// Splits de regel in naam, level, minNodes, maxNodes en comment, met 3 spaties als scheiding
		parts := strings.Split(line, "   ")

		if len(parts) < 4 {
			// Onvoldoende velden, schrijf de regel zoals deze is
			writer.WriteString(line + "\n")
			continue
		}

		// Parse LEVEL en MAXNODES
		level, _ := strconv.ParseUint(strings.TrimSpace(parts[1]), 10, 64)
		maxNodes, _ := strconv.ParseUint(strings.TrimSpace(parts[3]), 10, 64)

		// Bereken minimum nodes als maxNodes groter is dan 0
		if maxNodes > 0 {
			minNodes := getMinimum(maxNodes, level)
			parts[2] = strconv.FormatUint(minNodes, 10) // Update MINNODES
		}

		// Voeg de delen opnieuw samen met precies 3 spaties tussen de velden
		updatedLine := fmt.Sprintf("%s   %s   %s   %s", parts[0], parts[1], parts[2], parts[3])

		// Als er een commentaar is, voeg dat toe
		if len(parts) > 4 {
			updatedLine += "   " + parts[4]
		}

		// Schrijf de aangepaste regel naar het outputbestand
		writer.WriteString(updatedLine + "\n")
	}

	// Schrijf naar output file en sluit de writer
	writer.Flush()

	fmt.Println("Verwerking voltooid, output opgeslagen in output.txt")
}
