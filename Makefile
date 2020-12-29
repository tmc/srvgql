generate:
	go generate ./ent
	@patch -p0 < ent/node.go.patch
	go generate ./graph
