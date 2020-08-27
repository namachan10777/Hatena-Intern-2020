package grpc

import (
	"context"
	"errors"
	"testing"

	pb "github.com/hatena/Hatena-Intern-2020/services/renderer-go/pb/renderer"
	pb_title_fetcher "github.com/hatena/Hatena-Intern-2020/services/renderer-go/pb/title_fetcher"
	"google.golang.org/grpc"
	"github.com/stretchr/testify/assert"
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

func Test_Server_Render(t *testing.T) {
	mock := MockTitleFetcher {}
	s := NewServer(mock)
	src := `foo https://google.com/ bar`
	reply, err := s.Render(context.Background(), &pb.RenderRequest{Src: src})
	assert.NoError(t, err)
	assert.Equal(t, `foo <a href="https://google.com/">https://google.com/</a> bar`, reply.Html)
}
