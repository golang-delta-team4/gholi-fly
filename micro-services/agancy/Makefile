PROTOC_GEN_GO := protoc -I=./ --go_out=./ --go_opt=paths=source_relative

AGENCY_PROTO := ./api/pb/agency.proto

.PHONY: agency-proto
agency-proto:
	$(PROTOC_GEN_GO) $(AGENCY_PROTO)

.PHONY: proto-all
proto-all: agency-proto
