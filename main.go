package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/tecnologer/go-secrets"
)

var idFlag = flag.String("id", "", "Id fo the bucket")
var keyFlag = flag.String("key", "", "Secret key")
var valueFlag = flag.String("val", "", "Secret value")

var bucket *secrets.Bucket

func main() {
	if len(os.Args) == 2 && (strings.ToLower(os.Args[1]) == "help" || strings.ToLower(os.Args[1]) == "-help") {
		help()
		return
	}

	if len(os.Args) < 3 {
		fmt.Println("Invalid action. Type go-secret-cli -help for more info")
		return
	}

	flag.CommandLine.Parse(os.Args[2:])

	bucketID, err := uuid.Parse(*idFlag)
	if err != nil {
		log.Fatalf("Invalid bucket id. Error: %v", err)
	}
	bucket, err = secrets.GetBucket(bucketID)
	if err != nil {
		log.Fatalf("Error getting the bucket: %v", err)
	}

	action := strings.ToLower(os.Args[1])
	switch action {
	case "set":
		set(*keyFlag, *valueFlag)
	case "get":
		get(*keyFlag)
	case "remove":
		remove(*keyFlag)
	case "help":
		fallthrough
	case "-help":
		help()
	default:
		fmt.Println("Invalid action. Type `go-secret-cli help` for more info")
		return
	}
}

func set(key string, value interface{}) {
	if key == "" {
		log.Fatalf("Invalid empty key")
	}
	bucket.Set(key, value)
}

func get(key string) {
	if key != "" {
		fmt.Printf("Key: %s\nValue: %v\n", key, bucket.Get(key))
		return
	}

	for key, val := range bucket.Secrets {
		fmt.Printf("%s: %v\n", key, val)
	}
}

func remove(key string) {
	if key == "" {
		log.Fatalf("Invalid empty key")
	}
	bucket.Remove(key)
	fmt.Printf("Key \"%s\" removed\n", key)
}

func help() {
	fmt.Println("* Set new secret:")
	fmt.Println("\tgo-secrets-cli set -id <uuid> -key <string> -value <string>")

	fmt.Println("* Get secret:")
	fmt.Println("\tgo-secrets-cli get -id <uuid> [-key <string>]")

	fmt.Println("* Remove secret:")
	fmt.Println("\tgo-secrets-cli remove -id <uuid> -key <string>")
}
