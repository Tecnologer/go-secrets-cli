package testdeep

import "github.com/tecnologer/go-secrets"

//GetKey hdyii
func GetKey(k string) interface{} {
	return secrets.GetKey(k)
}

//GetGroup sss
func GetGroup(group string) (secrets.Secret, error) {
	return secrets.GetGroup(group)
}
