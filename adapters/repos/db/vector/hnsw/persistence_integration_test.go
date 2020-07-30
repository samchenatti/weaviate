//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// +build integrationTest

package hnsw

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHnswPersistence(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	dirName := fmt.Sprintf("./testdata/%d", rand.Intn(10000000))
	os.MkdirAll(dirName, 0777)
	defer func() {
		err := os.RemoveAll(dirName)
		fmt.Println(err)
	}()

	cl := NewCommitLogger(dirName, "integrationtest")
	index, err := New(dirName, "integrationtest", cl, 30, 60, testVectorForID)
	require.Nil(t, err)

	for i, vec := range testVectors {
		err := index.Add(i, vec)
		require.Nil(t, err)
	}

	// see index_test.go for more context
	expectedResults := []int{
		3, 5, 4, // cluster 2
		7, 8, 6, // cluster 3
		2, 1, 0, // cluster 1
	}

	t.Run("verify that the results match originally", func(t *testing.T) {
		position := 3
		res, err := index.knnSearchByVector(testVectors[position], 50, 36)
		require.Nil(t, err)
		assert.Equal(t, expectedResults, res)
	})

	// destory the index
	index = nil

	// build a new index from the (uncondensed) commit log
	secondIndex, err := New(dirName, "integrationtest", cl, 30, 60,
		testVectorForID)
	require.Nil(t, err)

	t.Run("verify that the results match after rebuiling from disk",
		func(t *testing.T) {
			position := 3
			res, err := secondIndex.knnSearchByVector(testVectors[position], 50, 36)
			require.Nil(t, err)
			assert.Equal(t, expectedResults, res)
		})
}