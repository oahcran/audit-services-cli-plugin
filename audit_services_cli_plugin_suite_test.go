package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAuditServicesCliPlugin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AuditServicesCliPlugin Suite")
}
