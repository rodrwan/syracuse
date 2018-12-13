

proto p:
	echo "[proto] Generating golang proto..."
	@protoc  -I citizens/ citizens/citizens.proto --go_out=plugins=grpc:citizens
