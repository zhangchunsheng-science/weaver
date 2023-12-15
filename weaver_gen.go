// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package weaver

import (
	"context"
	"errors"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"github.com/ServiceWeaver/weaver/runtime/protos"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:  "github.com/ServiceWeaver/weaver/Logger",
		Iface: reflect.TypeOf((*Logger)(nil)).Elem(),
		Impl:  reflect.TypeOf(stderrLogger{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return logger_local_stub{impl: impl.(Logger), tracer: tracer, logBatchMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/Logger", Method: "LogBatch", Remote: false})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return logger_client_stub{stub: stub, logBatchMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/Logger", Method: "LogBatch", Remote: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return logger_server_stub{impl: impl.(Logger), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return logger_reflect_stub{caller: caller}
		},
		RefData: "",
	})
	codegen.Register(codegen.Registration{
		Name:  "github.com/ServiceWeaver/weaver/controller",
		Iface: reflect.TypeOf((*controller)(nil)).Elem(),
		Impl:  reflect.TypeOf(noopController{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return controller_local_stub{impl: impl.(controller), tracer: tracer, getHealthMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/controller", Method: "GetHealth", Remote: false}), updateComponentsMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/controller", Method: "UpdateComponents", Remote: false})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return controller_client_stub{stub: stub, getHealthMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/controller", Method: "GetHealth", Remote: true}), updateComponentsMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/controller", Method: "UpdateComponents", Remote: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return controller_server_stub{impl: impl.(controller), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return controller_reflect_stub{caller: caller}
		},
		RefData: "",
	})
}

// weaver.InstanceOf checks.
var _ InstanceOf[Logger] = (*stderrLogger)(nil)
var _ InstanceOf[controller] = (*noopController)(nil)

// weaver.Router checks.
var _ Unrouted = (*stderrLogger)(nil)
var _ Unrouted = (*noopController)(nil)

// Local stub implementations.

type logger_local_stub struct {
	impl            Logger
	tracer          trace.Tracer
	logBatchMetrics *codegen.MethodMetrics
}

// Check that logger_local_stub implements the Logger interface.
var _ Logger = (*logger_local_stub)(nil)

func (s logger_local_stub) LogBatch(ctx context.Context, a0 *protos.LogEntryBatch) (err error) {
	// Update metrics.
	begin := s.logBatchMetrics.Begin()
	defer func() { s.logBatchMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "weaver.Logger.LogBatch", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.LogBatch(ctx, a0)
}

type controller_local_stub struct {
	impl                    controller
	tracer                  trace.Tracer
	getHealthMetrics        *codegen.MethodMetrics
	updateComponentsMetrics *codegen.MethodMetrics
}

// Check that controller_local_stub implements the controller interface.
var _ controller = (*controller_local_stub)(nil)

func (s controller_local_stub) GetHealth(ctx context.Context, a0 *protos.GetHealthRequest) (r0 *protos.GetHealthReply, err error) {
	// Update metrics.
	begin := s.getHealthMetrics.Begin()
	defer func() { s.getHealthMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "weaver.controller.GetHealth", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetHealth(ctx, a0)
}

func (s controller_local_stub) UpdateComponents(ctx context.Context, a0 *protos.UpdateComponentsRequest) (r0 *protos.UpdateComponentsReply, err error) {
	// Update metrics.
	begin := s.updateComponentsMetrics.Begin()
	defer func() { s.updateComponentsMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "weaver.controller.UpdateComponents", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.UpdateComponents(ctx, a0)
}

// Client stub implementations.

type logger_client_stub struct {
	stub            codegen.Stub
	logBatchMetrics *codegen.MethodMetrics
}

// Check that logger_client_stub implements the Logger interface.
var _ Logger = (*logger_client_stub)(nil)

func (s logger_client_stub) LogBatch(ctx context.Context, a0 *protos.LogEntryBatch) (err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.logBatchMetrics.Begin()
	defer func() { s.logBatchMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "weaver.Logger.LogBatch", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Encode arguments.
	enc := codegen.NewEncoder()
	serviceweaver_enc_ptr_LogEntryBatch_fec9a5d4(enc, a0)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

type controller_client_stub struct {
	stub                    codegen.Stub
	getHealthMetrics        *codegen.MethodMetrics
	updateComponentsMetrics *codegen.MethodMetrics
}

// Check that controller_client_stub implements the controller interface.
var _ controller = (*controller_client_stub)(nil)

func (s controller_client_stub) GetHealth(ctx context.Context, a0 *protos.GetHealthRequest) (r0 *protos.GetHealthReply, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.getHealthMetrics.Begin()
	defer func() { s.getHealthMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "weaver.controller.GetHealth", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Encode arguments.
	enc := codegen.NewEncoder()
	serviceweaver_enc_ptr_GetHealthRequest_fd6083fb(enc, a0)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_ptr_GetHealthReply_b2d11423(dec)
	err = dec.Error()
	return
}

func (s controller_client_stub) UpdateComponents(ctx context.Context, a0 *protos.UpdateComponentsRequest) (r0 *protos.UpdateComponentsReply, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.updateComponentsMetrics.Begin()
	defer func() { s.updateComponentsMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "weaver.controller.UpdateComponents", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Encode arguments.
	enc := codegen.NewEncoder()
	serviceweaver_enc_ptr_UpdateComponentsRequest_d1b56e1f(enc, a0)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 1, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_ptr_UpdateComponentsReply_93bebb77(dec)
	err = dec.Error()
	return
}

// Note that "weaver generate" will always generate the error message below.
// Everything is okay. The error message is only relevant if you see it when
// you run "go build" or "go run".
var _ codegen.LatestVersion = codegen.Version[[0][20]struct{}](`

ERROR: You generated this file with 'weaver generate' (devel) (codegen
version v0.20.0). The generated code is incompatible with the version of the
github.com/ServiceWeaver/weaver module that you're using. The weaver module
version can be found in your go.mod file or by running the following command.

    go list -m github.com/ServiceWeaver/weaver

We recommend updating the weaver module and the 'weaver generate' command by
running the following.

    go get github.com/ServiceWeaver/weaver@latest
    go install github.com/ServiceWeaver/weaver/cmd/weaver@latest

Then, re-run 'weaver generate' and re-build your code. If the problem persists,
please file an issue at https://github.com/ServiceWeaver/weaver/issues.

`)

// Server stub implementations.

type logger_server_stub struct {
	impl    Logger
	addLoad func(key uint64, load float64)
}

// Check that logger_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*logger_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s logger_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "LogBatch":
		return s.logBatch
	default:
		return nil
	}
}

func (s logger_server_stub) logBatch(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 *protos.LogEntryBatch
	a0 = serviceweaver_dec_ptr_LogEntryBatch_fec9a5d4(dec)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.LogBatch(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

type controller_server_stub struct {
	impl    controller
	addLoad func(key uint64, load float64)
}

// Check that controller_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*controller_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s controller_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "GetHealth":
		return s.getHealth
	case "UpdateComponents":
		return s.updateComponents
	default:
		return nil
	}
}

func (s controller_server_stub) getHealth(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 *protos.GetHealthRequest
	a0 = serviceweaver_dec_ptr_GetHealthRequest_fd6083fb(dec)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.GetHealth(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_ptr_GetHealthReply_b2d11423(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s controller_server_stub) updateComponents(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 *protos.UpdateComponentsRequest
	a0 = serviceweaver_dec_ptr_UpdateComponentsRequest_d1b56e1f(dec)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.UpdateComponents(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_ptr_UpdateComponentsReply_93bebb77(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

// Reflect stub implementations.

type logger_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that logger_reflect_stub implements the Logger interface.
var _ Logger = (*logger_reflect_stub)(nil)

func (s logger_reflect_stub) LogBatch(ctx context.Context, a0 *protos.LogEntryBatch) (err error) {
	err = s.caller("LogBatch", ctx, []any{a0}, []any{})
	return
}

type controller_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that controller_reflect_stub implements the controller interface.
var _ controller = (*controller_reflect_stub)(nil)

func (s controller_reflect_stub) GetHealth(ctx context.Context, a0 *protos.GetHealthRequest) (r0 *protos.GetHealthReply, err error) {
	err = s.caller("GetHealth", ctx, []any{a0}, []any{&r0})
	return
}

func (s controller_reflect_stub) UpdateComponents(ctx context.Context, a0 *protos.UpdateComponentsRequest) (r0 *protos.UpdateComponentsReply, err error) {
	err = s.caller("UpdateComponents", ctx, []any{a0}, []any{&r0})
	return
}

// Encoding/decoding implementations.

func serviceweaver_enc_ptr_LogEntryBatch_fec9a5d4(enc *codegen.Encoder, arg *protos.LogEntryBatch) {
	if arg == nil {
		enc.Bool(false)
	} else {
		enc.Bool(true)
		enc.EncodeProto(arg)
	}
}

func serviceweaver_dec_ptr_LogEntryBatch_fec9a5d4(dec *codegen.Decoder) *protos.LogEntryBatch {
	if !dec.Bool() {
		return nil
	}
	var res protos.LogEntryBatch
	dec.DecodeProto(&res)
	return &res
}

func serviceweaver_enc_ptr_GetHealthRequest_fd6083fb(enc *codegen.Encoder, arg *protos.GetHealthRequest) {
	if arg == nil {
		enc.Bool(false)
	} else {
		enc.Bool(true)
		enc.EncodeProto(arg)
	}
}

func serviceweaver_dec_ptr_GetHealthRequest_fd6083fb(dec *codegen.Decoder) *protos.GetHealthRequest {
	if !dec.Bool() {
		return nil
	}
	var res protos.GetHealthRequest
	dec.DecodeProto(&res)
	return &res
}

func serviceweaver_enc_ptr_GetHealthReply_b2d11423(enc *codegen.Encoder, arg *protos.GetHealthReply) {
	if arg == nil {
		enc.Bool(false)
	} else {
		enc.Bool(true)
		enc.EncodeProto(arg)
	}
}

func serviceweaver_dec_ptr_GetHealthReply_b2d11423(dec *codegen.Decoder) *protos.GetHealthReply {
	if !dec.Bool() {
		return nil
	}
	var res protos.GetHealthReply
	dec.DecodeProto(&res)
	return &res
}

func serviceweaver_enc_ptr_UpdateComponentsRequest_d1b56e1f(enc *codegen.Encoder, arg *protos.UpdateComponentsRequest) {
	if arg == nil {
		enc.Bool(false)
	} else {
		enc.Bool(true)
		enc.EncodeProto(arg)
	}
}

func serviceweaver_dec_ptr_UpdateComponentsRequest_d1b56e1f(dec *codegen.Decoder) *protos.UpdateComponentsRequest {
	if !dec.Bool() {
		return nil
	}
	var res protos.UpdateComponentsRequest
	dec.DecodeProto(&res)
	return &res
}

func serviceweaver_enc_ptr_UpdateComponentsReply_93bebb77(enc *codegen.Encoder, arg *protos.UpdateComponentsReply) {
	if arg == nil {
		enc.Bool(false)
	} else {
		enc.Bool(true)
		enc.EncodeProto(arg)
	}
}

func serviceweaver_dec_ptr_UpdateComponentsReply_93bebb77(dec *codegen.Decoder) *protos.UpdateComponentsReply {
	if !dec.Bool() {
		return nil
	}
	var res protos.UpdateComponentsReply
	dec.DecodeProto(&res)
	return &res
}
