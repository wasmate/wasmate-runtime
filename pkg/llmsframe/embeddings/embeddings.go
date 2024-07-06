package embeddings

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/vectorstores"
	"github.com/tmc/langchaingo/vectorstores/weaviate"
)

// randomIndexName generates a unique index name for testing purposes.
// It uses a UUID to ensure uniqueness and appends it to the prefix "Test".
// The generated index name will have the format "Test<UUID>".
func randomIndexName() string {
	return "Test" + strings.ReplaceAll(uuid.New().String(), "-", "")
}

/*
	//"github.com/tmc/langchaingo/embeddings/cybertron"
	emc, err := cybertron.NewCybertron(
		cybertron.WithModelsDir("models"),
		//cybertron.WithModel("sentence-transformers/all-MiniLM-L6-v2"),BAAI/bge-small-en-v1.5
		cybertron.WithModel("google-bert/bert-base-multilingual-cased"),
	)
	if err != nil {
		log.Fatal("create embedder client", err)
	}

	// Create an embedder from the previously created client.
	emb, err := embeddings.NewEmbedder(emc,
		embeddings.WithStripNewLines(false),
	)
	if err != nil {
		log.Fatal("create embedder", err)
	}
*/

// EmbeddingsFrame defines the structure for the framework.
type EmbeddingsFrame struct {
	Embedder    embeddings.Embedder
	Scheme      string
	Host        string
	IndexName   string
	VectorStore weaviate.Store
}

// New creates a new instance of EmbeddingsFrame.
func New(embedder embeddings.Embedder, scheme, host, indexName string) (*EmbeddingsFrame, error) {
	client, err := weaviate.New(
		weaviate.WithEmbedder(embedder),
		weaviate.WithScheme(scheme),
		weaviate.WithHost(host),
		weaviate.WithIndexName(indexName),
	)
	if err != nil {
		return nil, fmt.Errorf("create weaviate client: %w", err)
	}

	return &EmbeddingsFrame{
		Embedder:    embedder,
		Scheme:      scheme,
		Host:        host,
		IndexName:   indexName,
		VectorStore: client,
	}, nil
}

// AddDocuments adds a list of documents to the vector store.
func (lf *EmbeddingsFrame) AddDocuments(ctx context.Context, docs []schema.Document) error {
	_, err := lf.VectorStore.AddDocuments(ctx, docs)
	if err != nil {
		return fmt.Errorf("add documents: %w", err)
	}
	return nil
}

// SimilaritySearch performs a similarity search and returns matches.
func (lf *EmbeddingsFrame) SimilaritySearch(ctx context.Context, query string, topK int, threshold float32) ([]schema.Document, error) {
	matches, err := lf.VectorStore.SimilaritySearch(ctx, query, topK, vectorstores.WithScoreThreshold(threshold))
	if err != nil {
		return nil, fmt.Errorf("similarity search: %w", err)
	}
	return matches, nil
}

// PrintMatches prints the search matches to the console.
func PrintMatches(matches []schema.Document, query string) {
	fmt.Println("Matches:")
	for _, match := range matches {
		fmt.Printf("(%0.2f)  [%s] -> [%s]\n", match.Score, query, match.PageContent)
	}
}
