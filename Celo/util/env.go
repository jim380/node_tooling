package util

import "os"

// SetEnv sets the necessary env variables
// needs a rewrite. too tired to do it today
func SetEnv() {
	celoImage := os.Getenv("CELO_IMAGE")
	os.Setenv("CELO_IMAGE", celoImage)
	networkID := os.Getenv("NETWORK_ID")
	os.Setenv("NETWORK_ID", networkID)
	celoValAddr := os.Getenv("CELO_VALIDATOR_ADDRESS")
	os.Setenv("CELO_VALIDATOR_ADDRESS", celoValAddr)
	celoValGroupAddr := os.Getenv("CELO_VALIDATOR_GROUP_ADDRESS")
	os.Setenv("CELO_VALIDATOR_GROUP_ADDRESS", celoValGroupAddr)
	celoValSignerAddr := os.Getenv("CELO_VALIDATOR_SIGNER_ADDRESS")
	os.Setenv("CELO_VALIDATOR_SIGNER_ADDRESS", celoValSignerAddr)
	celoValSignerPubKey := os.Getenv("CELO_VALIDATOR_SIGNER_PUBLIC_KEY")
	os.Setenv("CELO_VALIDATOR_SIGNER_PUBLIC_KEY", celoValSignerPubKey)
	celoValSignerSig := os.Getenv("CELO_VALIDATOR_SIGNER_SIGNATURE")
	os.Setenv("CELO_VALIDATOR_SIGNER_SIGNATURE", celoValSignerSig)
	celoValSignerBlsPubKey := os.Getenv("CELO_VALIDATOR_SIGNER_BLS_PUBLIC_KEY")
	os.Setenv("CELO_VALIDATOR_SIGNER_BLS_PUBLIC_KEY", celoValSignerBlsPubKey)
	celoValSignerBlsSig := os.Getenv("CELO_VALIDATOR_SIGNER_BLS_SIGNATURE")
	os.Setenv("CELO_VALIDATOR_SIGNER_BLS_SIGNATURE", celoValSignerBlsSig)
	celoProxyEnode := os.Getenv("PROXY_ENODE")
	os.Setenv("PROXY_ENODE", celoProxyEnode)
	celoProxyExternalIP := os.Getenv("PROXY_EXTERNAL_IP")
	os.Setenv("PROXY_EXTERNAL_IP", celoProxyExternalIP)
	celoProxyInternalIP := os.Getenv("PROXY_INTERNAL_IP")
	os.Setenv("PROXY_INTERNAL_IP", celoProxyInternalIP)
	celoValName := os.Getenv("VALIDATOR_NAME")
	os.Setenv("VALIDATOR_NAME", celoValName)

	// if celoImage == "" || networkID == "" {
	// 	log.Fatal("Missing fields in .env file")
	// }
}
