// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: byzq.proto

package byzq

// import (
// 	bytes "bytes"
// 	context "context"
// 	fmt "fmt"
// 	math "math"
// 	time "time"

// 	_ "github.com/gogo/protobuf/gogoproto"
// 	proto "github.com/gogo/protobuf/proto"
// 	_ "github.com/relab/gorums"
// 	"golang.org/x/net/trace"
// 	grpc "google.golang.org/grpc"
// 	codes "google.golang.org/grpc/codes"
// 	status "google.golang.org/grpc/status"
// )

// // Reference imports to suppress errors if they are not otherwise used.
// var _ = proto.Marshal
// var _ = fmt.Errorf
// var _ = math.Inf

// const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// func init() {
// 	proto.RegisterType((*Value)(nil), "byzq.Value")
// 	proto.RegisterType((*Content)(nil), "byzq.Content")
// 	proto.RegisterType((*WriteResponse)(nil), "byzq.WriteResponse")
// }

// func init() { proto.RegisterFile("byzq.proto", fileDescriptor_f0624bb1535ad1f6) }

// // Reference Gorums specific imports to suppress errors if they are not otherwise used.
// var _ = codes.OK
// var _ = bytes.MinRead

// /* Code generated by protoc-gen-gorums - template source file: calltype_datatypes.tmpl */

// /* Code generated by protoc-gen-gorums - template source file: calltype_quorumcall.tmpl */

// /* Exported types and methods for quorum call method EchoWrite */

// // EchoWrite is invoked as a quorum call on all nodes in configuration c,
// // using the same argument arg, and returns the result.
// func (c *Configuration) EchoWrite(ctx context.Context, a *Value) (resp *WriteResponse, err error) {
// 	var ti traceInfo
// 	if c.mgr.opts.trace {
// 		ti.Trace = trace.New("gorums."+c.tstring()+".Sent", "EchoWrite")
// 		defer ti.Finish()

// 		ti.firstLine.cid = c.id
// 		if deadline, ok := ctx.Deadline(); ok {
// 			ti.firstLine.deadline = time.Until(deadline)
// 		}
// 		ti.LazyLog(&ti.firstLine, false)
// 		ti.LazyLog(&payload{sent: true, msg: a}, false)

// 		defer func() {
// 			ti.LazyLog(&qcresult{
// 				reply: resp,
// 				err:   err,
// 			}, false)
// 			if err != nil {
// 				ti.SetError()
// 			}
// 		}()
// 	}

// 	expected := c.n
// 	replyChan := make(chan internalWriteResponse, expected)
// 	for _, n := range c.nodes {
// 		go callGRPCEchoWrite(ctx, n, a, replyChan)
// 	}

// 	var (
// 		replyValues = make([]*WriteResponse, 0, expected)
// 		errs        []GRPCError
// 		quorum      bool
// 	)

// 	for {
// 		select {
// 		case r := <-replyChan:
// 			if r.err != nil {
// 				errs = append(errs, GRPCError{r.nid, r.err})
// 				break
// 			}
// 			if c.mgr.opts.trace {
// 				ti.LazyLog(&payload{sent: false, id: r.nid, msg: r.reply}, false)
// 			}
// 			replyValues = append(replyValues, r.reply)
// 			if resp, quorum = c.qspec.EchoWriteQF(a, replyValues); quorum {
// 				return resp, nil
// 			}
// 		case <-ctx.Done():
// 			return resp, QuorumCallError{ctx.Err().Error(), len(replyValues), errs}
// 		}

// 		if len(errs)+len(replyValues) == expected {
// 			return resp, QuorumCallError{"incomplete call", len(replyValues), errs}
// 		}
// 	}
// }

// func callGRPCEchoWrite(ctx context.Context, node *Node, arg *Value, replyChan chan<- internalWriteResponse) {
// 	reply := new(WriteResponse)
// 	start := time.Now()
// 	err := grpc.Invoke(
// 		ctx,
// 		"/byzq.Storage/EchoWrite",
// 		arg,
// 		reply,
// 		node.conn,
// 	)
// 	s, ok := status.FromError(err)
// 	if ok && (s.Code() == codes.OK || s.Code() == codes.Canceled) {
// 		node.setLatency(time.Since(start))
// 	} else {
// 		node.setLastErr(err)
// 	}
// 	replyChan <- internalWriteResponse{node.id, reply, err}
// }

// /* Code generated by protoc-gen-gorums - template source file: qspec.tmpl */

// // QuorumSpec is the interface that wraps every quorum function.
// type QuorumSpec interface {
// 	// EchoWriteQF is the quorum function for the EchoWrite
// 	// quorum call method.
// 	EchoWriteQF(req *Value, replies []*WriteResponse) (*WriteResponse, bool)
// }

// // Reference imports to suppress errors if they are not otherwise used.
// var _ context.Context
// var _ grpc.ClientConn

// // This is a compile-time assertion to ensure that this generated file
// // is compatible with the grpc package it is being compiled against.
// const _ = grpc.SupportPackageIsVersion4

// func (c *storageClient) EchoWrite(ctx context.Context, in *Value, opts ...grpc.CallOption) (*WriteResponse, error) {
// 	out := new(WriteResponse)
// 	err := c.cc.Invoke(ctx, "/byzq.Storage/EchoWrite", in, out, opts...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return out, nil
// }

// // StorageServer is the server API for Storage service.
// type StorageServer interface {
// 	EchoWrite(context.Context, *Value) (*WriteResponse, error)
// }

// func (*UnimplementedStorageServer) EchoWrite(ctx context.Context, req *Value) (*WriteResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method EchoWrite not implemented")
// }

// func _Storage_EchoWrite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
// 	in := new(Value)
// 	if err := dec(in); err != nil {
// 		return nil, err
// 	}
// 	if interceptor == nil {
// 		return srv.(StorageServer).EchoWrite(ctx, in)
// 	}
// 	info := &grpc.UnaryServerInfo{
// 		Server:     srv,
// 		FullMethod: "/byzq.Storage/EchoWrite",
// 	}
// 	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
// 		return srv.(StorageServer).EchoWrite(ctx, req.(*Value))
// 	}
// 	return interceptor(ctx, in, info, handler)
// }

// var _Storage_serviceDesc = grpc.ServiceDesc{
// 	ServiceName: "byzq.Storage",
// 	HandlerType: (*StorageServer)(nil),
// 	Methods: []grpc.MethodDesc{
// 		{
// 			MethodName: "EchoWrite",
// 			Handler:    _Storage_EchoWrite_Handler,
// 		},
// 	},
// 	Streams:  []grpc.StreamDesc{},
// 	Metadata: "byzq.proto",
// }

// func (m *Value) MarshalToSizedBuffer(dAtA []byte) (int, error) {
// 	i := len(dAtA)
// 	_ = i
// 	var l int
// 	_ = l
// 	if len(m.SignatureS) > 0 {
// 		i -= len(m.SignatureS)
// 		copy(dAtA[i:], m.SignatureS)
// 		i = encodeVarintByzq(dAtA, i, uint64(len(m.SignatureS)))
// 		i--
// 		dAtA[i] = 0x1a
// 	}
// 	if len(m.SignatureR) > 0 {
// 		i -= len(m.SignatureR)
// 		copy(dAtA[i:], m.SignatureR)
// 		i = encodeVarintByzq(dAtA, i, uint64(len(m.SignatureR)))
// 		i--
// 		dAtA[i] = 0x12
// 	}
// 	if m.C != nil {
// 		{
// 			size, err := m.C.MarshalToSizedBuffer(dAtA[:i])
// 			if err != nil {
// 				return 0, err
// 			}
// 			i -= size
// 			i = encodeVarintByzq(dAtA, i, uint64(size))
// 		}
// 		i--
// 		dAtA[i] = 0xa
// 	}
// 	return len(dAtA) - i, nil
// }
