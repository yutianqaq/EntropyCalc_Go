package main

import (
	"debug/pe"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

func calculateEntropy(buffer []byte) float64 {
	l := float64(len(buffer))
	m := map[byte]float64{}
	for _, b := range buffer {
		m[b]++
	}

	var hm float64
	for _, c := range m {
		hm += c * math.Log2(c)
	}

	return math.Log2(l) - hm/l
}

func calculateFileEntropy(filename string) {
	fileBuffer, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	fileEntropy := calculateEntropy(fileBuffer)

	fmt.Printf("Entropy of \033[36m%s \033[0m as a whole file is: ", filename)

	if fileEntropy >= 5.6 && fileEntropy <= 6.8 {
		fmt.Printf("\033[32m%f\033[0m\n", fileEntropy) // Green - legitimate
	} else if fileEntropy > 7.2 && fileEntropy <= 8.0 {
		fmt.Printf("\033[31m%f\033[0m\n", fileEntropy) // Red - malicious
	} else {
		fmt.Printf("%f\n", fileEntropy)
	}
}

func calculatePESectionEntropy(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	filePE, err := pe.NewFile(file)
	if err != nil {
		log.Fatal(err)
	}

	calculateFileEntropy(filename)

	fmt.Printf("[i] Parsing \033[36m%s\033[0m 's PE Section Headers ...\n", filename)

	colorIndex := 0
	colors := []string{"\033[33m", "\033[32m", "\033[36m", "\033[35m", "\033[34m"}

	for _, section := range filePE.Sections {
		sectionName := string(section.Name[:])
		sectionSize := section.Size

		switch sectionName {
		case ".text", ".data", ".rdata", ".pdata", ".xdata", ".CRT", ".rsrc", ".reloc":
			sectionContent := make([]byte, sectionSize)
			_, err := file.Seek(int64(section.Offset), 0)
			if err != nil {
				log.Fatal(err)
			}
			_, err = io.ReadFull(file, sectionContent)
			if err != nil {
				log.Fatal(err)
			}

			sectionEntropy := calculateEntropy(sectionContent)

			color := colors[colorIndex]
			colorIndex = (colorIndex + 1) % len(colors)

			fmt.Printf("\t>>> %s%s%s Scored Entropy Of Value: %f\033[0m\n", color, "\""+sectionName+"\"", color, sectionEntropy)
		}
	}

}

func main() {
	filename := flag.String("file", "", "File to calculate entropy")
	flag.Parse()

	if *filename == "" {
		flag.Usage()
		return
	}

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = pe.NewFile(file)
	if err == nil {
		calculatePESectionEntropy(*filename)
	} else {
		calculateFileEntropy(*filename)
	}
}
