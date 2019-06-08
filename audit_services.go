package main

import (
	"code.cloudfoundry.org/cli/plugin"
	"errors"
	"fmt"
	"os"
	"sort"

	"github.com/oahcran/audit-services-cli-plugin/cf"
	"github.com/olekukonko/tablewriter"
)

type AuditServicesPlugin struct {
	CfClient cf.CfClient
}

func main() {
	plugin.Start(&AuditServicesPlugin{
		CfClient: &cf.ClientHelper{},
	})
}

func (c *AuditServicesPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	// Ensure that we called the command basic-plugin-command
	if args[0] == "audit-service-instances" {
		//fmt.Println("Running the audit-service")
		err := c.AuditServices(cliConnection, args)
		if err != nil {
			fmt.Println(err)

			os.Exit(1)
		}
	}
}

func (c *AuditServicesPlugin) GetMetadata() plugin.PluginMetadata {

	return plugin.PluginMetadata{
		Name: "audit-services",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 1,
			Build: 0,
		},
		MinCliVersion: plugin.VersionType{
			Major: 6,
			Minor: 7,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "audit-service-instances",
				HelpText: "Audit the services instances for given Cloud Foundry env",

				// UsageDetails is optional
				// It is used to show help of usage of each command
				UsageDetails: plugin.Usage{
					Usage: "cf audit-service-instances [--csv]",
				},
			},
		},
	}
}

func parseArgs(args []string) (csv bool, err error) {
	if len(args) == 1 {
		return false, nil
	}

	if len(args) > 2 {
		return false, errors.New("wrong arguments provided")
	}

	flag := args[1]
	if flag != "--csv" {
		return false, errors.New("wrong arguments provided")
	}

	return true, nil
}

type ServiceInstanceInfo struct {
	InstanceName string
	InstanceGuid string
	OrgName      string
	SpaceName    string
	ServiceName  string
	ServicePlan  string
	NumBoundApps int
}

type ServicePlanInfo struct {
	Name        string
	Guid        string
	ServiceName string
}

func (c *AuditServicesPlugin) AuditServices(cliConnection plugin.CliConnection, args []string) error {

	csvFormat, err := parseArgs(args)

	if err != nil {
		return err
	}

	serviceInstanceStat, err := c.RetrieveServices(cliConnection, args)

	if err != nil {
		return err
	}

	sort.Sort(sortServiceInstances(serviceInstanceStat))

	if csvFormat != true {

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Org", "Space", "Service Name", "Service Plan", "Service Instance Name", "Service Instance Guid", "Bound App Count"})

		for _, v := range serviceInstanceStat {
			table.Append(v.toValueList())
		}

		table.Render()
	} else {

		fmt.Printf("%s,%s,%s,%s,%s,%s,%s\n", "Org", "Space", "Service Name", "Service Plan", "Service Instance Name", "Service Instance Guid", "Bound App Count")

		for _, val := range serviceInstanceStat {
			fmt.Printf("%s,%s,%s,%s,%s,%s,%d\n", val.OrgName, val.SpaceName, val.ServiceName, val.ServicePlan, val.InstanceName, val.InstanceGuid, val.NumBoundApps)
		}
	}

	return nil
}

func (c *AuditServicesPlugin) RetrieveServices(cliConnection plugin.CliConnection, args []string) ([]ServiceInstanceInfo, error) {

	// get services and service plans information
	servicesMap := make(map[string]string)
	servicePlanMap := make(map[string]ServicePlanInfo)

	services, err := c.CfClient.ListServices(cliConnection)
	if err != nil {
		return nil, err
	}

	for _, service := range services {
		servicesMap[service.Guid] = service.Name
	}

	plans, err := c.CfClient.ListServicePlans(cliConnection)
	if err != nil {
		return nil, err
	}

	for _, servicePlan := range plans {
		servicePlanMap[servicePlan.Guid] = ServicePlanInfo{
			Name:        servicePlan.Name,
			Guid:        servicePlan.Guid,
			ServiceName: servicesMap[servicePlan.ServiceGuid],
		}
	}

	// get orgs and spaces, start with service instance url under the space
	orgs, err := c.CfClient.ListOrgs(cliConnection)
	if err != nil {
		return nil, err
	}

	var serviceInstanceStat []ServiceInstanceInfo

	for _, org := range orgs {
		//fmt.Printf("name: %s, guid: %s, spaces url: %s\n", org.Name, org.Guid, org.SpacesUrl)

		spaces, err := c.CfClient.ListSpacesByOrg(cliConnection, org.SpacesUrl)

		if err != nil {
			return nil, err
		}

		// TODO: when set up multi goroutines it would fail, to be investigated.
		//swg := sizedwaitgroup.New(1)

		for _, space := range spaces {
			//fmt.Printf("name: %s, guid: %s, service instances url: %s\n", space.Name, space.Guid, space.ServiceInstancesUrl)

			summaries, err := c.CfClient.GetServiceSummaryBySpace(cliConnection, space.Url)
			if err != nil {
				return nil, err
			}

			serviceBoundAppMaps := make(map[string]int)

			for _, summary := range summaries {
				serviceBoundAppMaps[summary.Guid] = summary.BoundAppCount
			}

			instances, err := c.CfClient.GetServiceInstancesByUrl(cliConnection, space.ServiceInstancesUrl)
			if err != nil {
				return nil, err
			}

			for _, serviceInstance := range instances {
				//fmt.Printf("service instance name: %s\n", serviceInstance.Name)

				siStat := ServiceInstanceInfo{
					InstanceName: serviceInstance.Name,
					InstanceGuid: serviceInstance.Guid,
					OrgName:      org.Name,
					SpaceName:    space.Name,
					ServiceName:  servicePlanMap[serviceInstance.ServicePlanGuid].ServiceName,
					ServicePlan:  servicePlanMap[serviceInstance.ServicePlanGuid].Name,
					NumBoundApps: serviceBoundAppMaps[serviceInstance.Guid],
				}

				serviceInstanceStat = append(serviceInstanceStat, siStat)
			}

		}

	}

	return serviceInstanceStat, nil
}

func (s *ServiceInstanceInfo) toValueList() []string {
	return []string{s.OrgName, s.SpaceName, s.ServiceName, s.ServicePlan, s.InstanceName, s.InstanceGuid, fmt.Sprintf("%d", s.NumBoundApps)}
}

type sortServiceInstances []ServiceInstanceInfo

func (s sortServiceInstances) Len() int { return len(s) }

func (s sortServiceInstances) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s sortServiceInstances) Less(i, j int) bool {
	if s[i].ServiceName != s[j].ServiceName {
		return s[i].ServiceName < s[j].ServiceName
	}
	if s[i].ServicePlan != s[j].ServicePlan {
		return s[i].ServicePlan < s[j].ServicePlan
	}
	if s[i].OrgName != s[j].OrgName {
		return s[i].OrgName < s[j].OrgName
	}
	if s[i].SpaceName != s[j].SpaceName {
		return s[i].SpaceName < s[j].SpaceName
	}
	return s[i].InstanceGuid < s[j].InstanceGuid
}
