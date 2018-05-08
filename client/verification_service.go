package client

import (
	"log"
)

// VerificationService is a wrapper for the Pact Provider Verifier Service.
type VerificationService struct {
	ServiceManager
}

// NewService creates a new VerificationService with default settings.
// Arguments allowed:
//
// 		--provider-base-url
// 		--pact-urls
// 		--provider-states-url
// 		--provider-states-setup-url
// 		--broker-username
// 		--broker-password
//    --publish-verification-results
//    --provider-app-version
//    --custom-provider-headers
func (v *VerificationService) NewService(args []string) Service {
	log.Printf("[DEBUG] starting verification service with args: %v\n", args)

	v.Args = args
	v.Cmd = getVerifierCommandPath()

	return v
}

func getVerifierCommandPath() string {
	return "pact-provider-verifier"
}