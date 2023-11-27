package seed

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	s, err := youtube.NewService(ctx, option.WithAPIKey(os.Getenv("YOUTUBE_API_KEY")))

	conn, err := pgx.Connect(context.Background(), os.Getenv("POSTGRES_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var i = 0
	var nextPageToken = ""
	for i < 5 {
		call := s.CommentThreads.List([]string{"snippet"}).AllThreadsRelatedToChannelId("UCti6dG0zSAetLGGYcgNML4Q").MaxResults(100).PageToken(nextPageToken)
		res, err := call.Do()
		nextPageToken = res.NextPageToken

		for _, item := range res.Items {
			
		}
		i++
	}

}
