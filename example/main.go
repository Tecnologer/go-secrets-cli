package main

import (
	"fmt"

	"github.com/tecnologer/go-secrets"
	"github.com/tecnologer/go-secrets-cli/example/testdeep"
	"github.com/tecnologer/go-secrets/config"
)

func main() {
	secrets.InitWithConfig(&config.Config{EncryptionEnabled: false})

	// if err != nil {
	// 	panic(err)
	// }

	// secrets.Set("SQL.Username", "tecno")
	// secrets.Set("SQL.pwd", "123")
	// secrets.Set("SQL.host", "localhost")
	// secrets.Set("SQL.database", "test")

	sql, err := testdeep.GetGroup("SQL")
	if err == nil {
		fmt.Println("SQL keys:")
		fmt.Printf("Server=%v;Database=%v;User Id=%v;Password=%v;\n", sql.Get("host"), sql.Get("database"), sql.Get("Username"), sql.Get("pwd"))
	}

	fmt.Println("All keys:")
	bucket, err := secrets.Get()
	if err != nil {
		panic(err)
	}

	for key, val := range bucket.Secrets {
		fmt.Printf("%s: %v\n", key, val)
	}

	key := "SQL.pwd"
	fmt.Printf("Get key in other package. {%s: %v}\n", key, testdeep.GetKey(key))

	fmt.Println(bucket.ID)
}
