package cf_test

import (
	"github.com/cloudfoundry/cli/plugin/pluginfakes"
	"github.com/oahcran/audit-services-cli-plugin/cf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func FileToString(fileName string) ([]string, error) {
	path, err := filepath.Abs(filepath.Join("testdata", fileName))
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(buf), "\n"), nil
}

var _ = Describe("Cf", func() {

	var (
		cliConn *pluginfakes.FakeCliConnection
		client  cf.CfClient
	)

	BeforeEach(func() {
		cliConn = &pluginfakes.FakeCliConnection{}
		client = &cf.ClientHelper{}
	})

	Describe("List Services", func() {
		var servicesJSON []string
		var err error

		It("should return two services", func() {

			servicesJSON, err = FileToString("services.json")
			Expect(err).ToNot(HaveOccurred())

			cliConn.CliCommandWithoutTerminalOutputReturns(servicesJSON, nil)
			services, _ := client.ListServices(cliConn)
			Expect(len(services)).To(Equal(12))
		})

	})

	Describe("List Orgs", func() {
		var orgsJSON []string
		var err error

		It("should return two services", func() {

			orgsJSON, err = FileToString("orgs.json")
			Expect(err).ToNot(HaveOccurred())

			cliConn.CliCommandWithoutTerminalOutputReturns(orgsJSON, nil)
			orgs, _ := client.ListOrgs(cliConn)
			Expect(len(orgs)).To(Equal(1))
		})

	})

	Describe("List Service Plans", func() {
		var servicePlansJSON []string
		var err error

		It("should return two services", func() {

			servicePlansJSON, err = FileToString("service_plans.json")
			Expect(err).ToNot(HaveOccurred())

			cliConn.CliCommandWithoutTerminalOutputReturns(servicePlansJSON, nil)
			servicePlans, _ := client.ListServicePlans(cliConn)
			Expect(len(servicePlans)).To(Equal(3))
		})

	})

	Describe("List Spaces By Org", func() {
		var spacesJSON []string
		var err error

		It("should return two services", func() {

			spacesJSON, err = FileToString("spaces.json")
			Expect(err).ToNot(HaveOccurred())

			cliConn.CliCommandWithoutTerminalOutputReturns(spacesJSON, nil)
			spaces, _ := client.ListSpacesByOrg(cliConn, "/v2/organizations/ca4a88f7-f89b-4f99-8292-b01a1fb6f190/spaces")
			Expect(len(spaces)).To(Equal(2))
		})

	})

	Describe("Get Service Instances By Url", func() {
		var serviceInstancesJSON []string
		var err error

		It("should return two services", func() {

			serviceInstancesJSON, err = FileToString("service_instances.json")
			Expect(err).ToNot(HaveOccurred())

			cliConn.CliCommandWithoutTerminalOutputReturns(serviceInstancesJSON, nil)
			si, _ := client.GetServiceInstancesByUrl(cliConn, "/v2/service_plans/38579721-1e1f-46eb-aa94-8b599cf14160/service_instances")
			Expect(len(si)).To(Equal(6))
		})

	})

	Describe("Get Service Instances By Url - Empty", func() {
		var serviceInstancesJSON []string
		var err error

		It("should return two services", func() {

			serviceInstancesJSON, err = FileToString("service_instances_empty.json")
			Expect(err).ToNot(HaveOccurred())

			cliConn.CliCommandWithoutTerminalOutputReturns(serviceInstancesJSON, nil)
			si, _ := client.GetServiceInstancesByUrl(cliConn, "/v2/service_plans/38579721-1e1f-46eb-aa94-8b599cf14160/service_instances")
			Expect(len(si)).To(Equal(0))
		})

	})

	Describe("Get Service Summary By Space", func() {
		var summaryJSON []string
		var err error

		It("should return two services", func() {

			summaryJSON, err = FileToString("summary.json")
			Expect(err).ToNot(HaveOccurred())

			cliConn.CliCommandWithoutTerminalOutputReturns(summaryJSON, nil)
			spaces, _ := client.GetServiceSummaryBySpace(cliConn, "/v2/spaces/01b830f3-eaa2-42b1-8547-14af266a20b8/summary")
			Expect(len(spaces)).To(Equal(6))
		})

	})

})
