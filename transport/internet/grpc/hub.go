package grpc

import (
	"context"
	"time"

	"github.com/xtls/xray-core/common"
	"github.com/xtls/xray-core/common/net"
	"github.com/xtls/xray-core/common/session"
	"github.com/xtls/xray-core/transport/internet"
	"github.com/xtls/xray-core/transport/internet/grpc/encoding"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type Listener struct {
	encoding.UnimplementedGRPCServiceServer
	ctx     context.Context
	handler internet.ConnHandler
	local   net.Addr
	config  *Config

	s *grpc.Server
}

func (l Listener) Tun(server encoding.GRPCService_TunServer) error {
	tunCtx, cancel := context.WithCancel(l.ctx)
	l.handler(encoding.NewHunkConn(server, cancel))
	<-tunCtx.Done()
	return nil
}

func (l Listener) TunMulti(server encoding.GRPCService_TunMultiServer) error {
	tunCtx, cancel := context.WithCancel(l.ctx)
	l.handler(encoding.NewMultiHunkConn(server, cancel))
	<-tunCtx.Done()
	return nil
}

func (l Listener) Close() error {
	l.s.Stop()
	return nil
}

func (l Listener) Addr() net.Addr {
	return l.local
}

func Listen(ctx context.Context, address net.Address, port net.Port, settings *internet.MemoryStreamConfig, handler internet.ConnHandler) (internet.Listener, error) {
	grpcSettings := settings.ProtocolSettings.(*Config)
	var listener *Listener
	if port == net.Port(0) { // unix
		listener = &Listener{
			handler: handler,
			local: &net.UnixAddr{
				Name: address.Domain(),
				Net:  "unix",
			},
			config: grpcSettings,
		}
	} else { // tcp
		listener = &Listener{
			handler: handler,
			local: &net.TCPAddr{
				IP:   address.IP(),
				Port: int(port),
			},
			config: grpcSettings,
		}
	}

	listener.ctx = ctx

	var options []grpc.ServerOption
	var s *grpc.Server
	if grpcSettings.IdleTimeout > 0 || grpcSettings.HealthCheckTimeout > 0 {
		options = append(options, grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    time.Second * time.Duration(grpcSettings.IdleTimeout),
			Timeout: time.Second * time.Duration(grpcSettings.HealthCheckTimeout),
		}))
	}

	s = grpc.NewServer(options...)
	listener.s = s

	go func() {
		var streamListener net.Listener
		var err error
		if port == net.Port(0) { // unix
			streamListener, err = internet.ListenSystem(ctx, &net.UnixAddr{
				Name: address.Domain(),
				Net:  "unix",
			}, settings.SocketSettings)
			if err != nil {
				newError("failed to listen on ", address).Base(err).AtError().WriteToLog(session.ExportIDToError(ctx))
				return
			}
		} else { // tcp
			streamListener, err = internet.ListenSystem(ctx, &net.TCPAddr{
				IP:   address.IP(),
				Port: int(port),
			}, settings.SocketSettings)
			if err != nil {
				newError("failed to listen on ", address, ":", port).Base(err).AtError().WriteToLog(session.ExportIDToError(ctx))
				return
			}
		}

		newError("gRPC listen for service name `" + grpcSettings.getServiceName() + "` tun `" + grpcSettings.getTunStreamName() + "` multi tun `" + grpcSettings.getTunMultiStreamName() + "`").AtDebug().WriteToLog()
		encoding.RegisterGRPCServiceServerX(s, listener, grpcSettings.getServiceName(), grpcSettings.getTunStreamName(), grpcSettings.getTunMultiStreamName())

		streamListener = settings.ToSecuredListener(streamListener)

		if err = s.Serve(streamListener); err != nil {
			newError("Listener for gRPC ended").Base(err).WriteToLog()
		}
	}()

	return listener, nil
}

func init() {
	common.Must(internet.RegisterTransportListener(protocolName, Listen))
}
