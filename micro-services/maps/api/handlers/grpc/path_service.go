package grpc

// import (
// 	"context"

// 	"gholi-fly-maps/internal/paths/port"
// 	pb "gholi-fly-maps/api/handlers/grpc" // Import generated protobuf package
// )

// type PathServiceServer struct {
// 	service port.PathService
// 	pb.UnimplementedPathServiceServer
// }

// // NewPathServiceServer creates a new PathServiceServer instance
// func NewPathServiceServer(service port.PathService) *PathServiceServer {
// 	return &PathServiceServer{service: service}
// }

// // GetPath handles the GetPath gRPC call
// func (s *PathServiceServer) GetPath(ctx context.Context, req *pb.GetPathRequest) (*pb.GetPathResponse, error) {
// 	// Fetch path by ID
// 	path, err := s.service.GetByID(ctx, req.GetId())
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Map domain model to gRPC response
// 	return &pb.GetPathResponse{
// 		Id:                  path.ID.String(),
// 		SourceTerminalId:    path.SourceTerminalID.String(),
// 		DestinationTerminalId: path.DestinationTerminalID.String(),
// 		DistanceKm:          path.DistanceKM,
// 		RouteCode:           path.RouteCode,
// 		VehicleType:         path.VehicleType,
// 	}, nil
// }
