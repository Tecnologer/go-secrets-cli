package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/tecnologer/go-secrets"
	"github.com/tecnologer/go-secrets/config"
)

// var idFlag = flag.String("id", "", "Id fo the bucket")
var keyFlag = flag.String("key", "", "Secret key")
var valueFlag = flag.String("val", "", "Secret value")
var encryptEnabled = flag.Bool("enc", false, "Flag to encrypt or not the secret bucket")

var (
	bucket              *secrets.Bucket
	currentPath         string
	localSecretFilePath string
)

const initSecretFilePath = ".secretid"

func init() {
	currentPath, _ = os.Getwd()

	if currentPath == "" {
		currentPath = "."
	}
	localSecretFilePath = fmt.Sprintf("%s/%s", currentPath, initSecretFilePath)
	// fmt.Println(localSecretFilePath)
}

func main() {
	if len(os.Args) == 2 {
		action := strings.ToLower(os.Args[1])
		switch action {
		case "init":
			initBucket()
		case "--help":
			fallthrough
		case "-help":
			fallthrough
		case "help":
			help()
		default:
			fmt.Println("Invalid action. Type `go-secret-cli help` for more info")
		}
		return
	}

	if len(os.Args) < 3 {
		fmt.Println("Invalid action. Type go-secret-cli -help for more info")
		return
	}

	flag.CommandLine.Parse(os.Args[2:])

	bucketID, err := getBucketIDFromFile()
	if err != nil {
		log.Fatalf("Invalid bucket id. Init bucket with: `go-secrets-cli init`. Error: %v", err)
	}

	bucket, err = secrets.GetWithConfig(&config.Config{BucketID: bucketID, EncryptionEnabled: *encryptEnabled})
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
	fmt.Println("\tgo-secrets-cli set -key <string> -value <string>")

	fmt.Println("* Get secret:")
	fmt.Println("\tgo-secrets-cli get [-key <string>]")

	fmt.Println("* Remove secret:")
	fmt.Println("\tgo-secrets-cli remove -key <string>")

	fmt.Println("* Init secret:")
	fmt.Println("\tgo-secrets-cli init")

	fmt.Println("* Enable Encryption: Flag to indicate if the bucket is encrypted. Default: False")
	fmt.Println("\tgo-secrets-cli <set|get|remove|init> -enc")
}

func initBucket() {
	var err error
	bucketID := uuid.New()
	write := true
	if secretExists(localSecretFilePath) {
		bucketID, err = getBucketIDFromFile()

		if err != nil {
			write = true
			bucketID = uuid.New()
		}
	}

	if write {
		ioutil.WriteFile(localSecretFilePath, []byte(bucketID.String()), 0644)
	}

	fmt.Printf("Secret initialized with id \"%v\"\n", bucketID)
}

func secretExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func getBucketIDFromFile() (uuid.UUID, error) {
	file, err := ioutil.ReadFile(localSecretFilePath)

	if err != nil {
		return uuid.UUID{}, errors.Wrap(err, fmt.Sprintf("error reading file %s", localSecretFilePath))
	}

	return uuid.Parse(string(file))
}
