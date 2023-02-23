package urlservice

import (
	"context"
	"net/url"

	"github.com/karalef/shlink/model"
	"github.com/karalef/shlink/repo"
	"github.com/karalef/shlink/urlservice/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// NewURLService creates a new service.
func NewURLService(repo repo.Repository, shortener ...Shortener) *URLService {
	var sh Shortener
	if len(shortener) > 0 {
		sh = shortener[0]
	} else {
		sh = NewCRC64Shortener()
	}
	return &URLService{
		repo:      repo,
		shortener: sh,
	}
}

var _ proto.URLServiceServer = (*URLService)(nil)

// URLService implements proto.URLService interface.
type URLService struct {
	repo      repo.Repository
	shortener Shortener

	proto.UnimplementedURLServiceServer
}

// CreateShort creates a short link.
func (s *URLService) CreateShort(ctx context.Context, origin *proto.Origin) (*proto.Short, error) {
	u := origin.GetUrl()
	if u == "" {
		return nil, status.Error(codes.InvalidArgument, "origin URL is required")
	}
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid URL")
	}

	id := s.shortener.Shorten(u)

	err = s.repo.Store(ctx, id, u)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.Short{Id: id}, nil
}

// GetOrigin returns an origin link by short link.
func (s *URLService) GetOrigin(ctx context.Context, short *proto.Short) (*proto.Origin, error) {
	u, err := s.get(ctx, short)
	if err != nil {
		return nil, err
	}
	return &proto.Origin{Url: u.Origin}, nil
}

func (s *URLService) get(ctx context.Context, short *proto.Short) (*model.URL, error) {
	sh := short.GetId()
	if sh == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	url, err := s.repo.Get(ctx, sh)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if url == nil {
		return nil, status.Error(codes.NotFound, "url not found")
	}
	return url, nil
}

// Get returns an URL by short link.
func (s *URLService) Get(ctx context.Context, short *proto.Short) (*proto.URL, error) {
	u, err := s.get(ctx, short)
	if err != nil {
		return nil, err
	}
	return &proto.URL{
		Short:     u.Short,
		Origin:    u.Origin,
		CreatedAt: timestamppb.New(u.CreatedAt),
		Views:     u.Views,
	}, nil
}
