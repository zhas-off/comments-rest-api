//go:build integration
// +build integration

package database

func TestCommentDatabase(t *testing.T) {
	t.Run("test create comment", func(t *testing.T) {
		fmt.Println("testing the creation of comments")
	})
}
