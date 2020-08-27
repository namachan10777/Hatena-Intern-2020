package renderer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	pb_title_fetcher "github.com/hatena/Hatena-Intern-2020/services/renderer-go/pb/title_fetcher"
)

type MockTitleFetcher struct {
}

func (m MockTitleFetcher) Fetch(ctx context.Context, req *pb_title_fetcher.FetchRequest, options ...grpc.CallOption) (*pb_title_fetcher.FetchReply, error) {
	reply := pb_title_fetcher.FetchReply {
		Title: "hoge",
	}
	return &reply, nil
}

func Test_Render(t *testing.T) {
	mock := MockTitleFetcher {}
	src := `foo https://google.com/ bar`
	html, err := Render(context.Background(), mock, src)
	assert.NoError(t, err)
	assert.Equal(t, `foo <a href="https://google.com/">https://google.com/</a> bar`, html)
}
