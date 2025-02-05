package main

import (
	"flag"
	"log"
	"time"

	utils "github.com/hossainsmshakib/GoKage/utils"

)

func main() {
	var dumpPath, query string
	flag.StringVar(&dumpPath, "p", "enwiki-latest-abstract1.xml.gz", "wiki abstract dump path")
	flag.StringVar(&query, "q", "Small wild cat", "search query")
	flag.Parse()

	log.Println("Running Full Text Search")

	start := time.Now()
	docs, err := utils.LoadDocuments(dumpPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))

	start = time.Now()
	idx := make(utils.Index)
	idx.Add(docs)
	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))

	start = time.Now()
	matchedIDs := idx.Search(query)
	elapsedNano := time.Since(start).Nanoseconds() // Get nanoseconds
	
	// If elapsed time is zero, use at least 1 ns
	if elapsedNano == 0 {
		elapsedNano = 1
	}
	
	elapsedMicro := float64(elapsedNano) / 1e3  // Convert to µs
	elapsedMilli := float64(elapsedNano) / 1e6  // Convert to ms
	elapsedSec := float64(elapsedNano) / 1e9    // Convert to sec
	
	if elapsedNano < 1000 {
		log.Printf("Search found %d documents in %.0f ns", len(matchedIDs), float64(elapsedNano))
	} else if elapsedNano < 1e6 {
		log.Printf("Search found %d documents in %.2f µs", len(matchedIDs), elapsedMicro)
	} else if elapsedNano < 1e9 {
		log.Printf("Search found %d documents in %.3f ms", len(matchedIDs), elapsedMilli)
	} else {
		log.Printf("Search found %d documents in %.9f seconds", len(matchedIDs), elapsedSec)
	}

	for _, id := range matchedIDs {
		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)
	}
}