all:
	make configs
	make protocol

protocol:
	protoc --proto_path=msg/def --go_out=msg/gen msg/def/*.proto

configs:
	utils/tools configs configs/def/

sync_submodules:
	git submodule sync