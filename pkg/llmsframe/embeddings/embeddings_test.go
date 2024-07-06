package embeddings

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/embeddings/cybertron"
)

func TestNew(t *testing.T) {
	// Create a mock embedder
	emc, err := cybertron.NewCybertron(
		cybertron.WithModelsDir("models"),
		cybertron.WithModel("google-bert/bert-base-multilingual-cased"),
	)

	if err != nil {
		t.Errorf("New function failed with error: %v", err)
	}

	emb, err := embeddings.NewEmbedder(emc,
		embeddings.WithStripNewLines(false),
	)

	if err != nil {
		t.Errorf("New function failed with error: %v", err)
	}

	mockEmbedder := emb

	// Test case 1: Valid inputs
	indexName := randomIndexName()
	_, err = New(mockEmbedder, "http", "localhost", indexName)
	if err != nil {
		t.Errorf("New function failed with error: %v", err)
	}

	// Test case 2: Invalid scheme
	_, err = New(mockEmbedder, "", "localhost", indexName)
	if err == nil {
		t.Errorf("New function did not fail with invalid scheme")
	}

	// Test case 3: Invalid host
	_, err = New(mockEmbedder, "http", "", indexName)
	if err == nil {
		t.Errorf("New function did not fail with invalid host")
	}

	// Test case 4: Invalid index name
	_, err = New(mockEmbedder, "http", "localhost", "")
	if err == nil {
		t.Errorf("New function did not fail with invalid index name")
	}
}

func TestRandomIndexName(t *testing.T) {
	for i := 0; i < 5; i++ {
		indexName := randomIndexName()
		require.NotEmpty(t, indexName)
		require.Contains(t, indexName, "Test")
		//require.True(t, strings.ContainsAny(indexName, "0123456789abcdef"))
	}
}

func TestRandomIndexNameUnique(t *testing.T) {
	indexName1 := randomIndexName()
	indexName2 := randomIndexName()
	require.NotEqual(t, indexName1, indexName2)
}

func TestRandomIndexNameLength(t *testing.T) {
	indexName := randomIndexName()
	require.GreaterOrEqual(t, len(indexName), 7)
	require.LessOrEqual(t, len(indexName), 36)
}

func TestRandomIndexNameUUIDFormat(t *testing.T) {
	indexName := randomIndexName()
	require.True(t, strings.ContainsAny(indexName, "Test"))
	//require.True(t, strings.ContainsAny(indexName, "0123456789abcdef"))
}

func TestRandomIndexNameWithPrefix(t *testing.T) {
	indexName := randomIndexName()
	require.True(t, strings.HasPrefix(indexName, "Test"))
}
