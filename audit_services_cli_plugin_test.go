package main_test

import (
	"code.cloudfoundry.org/cli/plugin/pluginfakes"
	"github.com/oahcran/audit-services-cli-plugin/cf"
	"github.com/oahcran/audit-services-cli-plugin/cf/cffakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/oahcran/audit-services-cli-plugin"
)

var _ = Describe("AuditServicesCliPlugin", func() {
	var cliConn *pluginfakes.FakeCliConnection
	var cliClient *cffakes.FakeCfClient

	var plugin *AuditServicesPlugin

	BeforeEach(func() {
		cliConn = &pluginfakes.FakeCliConnection{}
		cliClient = &cffakes.FakeCfClient{}
		plugin = &AuditServicesPlugin{
			CfClient: cliClient,
		}
	})

	Describe("AuditServices", func() {

		It("should return a service", func() {
			args := []string{"audit-service-instances"}

			err := plugin.AuditServices(cliConn, args)
			Expect(err).To(BeNil())

		})

	})

	Describe("RetrieveServices return nothing", func() {

		It("should return 0 result", func() {
			args := []string{"audit-service-instances"}

			info, err := plugin.RetrieveServices(cliConn, args)
			Expect(err).To(BeNil())
			Expect(len(info)).To(Equal(0))

		})

	})

	Describe("RetrieveServices return result", func() {
		service := cf.Service{
			Name: "service-name-1",
			Guid: "service-guid-1",
		}

		servicePlan := cf.ServicePlan{
			Name:        "service-plan-name-1",
			Guid:        "service-plan-guid-1",
			ServiceGuid: "service-guid-1",
		}

		org := cf.Org{
			Name:      "org-name-1",
			Guid:      "org-guid-1",
			SpacesUrl: "space-url-1",
		}

		space := cf.Space{
			Name:                "space-name-1",
			Guid:                "space-guid-1",
			Url:                 "space-url-1",
			OrgGuid:             "org-guid-1",
			ServiceInstancesUrl: "service-instances-url-1",
		}

		svcSummary := cf.ServiceSummaryBySpace{
			Name:          "service-instance-name-1",
			Guid:          "service-instance-guid-1",
			BoundAppCount: 2,
		}

		serviceInstance := cf.ServiceInstance{
			Name:            "service-instance-name-1",
			Guid:            "service-instance-guid-1",
			ServicePlanGuid: "service-plan-guid-1",
		}

		BeforeEach(func() {
			cliClient.ListServicesReturns([]cf.Service{service}, nil)
			cliClient.ListServicePlansReturns([]cf.ServicePlan{servicePlan}, nil)
			cliClient.ListOrgsReturns([]cf.Org{org}, nil)
			cliClient.ListSpacesByOrgReturns([]cf.Space{space}, nil)
			cliClient.GetServiceSummaryBySpaceReturns([]cf.ServiceSummaryBySpace{svcSummary}, nil)
			cliClient.GetServiceInstancesByUrlReturns([]cf.ServiceInstance{serviceInstance}, nil)
		})

		It("should return 1 result", func() {
			args := []string{"audit-service-instances"}

			info, err := plugin.RetrieveServices(cliConn, args)
			Expect(err).To(BeNil())
			Expect(len(info)).To(Equal(1))

			err = plugin.AuditServices(cliConn, args)
			Expect(err).To(BeNil())
		})

	})

})
