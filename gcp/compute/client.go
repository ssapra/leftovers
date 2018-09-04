package compute

import (
	"fmt"
	"time"

	gcpcompute "google.golang.org/api/compute/v1"
	"google.golang.org/api/googleapi"
)

type client struct {
	project string
	logger  logger

	service               *gcpcompute.Service
	addresses             *gcpcompute.AddressesService
	globalAddresses       *gcpcompute.GlobalAddressesService
	backendServices       *gcpcompute.BackendServicesService
	disks                 *gcpcompute.DisksService
	globalHealthChecks    *gcpcompute.HealthChecksService
	httpHealthChecks      *gcpcompute.HttpHealthChecksService
	httpsHealthChecks     *gcpcompute.HttpsHealthChecksService
	images                *gcpcompute.ImagesService
	instanceTemplates     *gcpcompute.InstanceTemplatesService
	instances             *gcpcompute.InstancesService
	instanceGroups        *gcpcompute.InstanceGroupsService
	instanceGroupManagers *gcpcompute.InstanceGroupManagersService
	firewalls             *gcpcompute.FirewallsService
	forwardingRules       *gcpcompute.ForwardingRulesService
	globalForwardingRules *gcpcompute.GlobalForwardingRulesService
	subnetworks           *gcpcompute.SubnetworksService
	sslCertificates       *gcpcompute.SslCertificatesService
	networks              *gcpcompute.NetworksService
	targetHttpProxies     *gcpcompute.TargetHttpProxiesService
	targetHttpsProxies    *gcpcompute.TargetHttpsProxiesService
	targetPools           *gcpcompute.TargetPoolsService
	urlMaps               *gcpcompute.UrlMapsService
	regions               *gcpcompute.RegionsService
	zones                 *gcpcompute.ZonesService
}

func NewClient(project string, service *gcpcompute.Service, logger logger) client {
	return client{
		project:               project,
		logger:                logger,
		service:               service,
		addresses:             service.Addresses,
		globalAddresses:       service.GlobalAddresses,
		backendServices:       service.BackendServices,
		disks:                 service.Disks,
		globalHealthChecks:    service.HealthChecks,
		httpHealthChecks:      service.HttpHealthChecks,
		httpsHealthChecks:     service.HttpsHealthChecks,
		images:                service.Images,
		instanceTemplates:     service.InstanceTemplates,
		instances:             service.Instances,
		instanceGroups:        service.InstanceGroups,
		instanceGroupManagers: service.InstanceGroupManagers,
		firewalls:             service.Firewalls,
		forwardingRules:       service.ForwardingRules,
		globalForwardingRules: service.GlobalForwardingRules,
		sslCertificates:       service.SslCertificates,
		subnetworks:           service.Subnetworks,
		networks:              service.Networks,
		targetHttpProxies:     service.TargetHttpProxies,
		targetHttpsProxies:    service.TargetHttpsProxies,
		targetPools:           service.TargetPools,
		urlMaps:               service.UrlMaps,
		regions:               service.Regions,
		zones:                 service.Zones,
	}
}

func (c client) ListAddresses(region string) ([]*gcpcompute.Address, error) {
	list := []*gcpcompute.Address{}

	for token := ""; token != ""; {
		resp, err := c.addresses.List(c.project, region).PageToken(token).Do()
		if err != nil {
			return nil, err
		}

		list = append(list, resp.Items...)

		token = resp.NextPageToken
		if token != "" {
			time.Sleep(time.Second)
		}
	}

	return list, nil
}

func (c client) DeleteAddress(region, address string) error {
	return c.wait(c.addresses.Delete(c.project, region, address))
}

func (c client) ListGlobalAddresses() ([]*gcpcompute.Address, error) {
	list := []*gcpcompute.Address{}

	for token := ""; token != ""; {
		resp, err := c.globalAddresses.List(c.project).PageToken(token).Do()
		if err != nil {
			return nil, err
		}

		list = append(list, resp.Items...)

		token = resp.NextPageToken
		if token != "" {
			time.Sleep(time.Second)
		}
	}

	return list, nil
}

func (c client) DeleteGlobalAddress(address string) error {
	return c.wait(c.globalAddresses.Delete(c.project, address))
}

func (c client) ListBackendServices() ([]*gcpcompute.BackendService, error) {
	list := []*gcpcompute.BackendService{}

	for token := ""; token != ""; {
		resp, err := c.backendServices.List(c.project).PageToken(token).Do()
		if err != nil {
			return nil, err
		}

		list = append(list, resp.Items...)

		token = resp.NextPageToken
		if token != "" {
			time.Sleep(time.Second)
		}
	}

	return list, nil
}

func (c client) DeleteBackendService(backendService string) error {
	return c.wait(c.backendServices.Delete(c.project, backendService))
}

// ListDisks returns the full list of disks.
func (c client) ListDisks(zone string) ([]*gcpcompute.Disk, error) {
	list := []*gcpcompute.Disk{}

	for token := ""; token != ""; {
		resp, err := c.disks.List(c.project, zone).PageToken(token).Do()
		if err != nil {
			return nil, err
		}

		list = append(list, resp.Items...)

		token = resp.NextPageToken
		if token != "" {
			time.Sleep(time.Second)
		}
	}

	return list, nil
}

func (c client) DeleteDisk(zone, disk string) error {
	return c.wait(c.disks.Delete(c.project, zone, disk))
}

// ListImages returns the full list of images.
func (c client) ListImages() ([]*gcpcompute.Image, error) {
	list := []*gcpcompute.Image{}

	for token := ""; token != ""; {
		resp, err := c.images.List(c.project).PageToken(token).Do()
		if err != nil {
			return nil, err
		}

		list = append(list, resp.Items...)

		token = resp.NextPageToken
		if token != "" {
			time.Sleep(time.Second)
		}
	}

	return list, nil
}

func (c client) DeleteImage(image string) error {
	return c.wait(c.images.Delete(c.project, image))
}

func (c client) ListInstances(zone string) ([]*gcpcompute.Instance, error) {
	list := []*gcpcompute.Instance{}

	for token := ""; token != ""; {
		resp, err := c.instances.List(c.project, zone).PageToken(token).Do()
		if err != nil {
			return nil, err
		}

		list = append(list, resp.Items...)

		token = resp.NextPageToken
		if token != "" {
			time.Sleep(time.Second)
		}
	}

	return list, nil
}

func (c client) DeleteInstance(zone, instance string) error {
	return c.wait(c.instances.Delete(c.project, zone, instance))
}

func (c client) ListInstanceTemplates() ([]*gcpcompute.InstanceTemplate, error) {
	list := []*gcpcompute.InstanceTemplate{}

	for token := ""; token != ""; {
		resp, err := c.instanceTemplates.List(c.project).PageToken(token).Do()
		if err != nil {
			return nil, err
		}

		list = append(list, resp.Items...)

		token = resp.NextPageToken
		if token != "" {
			time.Sleep(time.Second)
		}
	}

	return list, nil
}

func (c client) DeleteInstanceTemplate(instanceTemplate string) error {
	return c.wait(c.instanceTemplates.Delete(c.project, instanceTemplate))
}

func (c client) ListInstanceGroups(zone string) ([]*gcpcompute.InstanceGroup, error) {
	list := []*gcpcompute.InstanceGroup{}

	for token := ""; token != ""; {
		resp, err := c.instanceGroups.List(c.project, zone).PageToken(token).Do()
		if err != nil {
			return nil, err
		}

		list = append(list, resp.Items...)

		token = resp.NextPageToken
		if token != "" {
			time.Sleep(time.Second)
		}
	}

	return list, nil
}

func (c client) DeleteInstanceGroup(zone, instanceGroup string) error {
	return c.wait(c.instanceGroups.Delete(c.project, zone, instanceGroup))
}

func (c client) ListInstanceGroupManagers(zone string) ([]*gcpcompute.InstanceGroupManager, error) {
	list := []*gcpcompute.InstanceGroupManager{}

	for token := ""; token != ""; {
		resp, err := c.instanceGroupManagers.List(c.project, zone).PageToken(token).Do()
		if err != nil {
			return nil, err
		}

		list = append(list, resp.Items...)

		token = resp.NextPageToken
		if token != "" {
			time.Sleep(time.Second)
		}
	}

	return list, nil
}

func (c client) DeleteInstanceGroupManager(zone, instanceGroupManager string) error {
	return c.wait(c.instanceGroupManagers.Delete(c.project, zone, instanceGroupManager))
}

func (c client) ListGlobalHealthChecks() ([]*gcpcompute.HealthCheck, error) {
	list := []*gcpcompute.HealthCheck{}

	for token := ""; token != ""; {
		resp, err := c.globalHealthChecks.List(c.project).PageToken(token).Do()
		if err != nil {
			return nil, err
		}

		list = append(list, resp.Items...)

		token = resp.NextPageToken
		if token != "" {
			time.Sleep(time.Second)
		}
	}

	return list, nil
}

func (c client) DeleteGlobalHealthCheck(globalHealthCheck string) error {
	return c.wait(c.globalHealthChecks.Delete(c.project, globalHealthCheck))
}

func (c client) ListHttpHealthChecks() ([]*gcpcompute.HttpHealthCheck, error) {
	list := []*gcpcompute.HttpHealthCheck{}

	for token := ""; token != ""; {
		resp, err := c.httpHealthChecks.List(c.project).PageToken(token).Do()
		if err != nil {
			return nil, err
		}

		list = append(list, resp.Items...)

		token = resp.NextPageToken
		if token != "" {
			time.Sleep(time.Second)
		}
	}

	return list, nil
}

func (c client) DeleteHttpHealthCheck(httpHealthCheck string) error {
	return c.wait(c.httpHealthChecks.Delete(c.project, httpHealthCheck))
}

func (c client) ListHttpsHealthChecks() ([]*gcpcompute.HttpsHealthCheck, error) {
	list := []*gcpcompute.HttpsHealthCheck{}

	for token := ""; token != ""; {
		resp, err := c.httpsHealthChecks.List(c.project).PageToken(token).Do()
		if err != nil {
			return nil, err
		}

		list = append(list, resp.Items...)

		token = resp.NextPageToken
		if token != "" {
			time.Sleep(time.Second)
		}
	}

	return list, nil
}

func (c client) DeleteHttpsHealthCheck(httpsHealthCheck string) error {
	return c.wait(c.httpsHealthChecks.Delete(c.project, httpsHealthCheck))
}

func (c client) ListFirewalls() ([]*gcpcompute.Firewall, error) {
	list := []*gcpcompute.Firewall{}

	for token := ""; token != ""; {
		resp, err := c.firewalls.List(c.project).PageToken(token).Do()
		if err != nil {
			return nil, err
		}

		list = append(list, resp.Items...)

		token = resp.NextPageToken
		if token != "" {
			time.Sleep(time.Second)
		}
	}

	return list, nil
}

func (c client) DeleteFirewall(firewall string) error {
	return c.wait(c.firewalls.Delete(c.project, firewall))
}

func (c client) ListGlobalForwardingRules() ([]*gcpcompute.ForwardingRule, error) {
	list := []*gcpcompute.ForwardingRule{}

	for token := ""; token != ""; {
		resp, err := c.globalForwardingRules.List(c.project).PageToken(token).Do()
		if err != nil {
			return nil, err
		}

		list = append(list, resp.Items...)

		token = resp.NextPageToken
		if token != "" {
			time.Sleep(time.Second)
		}
	}

	return list, nil
}

func (c client) DeleteGlobalForwardingRule(globalForwardingRule string) error {
	return c.wait(c.globalForwardingRules.Delete(c.project, globalForwardingRule))
}

func (c client) ListForwardingRules(region string) ([]*gcpcompute.ForwardingRule, error) {
	list := []*gcpcompute.ForwardingRule{}

	for token := ""; token != ""; {
		resp, err := c.forwardingRules.List(c.project, region).PageToken(token).Do()
		if err != nil {
			return nil, err
		}

		list = append(list, resp.Items...)

		token = resp.NextPageToken
		if token != "" {
			time.Sleep(time.Second)
		}
	}

	return list, nil
}

func (c client) DeleteForwardingRule(region, forwardingRule string) error {
	return c.wait(c.forwardingRules.Delete(c.project, region, forwardingRule))
}

func (c client) ListNetworks() ([]*gcpcompute.Network, error) {
	list := []*gcpcompute.Network{}

	for token := ""; token != ""; {
		resp, err := c.networks.List(c.project).PageToken(token).Do()
		if err != nil {
			return nil, err
		}

		list = append(list, resp.Items...)

		token = resp.NextPageToken
		if token != "" {
			time.Sleep(time.Second)
		}
	}

	return list, nil
}

func (c client) DeleteNetwork(network string) error {
	return c.wait(c.networks.Delete(c.project, network))
}

func (c client) ListSubnetworks(region string) ([]*gcpcompute.Subnetwork, error) {
	list := []*gcpcompute.Subnetwork{}

	for token := ""; token != ""; {
		resp, err := c.subnetworks.List(c.project, region).PageToken(token).Do()
		if err != nil {
			return nil, err
		}

		list = append(list, resp.Items...)

		token = resp.NextPageToken
		if token != "" {
			time.Sleep(time.Second)
		}
	}

	return list, nil
}

func (c client) DeleteSubnetwork(region, subnetwork string) error {
	return c.wait(c.subnetworks.Delete(c.project, region, subnetwork))
}

func (c client) ListSslCertificates() (*gcpcompute.SslCertificateList, error) {
	return c.sslCertificates.List(c.project).Do()
}

func (c client) DeleteSslCertificate(certificate string) error {
	return c.wait(c.sslCertificates.Delete(c.project, certificate))
}

func (c client) ListTargetHttpProxies() (*gcpcompute.TargetHttpProxyList, error) {
	return c.targetHttpProxies.List(c.project).Do()
}

func (c client) DeleteTargetHttpProxy(targetHttpProxy string) error {
	return c.wait(c.targetHttpProxies.Delete(c.project, targetHttpProxy))
}

func (c client) ListTargetHttpsProxies() (*gcpcompute.TargetHttpsProxyList, error) {
	return c.targetHttpsProxies.List(c.project).Do()
}

func (c client) DeleteTargetHttpsProxy(targetHttpsProxy string) error {
	return c.wait(c.targetHttpsProxies.Delete(c.project, targetHttpsProxy))
}

func (c client) ListTargetPools(region string) (*gcpcompute.TargetPoolList, error) {
	return c.targetPools.List(c.project, region).Do()
}

func (c client) DeleteTargetPool(region string, targetPool string) error {
	return c.wait(c.targetPools.Delete(c.project, region, targetPool))
}

func (c client) ListUrlMaps() (*gcpcompute.UrlMapList, error) {
	return c.urlMaps.List(c.project).Do()
}

func (c client) DeleteUrlMap(urlMap string) error {
	return c.wait(c.urlMaps.Delete(c.project, urlMap))
}

func (c client) ListRegions() (map[string]string, error) {
	regions := map[string]string{}

	list, err := c.regions.List(c.project).Do()
	if err != nil {
		return regions, fmt.Errorf("List Regions: %s", err)
	}

	for _, r := range list.Items {
		regions[r.SelfLink] = r.Name
	}
	return regions, nil
}

func (c client) ListZones() (map[string]string, error) {
	zones := map[string]string{}

	list, err := c.zones.List(c.project).Do()
	if err != nil {
		return zones, fmt.Errorf("List Zones: %s", err)
	}

	for _, z := range list.Items {
		zones[z.SelfLink] = z.Name
	}
	return zones, nil
}

type request interface {
	Do(...googleapi.CallOption) (*gcpcompute.Operation, error)
}

func (c client) wait(request request) error {
	op, err := request.Do()
	if err != nil {
		if gerr, ok := err.(*googleapi.Error); ok {
			if gerr.Code == 404 {
				return nil
			}
		}
		return fmt.Errorf("Do request: %s", err)
	}

	waiter := NewOperationWaiter(op, c.service, c.project, c.logger)

	return waiter.Wait()
}
