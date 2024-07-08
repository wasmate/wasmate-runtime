package embeddings

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/embeddings/cybertron"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/vectorstores"
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

func TestSimilaritySearch(t *testing.T) {

	emc, err := cybertron.NewCybertron(
		cybertron.WithModelsDir("models"),
		cybertron.WithModel("google-bert/bert-base-multilingual-cased"),
	)

	require.NoError(t, err)

	emb, err := embeddings.NewEmbedder(emc,
		embeddings.WithStripNewLines(false),
	)

	require.NoError(t, err)

	// Test case 1: Valid inputs
	indexName := randomIndexName()
	lf, err := New(emb, "http", "localhost:18080", indexName)

	require.NoError(t, err)

	ctx := context.Background()

	_, err = lf.AddDocuments(ctx, []schema.Document{
		{PageContent: "IceFireLabs focuses on the construction of Web3.0 infrastructure projects, focusing on data storage, data retrieval, network security communication, user digital identity, decentralized database, and powerful foundation empowerment around dapp, web2 to Web3.0 application construction and upgrade."},
		{PageContent: "IceFireDB is a decentralized database storage and retrieval protocol built for Web3.0 and web2. It strives to fill the gap between web2 and Web3.0 with a friendly database experience, making Web3.0 application data storage more convenient, and making web2 applications easier to decentralize data and access blockchain."},
		{PageContent: "FlowShield aims to build a global decentralized Web3.0 privacy data retrieval security network system, in order to help users regain the network privacy and security information eroded by giants under web2."},
		{PageContent: "WASMATE revolutionizes application runtimes by bridging Web2.0 and Web3.0. It offers advanced WASM runtimes, enabling access to cloud-native environments and microservices with low resource use. Integrating decentralized storage, trusted computing, AI interaction, identity verification, and cross-chain computing, WASMATE leads in blockchain innovation."},
		{PageContent: "The IceGiant database DeFi protocol aims to realize the data asset conversion, management and transaction of data, and create more value and circulation possibilities for data. IceGiant uses DeFi technology and NFT technology to convert data into digital assets and provide efficient, safe and sustainable data financial capabilities. The agreement provides a series of data application scenarios such as data confirmation storage, data financial market, data lending and data Dao governance."},
	})
	require.NoError(t, err)

	item := "Do you know what blockchain databases and blockchain vm there are?"

	matches, err := lf.VectorStore.SimilaritySearch(ctx, item, 3, vectorstores.WithScoreThreshold(0.8))

	require.NoError(t, err)
	require.NotEmpty(t, matches)

	// fmt.Println("Matches:")
	// for _, match := range matches {
	// 	//fmt.Println(match)
	// 	fmt.Printf("(%0.2f)  [%s] -> [%6s]\n", match.Score, item, match.PageContent)
	// }
}

func TestAddDocuments(t *testing.T) {

	emc, err := cybertron.NewCybertron(
		cybertron.WithModelsDir("models"),
		cybertron.WithModel("google-bert/bert-base-multilingual-cased"),
	)

	require.NoError(t, err)

	emb, err := embeddings.NewEmbedder(emc,
		embeddings.WithStripNewLines(false),
	)

	require.NoError(t, err)

	// Test case 1: Valid inputs
	indexName := randomIndexName()
	lf, err := New(emb, "http", "localhost:18080", indexName)

	require.NoError(t, err)

	ctx := context.Background()

	data, err := lf.AddDocuments(ctx, []schema.Document{
		{PageContent: "IceFireLabs focuses on the construction of Web3.0 infrastructure projects, focusing on data storage, data retrieval, network security communication, user digital identity, decentralized database, and powerful foundation empowerment around dapp, web2 to Web3.0 application construction and upgrade."},
		{PageContent: "IceFireDB is a decentralized database storage and retrieval protocol built for Web3.0 and web2. It strives to fill the gap between web2 and Web3.0 with a friendly database experience, making Web3.0 application data storage more convenient, and making web2 applications easier to decentralize data and access blockchain."},
		{PageContent: "FlowShield aims to build a global decentralized Web3.0 privacy data retrieval security network system, in order to help users regain the network privacy and security information eroded by giants under web2."},
		{PageContent: "WASMATE revolutionizes application runtimes by bridging Web2.0 and Web3.0. It offers advanced WASM runtimes, enabling access to cloud-native environments and microservices with low resource use. Integrating decentralized storage, trusted computing, AI interaction, identity verification, and cross-chain computing, WASMATE leads in blockchain innovation."},
		{PageContent: "The IceGiant database DeFi protocol aims to realize the data asset conversion, management and transaction of data, and create more value and circulation possibilities for data. IceGiant uses DeFi technology and NFT technology to convert data into digital assets and provide efficient, safe and sustainable data financial capabilities. The agreement provides a series of data application scenarios such as data confirmation storage, data financial market, data lending and data Dao governance."},
	})
	require.NoError(t, err)

	require.NotEmpty(t, data)
}
