package forms

import (
	"context"
	"encoding/json"
	"log"
	"net"

	formsgrpc "github.com/coopgo/forms/grpc"
	jtd "github.com/jsontypedef/json-typedef-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type GRPCTransport struct {
	Addr   string
	Server *grpc.Server
}

// NewRESTTransport creates a new transport using the REST protocol
func NewGRPCTransport(addr string) *GRPCTransport {
	grpcServer := grpc.NewServer()

	return &GRPCTransport{
		Addr:   addr,
		Server: grpcServer,
	}
}

func (t GRPCTransport) Run(s *Service) {
	formsgrpc.RegisterFormsServiceServer(t.Server, NewGRPCFormServiceServer(s))
	reflection.Register(t.Server)
	l, err := net.Listen("tcp", t.Addr)
	if err != nil {
		log.Fatal(err)
	}
	if err := t.Server.Serve(l); err != nil {
		log.Fatal(err)
	}
}

type GRPCFormsServiceServer struct {
	Service *Service
	formsgrpc.UnimplementedFormsServiceServer
}

func NewGRPCFormServiceServer(s *Service) *GRPCFormsServiceServer {
	return &GRPCFormsServiceServer{
		Service: s,
	}
}

func (s *GRPCFormsServiceServer) AddForm(context context.Context, req *formsgrpc.AddFormRequest) (*formsgrpc.AddFormResponse, error) {
	err := s.Service.AddForm(form(req.Form))
	if err != nil {
		return nil, err
	}

	return &formsgrpc.AddFormResponse{
		Form: req.Form,
	}, nil
}
func (s *GRPCFormsServiceServer) DeleteForm(context context.Context, req *formsgrpc.DeleteFormRequest) (*formsgrpc.DeleteFormResponse, error) {
	if err := s.Service.DeleteForm(req.Formid); err != nil {
		return nil, err
	}

	return &formsgrpc.DeleteFormResponse{}, nil
}
func (s *GRPCFormsServiceServer) GetForms(context.Context, *formsgrpc.GetFormsRequest) (*formsgrpc.GetFormsResponse, error) {
	forms, err := s.Service.GetForms()
	if err != nil {
		return nil, err
	}
	grpcforms := []*formsgrpc.GRPCForm{}

	for _, f := range forms {
		gf := grpcForm(f)
		grpcforms = append(grpcforms, &gf)
	}
	return &formsgrpc.GetFormsResponse{
		Forms: grpcforms,
	}, nil
}
func (s *GRPCFormsServiceServer) GetForm(c context.Context, req *formsgrpc.GetFormRequest) (*formsgrpc.GetFormResponse, error) {
	form, err := s.Service.GetForm(req.Formid)
	if err != nil {
		return nil, err
	}
	gf := grpcForm(form)
	return &formsgrpc.GetFormResponse{
		Form: &gf,
	}, nil
}
func (s *GRPCFormsServiceServer) SubmitResponse(context context.Context, req *formsgrpc.SubmitResponseRequest) (*formsgrpc.SubmitResponseResponse, error) {
	var response Response
	if err := json.Unmarshal([]byte(req.Response), &response); err != nil {
		return nil, err
	}
	if err := s.Service.SubmitResponse(req.Formid, response); err != nil {
		return nil, err
	}
	return &formsgrpc.SubmitResponseResponse{}, nil
}
func (s *GRPCFormsServiceServer) GetFormResponses(context context.Context, req *formsgrpc.GetFormResponsesRequest) (*formsgrpc.GetFormResponsesResponse, error) {
	responses, err := s.Service.GetFormResponses(req.Formid)
	if err != nil {
		return nil, err
	}
	grpcresponses := []string{}
	for _, r := range responses {
		stringresponse, _ := json.Marshal(r)
		grpcresponses = append(grpcresponses, string(stringresponse))
	}
	return &formsgrpc.GetFormResponsesResponse{
		Responses: grpcresponses,
	}, nil
}
func (s *GRPCFormsServiceServer) AddFormBackend(context.Context, *formsgrpc.AddFormBackendRequest) (*formsgrpc.AddFormBackendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFormBackend not implemented")
}
func (s *GRPCFormsServiceServer) GetFormBackends(context.Context, *formsgrpc.GetFormBackendsRequest) (*formsgrpc.GetFormBackendsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFormBackends not implemented")
}
func (s *GRPCFormsServiceServer) mustEmbedUnimplementedFormsServiceServer() {
}

func grpcForm(f Form) formsgrpc.GRPCForm {
	schema, _ := json.Marshal(f.GetSchema())
	return formsgrpc.GRPCForm{
		Id:     f.ID(),
		Type:   f.Type(),
		Schema: string(schema),
		//AdditionalBackends: f.GetAdditionalBackends(),
	}
}

func form(pf *formsgrpc.GRPCForm) Form {
	if pf.Type == "Structured" {
		var schema jtd.Schema
		json.Unmarshal([]byte(pf.Schema), &schema)
		return NewStructuredForm(pf.Id, &schema)
	}
	return NewUnstructuredForm(pf.Id)
}
