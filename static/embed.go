package static

import (
	_ "embed"
)

//go:embed tls/ca_crt.pem
var CaCert []byte

//go:embed tls/ca_key.pem
var CaKey []byte

//go:embed tls/server_self_crt.pem
var ServerSelfCert []byte

//go:embed tls/server_self_key.pem
var ServerSelfKey []byte
