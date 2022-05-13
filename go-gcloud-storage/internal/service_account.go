package internal

import (
	"encoding/json"
	"log"
)

type ServiceAccount interface {
	GetGoogleAccessID() string
	GetPrivateKey() string
}
type serviceAccount struct {
	Type                    string `json:"type"`
	ProjectId               string `json:"project_id"`
	PrivateKeyId            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientId                string `json:"client_id"`
	AuthUri                 string `json:"auth_uri"`
	TokenUri                string `json:"token_uri"`
	AuthProviderX509CertUrl string `json:"auth_provider_x509_cert_url"`
	ClientX509CertUrl       string `json:"client_x509_cert_url"`
}

func NewServiceAccount(jsonData []byte) ServiceAccount {
	sa := new(serviceAccount)
	err := json.Unmarshal(jsonData, sa)
	if err != nil {
		log.Panicln("unable to decode service account from json: ", err.Error())
	}
	return sa
}

func (sa *serviceAccount) GetGoogleAccessID() string {
	return sa.ClientEmail
}

func (sa *serviceAccount) GetPrivateKey() string {
	return sa.PrivateKey
}
