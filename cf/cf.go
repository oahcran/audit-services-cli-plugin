package cf

import (
	"code.cloudfoundry.org/cli/plugin"
	"fmt"
	"github.com/krujos/cfcurl"
	"strconv"
)

//counterfeiter:generate . CfClient

type CfClient interface {
	ListOrgs(cliConnection plugin.CliConnection) ([]Org, error)
	ListSpacesByOrg(cliConnection plugin.CliConnection, spacesUrl string) ([]Space, error)
	ListServices(cliConnection plugin.CliConnection) ([]Service, error)
	ListServicePlans(cliConnection plugin.CliConnection) ([]ServicePlan, error)
	GetServiceInstancesByUrl(cliConnection plugin.CliConnection, serviceInstancesUrl string) ([]ServiceInstance, error)
	GetServiceSummaryBySpace(cliConnection plugin.CliConnection, spaceUrl string) ([]ServiceSummaryBySpace, error)
}

// implementation of interface CfClient
type ClientHelper struct{}

type Org struct {
	Name      string
	Guid      string
	SpacesUrl string
}

type Space struct {
	Name                string
	Guid                string
	Url                 string
	OrgGuid             string
	ServiceInstancesUrl string
}

type Service struct {
	Name string
	Guid string
}

type ServicePlan struct {
	Name        string
	Guid        string
	ServiceGuid string
}

type ServiceInstance struct {
	Name            string
	Guid            string
	ServicePlanGuid string
	SpaceGuid       string
}

type ServiceSummaryBySpace struct {
	Name          string
	Guid          string
	BoundAppCount int
}

type process func(metadata map[string]interface{}, entity map[string]interface{}) interface{}

//Base method to process paged results from API calls
//Borrow from https://github.com/cdelashmutt-pivotal/service-use/blob/master/apihelper/apihelper.go#L32
func (helper *ClientHelper) processV2PagedResults(cliConnection plugin.CliConnection, url string, fn process) ([]interface{}, error) {

	theJSON, err := cfcurl.Curl(cliConnection, url)
	if nil != err {
		return nil, err
	}

	pages := int(theJSON["total_pages"].(float64))
	var objects []interface{}
	for i := 1; i <= pages; i++ {
		if 1 != i {
			theJSON, err = cfcurl.Curl(cliConnection, url+"?page="+strconv.Itoa(i))
		}
		for _, o := range theJSON["resources"].([]interface{}) {
			theObj := o.(map[string]interface{})
			entity := theObj["entity"].(map[string]interface{})
			metadata := theObj["metadata"].(map[string]interface{})
			objects = append(objects, fn(metadata, entity))
		}

	}

	return objects, nil
}

// list all the orgs in the foundation
func (helper *ClientHelper) ListOrgs(cliConnection plugin.CliConnection) ([]Org, error) {

	orgs, err := helper.processV2PagedResults(cliConnection, "/v2/organizations",
		func(metadata map[string]interface{}, entity map[string]interface{}) interface{} {
			return Org{
				Name:      entity["name"].(string),
				Guid:      metadata["guid"].(string),
				SpacesUrl: entity["spaces_url"].(string),
			}
		})

	if err != nil {
		return nil, err
	}
	retVal := make([]Org, len(orgs))
	for i := range orgs {
		retVal[i] = orgs[i].(Org)
	}

	return retVal, nil
}

// list all the spaces under the particular org
func (helper *ClientHelper) ListSpacesByOrg(cliConnection plugin.CliConnection, spacesUrl string) ([]Space, error) {

	spaces, err := helper.processV2PagedResults(cliConnection, spacesUrl,
		func(metadata map[string]interface{}, entity map[string]interface{}) interface{} {
			return Space{
				Name:                entity["name"].(string),
				Guid:                metadata["guid"].(string),
				Url:                 metadata["url"].(string),
				OrgGuid:             entity["organization_guid"].(string),
				ServiceInstancesUrl: entity["service_instances_url"].(string),
			}
		})

	if err != nil {
		return nil, err
	}
	retVal := make([]Space, len(spaces))
	for i := range spaces {
		retVal[i] = spaces[i].(Space)
	}

	return retVal, nil
}

// list all the services
func (helper *ClientHelper) ListServices(cliConnection plugin.CliConnection) ([]Service, error) {

	services, err := helper.processV2PagedResults(cliConnection, "/v2/services",
		func(metadata map[string]interface{}, entity map[string]interface{}) interface{} {
			return Service{
				Name: entity["label"].(string),
				Guid: metadata["guid"].(string),
			}
		})

	if err != nil {
		return nil, err
	}
	retVal := make([]Service, len(services))
	for i := range services {
		retVal[i] = services[i].(Service)
	}

	return retVal, nil
}

// list all the service plans
func (helper *ClientHelper) ListServicePlans(cliConnection plugin.CliConnection) ([]ServicePlan, error) {

	plans, err := helper.processV2PagedResults(cliConnection, "/v2/service_plans",
		func(metadata map[string]interface{}, entity map[string]interface{}) interface{} {
			return ServicePlan{
				Name:        entity["name"].(string),
				Guid:        metadata["guid"].(string),
				ServiceGuid: entity["service_guid"].(string),
			}
		})

	if err != nil {
		return nil, err
	}
	retVal := make([]ServicePlan, len(plans))
	for i := range plans {
		retVal[i] = plans[i].(ServicePlan)
	}

	return retVal, nil
}

// get all service instances from service_instances_url under space response
func (helper *ClientHelper) GetServiceInstancesByUrl(cliConnection plugin.CliConnection, serviceInstancesUrl string) ([]ServiceInstance, error) {

	instances, err := helper.processV2PagedResults(cliConnection, serviceInstancesUrl,
		func(metadata map[string]interface{}, entity map[string]interface{}) interface{} {
			return ServiceInstance{
				Name:            entity["name"].(string),
				Guid:            metadata["guid"].(string),
				ServicePlanGuid: entity["service_plan_guid"].(string),
				SpaceGuid:       entity["space_guid"].(string),
			}
		})

	if err != nil {
		return nil, err
	}

	retVal := make([]ServiceInstance, len(instances))
	for i := range instances {
		retVal[i] = instances[i].(ServiceInstance)
	}

	return retVal, nil
}

// get service summary for specific space to get bound_app_count numbers
func (helper *ClientHelper) GetServiceSummaryBySpace(cliConnection plugin.CliConnection, spaceUrl string) ([]ServiceSummaryBySpace, error) {

	theJSON, err := cfcurl.Curl(cliConnection, fmt.Sprintf("%s/summary", spaceUrl))

	if nil != err {
		return nil, err
	}

	var objects []ServiceSummaryBySpace

	for _, o := range theJSON["services"].([]interface{}) {

		theObj := o.(map[string]interface{})

		summary := ServiceSummaryBySpace{
			Name:          theObj["name"].(string),
			Guid:          theObj["guid"].(string),
			BoundAppCount: int(theObj["bound_app_count"].(float64)),
		}

		objects = append(objects, summary)

	}

	return objects, nil
}
