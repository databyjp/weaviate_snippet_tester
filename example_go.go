package main

import (
	"context"
	"fmt"
	"os"

	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/auth"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/graphql"
)

func main() {
	// // // For local client
	// cfg := weaviate.Config{
	// 	Host:   "localhost:8080", // Replace with your endpoint
	// 	Scheme: "http",
	// 	Headers: map[string]string{
	// 		"X-OpenAI-Api-Key":      os.Getenv("OPENAI_APIKEY"),
	// 		"X-Cohere-Api-Key":      os.Getenv("COHERE_APIKEY"),
	// 		"X-HuggingFace-Api-Key": os.Getenv("HUGGINGFACE_APIKEY"),
	// 	}, // Replace with your inference API key
	// }

	// For edu-demo client
	cfg := weaviate.Config{
		Host:       "edu-demo.weaviate.network", // Replace with your endpoint
		Scheme:     "https",
		AuthConfig: auth.ApiKey{Value: "learn-weaviate"}, // Replace w/ your Weaviate instance API key
		Headers: map[string]string{
			"X-OpenAI-Api-Key":      os.Getenv("OPENAI_APIKEY"),
			"X-Cohere-Api-Key":      os.Getenv("COHERE_APIKEY"),
			"X-HuggingFace-Api-Key": os.Getenv("HUGGINGFACE_APIKEY"),
		}, // Replace with your inference API key
	}

	client, err := weaviate.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	fields := []graphql.Field{
		{Name: "question"},
		{Name: "answer"},
	}

	nearText := client.GraphQL().
		NearTextArgBuilder().
		WithConcepts([]string{"biology"})

	result, err := client.GraphQL().Get().
		WithClassName("JeopardyQuestion").
		WithFields(fields...).
		WithNearText(nearText).
		WithLimit(2).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", result)
}
