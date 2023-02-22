package urlservice

import (
	"context"

	"github.com/karalef/shlink/urlservice/proto"
	"google.golang.org/grpc"
)

// Service represents wrapped grpc service.
type Service interface {
	CreateShort(ctx context.Context, origin string, opts ...grpc.CallOption) (string, error)
	GetOrigin(ctx context.Context, short string, opts ...grpc.CallOption) (string, error)
	Get(ctx context.Context, short string, opts ...grpc.CallOption) (*proto.URL, error)
}

// NewService creates and wraps a new service.
func NewService(conn *grpc.ClientConn) Service {
	return &srvcWrapper{
		proto.NewURLServiceClient(conn),
	}
}

var _ Service = (*srvcWrapper)(nil)

type srvcWrapper struct {
	srvc proto.URLServiceClient
}

func (w *srvcWrapper) CreateShort(ctx context.Context, origin string, opts ...grpc.CallOption) (string, error) {
	sh, err := w.srvc.CreateShort(ctx, &proto.Origin{Url: origin}, opts...)
	return sh.GetId(), err
}

func (w *srvcWrapper) GetOrigin(ctx context.Context, short string, opts ...grpc.CallOption) (string, error) {
	orig, err := w.srvc.GetOrigin(ctx, &proto.Short{Id: short}, opts...)
	return orig.GetUrl(), err
}

func (w *srvcWrapper) Get(ctx context.Context, short string, opts ...grpc.CallOption) (*proto.URL, error) {
	return w.srvc.Get(ctx, &proto.Short{Id: short}, opts...)
}
