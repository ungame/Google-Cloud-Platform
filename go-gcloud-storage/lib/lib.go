package lib

import (
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	"github.com/joho/godotenv"
	"go-gcloud-storage/env"
	"go-gcloud-storage/internal"
	"go-gcloud-storage/utils"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Panicln("unable to load env vars: ", err.Error())
	}
}

type GoogleCloudStorage interface {
	Create(ctx context.Context, filename string, file []byte) error
	ReadFile(ctx context.Context, filename string) ([]byte, error)
	ShareFileWithTimeout(_ context.Context, filename string, timeout time.Duration) (string, error)
	Close()
}

type googleCloudStorage struct {
	client      *storage.Client
	credentials *google.Credentials
	bucket      string
}

func New() GoogleCloudStorage {
	sa, err := ioutil.ReadFile(os.Getenv(env.CREDENTIALS_PATH))
	if err != nil {
		log.Panicln("unable to read credentials file: ", err.Error())
	}

	ctx := context.Background()

	credentials, err := google.CredentialsFromJSON(ctx, sa, storage.ScopeReadWrite)
	if err != nil {
		log.Panicln("unable to create credentials from json: ", err.Error())
	}

	client, err := storage.NewClient(ctx, option.WithCredentials(credentials))
	if err != nil {
		log.Panicln("unable to create storage client: ", err.Error())
	}

	return &googleCloudStorage{
		client:      client,
		credentials: credentials,
		bucket:      os.Getenv(env.DEFAULT_BUCKET),
	}
}

func (gcs *googleCloudStorage) Close() {
	utils.HandleClose(gcs.client)
}

func (gcs *googleCloudStorage) ReadFile(ctx context.Context, filename string) ([]byte, error) {
	objectHandle := gcs.object(filename)
	reader, err := objectHandle.NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer utils.HandleClose(reader)
	var buf bytes.Buffer
	_, err = buf.ReadFrom(reader)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (gcs *googleCloudStorage) Create(ctx context.Context, filename string, file []byte) error {
	objectHandle := gcs.object(filename)
	writer := objectHandle.NewWriter(ctx)
	_, err := writer.Write(file)
	if err != nil {
		return err
	}
	return writer.Close()
}

func (gcs *googleCloudStorage) ShareFileWithTimeout(_ context.Context, filename string, timeout time.Duration) (string, error) {
	bucketHandle := gcs.client.Bucket(gcs.bucket)

	sa := internal.NewServiceAccount(gcs.credentials.JSON)

	opts := &storage.SignedURLOptions{
		Scheme:         storage.SigningSchemeV4,
		GoogleAccessID: sa.GetGoogleAccessID(),
		PrivateKey:     []byte(sa.GetPrivateKey()),
		Method:         http.MethodGet,
		Expires:        time.Now().Add(timeout),
	}

	return bucketHandle.SignedURL(filename, opts)
}

func (gcs *googleCloudStorage) object(name string) *storage.ObjectHandle {
	return gcs.client.Bucket(gcs.bucket).Object(name)
}
