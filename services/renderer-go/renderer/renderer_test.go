package renderer

import (
	"context"
	"errors"
	"testing"

	pb_title_fetcher "github.com/hatena/Hatena-Intern-2020/services/renderer-go/pb/title_fetcher"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

type MockTitleFetcher struct {
}

func (m MockTitleFetcher) Fetch(ctx context.Context, req *pb_title_fetcher.FetchRequest, options ...grpc.CallOption) (*pb_title_fetcher.FetchReply, error) {
	reply := new(pb_title_fetcher.FetchReply)
	if req.Url == "https://namachan10777.dev" {
		reply.Title = "namachan10777"
		return reply, nil
	}
	return nil, errors.New("Something happened")
}

func Test_Render_Error(t *testing.T) {
	mock := MockTitleFetcher {}
	src := `foo https://google.com/ bar`
	html, err := Render(context.Background(), mock, src)
	assert.NoError(t, err)
	assert.Equal(t, `foo <a href="https://google.com/">https://google.com/</a> bar`, html)
}

func Test_Render_Success(t *testing.T) {
	mock := MockTitleFetcher {}
	src := `foo https://namachan10777.dev bar`
	html, err := Render(context.Background(), mock, src)
	assert.NoError(t, err)
	assert.Equal(t, `foo <a href="https://namachan10777.dev">namachan10777</a> bar`, html)
}
