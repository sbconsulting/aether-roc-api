// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package server

import (
	"fmt"
	externalRef0Svr "github.com/onosproject/aether-roc-api/pkg/aether_2_0_0/server"
	externalRef0 "github.com/onosproject/aether-roc-api/pkg/aether_2_0_0/types"
	externalRef2Svr "github.com/onosproject/aether-roc-api/pkg/aether_4_0_0/server"
	externalRef2 "github.com/onosproject/aether-roc-api/pkg/aether_4_0_0/types"
	"github.com/onosproject/aether-roc-api/pkg/toplevel/types"
	"github.com/openconfig/gnmi/proto/gnmi"
	"regexp"
)

var re = regexp.MustCompile(`[0-9a-z\-\._]+`)

func encodeToGnmiPatchBody(jsonObj *types.PatchBody) ([]*gnmi.Update, []*gnmi.Path, *string, *string, *string, string, error) {
	updates := make([]*gnmi.Update, 0)
	deletes := make([]*gnmi.Path, 0)
	var ext100Name *string
	var ext101Version *string
	var ext102Type *string

	if jsonObj.Extensions != nil {
		ext100Name = jsonObj.Extensions.ChangeName100
		ext101Version = jsonObj.Extensions.ModelVersion101
		ext102Type = jsonObj.Extensions.ModelType102
	}

	if !re.MatchString(jsonObj.DefaultTarget) {
		return nil, nil, ext100Name, ext101Version, ext102Type, jsonObj.DefaultTarget, fmt.Errorf("default-target cannot be blank")
	}

	gnmiUpdates, err := encodeToGnmiElements(jsonObj.Updates, jsonObj.DefaultTarget, false)
	if err != nil {
		return nil, nil, ext100Name, ext101Version, ext102Type, jsonObj.DefaultTarget, fmt.Errorf("encodeToGnmiElements() %s", err.Error())
	}
	updates = append(updates, gnmiUpdates...)

	gnmiDeletes, err := encodeToGnmiElements(jsonObj.Deletes, jsonObj.DefaultTarget, true)
	if err != nil {
		return nil, nil, ext100Name, ext101Version, ext102Type, jsonObj.DefaultTarget, fmt.Errorf("encodeToGnmiElements() %s", err.Error())
	}
	for _, gd := range gnmiDeletes {
		deletes = append(deletes, gd.Path)
	}

	return updates, deletes, ext100Name, ext101Version, ext102Type, jsonObj.DefaultTarget, nil
}

func encodeToGnmiElements(elements *types.Elements, target string, forDelete bool) ([]*gnmi.Update, error) {
	if elements == nil {
		return nil, nil
	}
	updates := make([]*gnmi.Update, 0)

	// Aether 2.x

	if elements.ConnectivityServices200 != nil {
		connectivityServiceUpdates, err := externalRef0Svr.EncodeToGnmiConnectivityServices(
			elements.ConnectivityServices200, false, forDelete, externalRef0.Target(target),
			"/connectivity-services")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiConnectivityService() %s", err)
		}
		updates = append(updates, connectivityServiceUpdates...)
	}

	if elements.Enterprises200 != nil {
		enterpriseUpdates, err := externalRef0Svr.EncodeToGnmiEnterprises(
			elements.Enterprises200, false, forDelete, externalRef0.Target(target),
			"/enterprises")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiEnterprise() %s", err)
		}
		updates = append(updates, enterpriseUpdates...)
	}

	// Aether 4.x

	if elements.ConnectivityService400 != nil {
		connectivityServiceUpdates, err := externalRef2Svr.EncodeToGnmiConnectivityService(
			elements.ConnectivityService400, false, forDelete, externalRef2.Target(target),
			"/connectivity-service")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiConnectivityService() %s", err)
		}
		updates = append(updates, connectivityServiceUpdates...)
	}

	if elements.Enterprise400 != nil {
		enterpriseUpdates, err := externalRef2Svr.EncodeToGnmiEnterprise(
			elements.Enterprise400, false, forDelete, externalRef2.Target(target),
			"/enterprise")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiEnterprise() %s", err)
		}
		updates = append(updates, enterpriseUpdates...)
	}

	if elements.Application400 != nil {
		applicationUpdates, err := externalRef2Svr.EncodeToGnmiApplication(
			elements.Application400, false, forDelete, externalRef2.Target(target),
			"/application")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiApplication() %s", err)
		}
		updates = append(updates, applicationUpdates...)
	}

	if elements.DeviceGroup400 != nil {
		deviceGroupUpdates, err := externalRef2Svr.EncodeToGnmiDeviceGroup(
			elements.DeviceGroup400, false, forDelete, externalRef2.Target(target),
			"/device-group")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiDeviceGroup() %s", err)
		}
		updates = append(updates, deviceGroupUpdates...)
	}

	if elements.IpDomain400 != nil {
		ipDomainUpdates, err := externalRef2Svr.EncodeToGnmiIpDomain(
			elements.IpDomain400, false, forDelete, externalRef2.Target(target),
			"/ip-domain")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiIpDomain() %s", err)
		}
		updates = append(updates, ipDomainUpdates...)
	}

	if elements.Site400 != nil {
		siteUpdates, err := externalRef2Svr.EncodeToGnmiSite(
			elements.Site400, false, forDelete, externalRef2.Target(target),
			"/site")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiSite() %s", err)
		}
		updates = append(updates, siteUpdates...)
	}

	if elements.Template400 != nil {
		templateUpdates, err := externalRef2Svr.EncodeToGnmiTemplate(
			elements.Template400, false, forDelete, externalRef2.Target(target),
			"/template")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiTemplate() %s", err)
		}
		updates = append(updates, templateUpdates...)
	}

	if elements.TrafficClass400 != nil {
		trafficClassUpdates, err := externalRef2Svr.EncodeToGnmiTrafficClass(
			elements.TrafficClass400, false, forDelete, externalRef2.Target(target),
			"/traffic-class")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiTrafficClass() %s", err)
		}
		updates = append(updates, trafficClassUpdates...)
	}

	if elements.Upf400 != nil {
		upfUpdates, err := externalRef2Svr.EncodeToGnmiUpf(
			elements.Upf400, false, forDelete, externalRef2.Target(target),
			"/upf")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiUpf() %s", err)
		}
		updates = append(updates, upfUpdates...)
	}

	if elements.Vcs400 != nil {
		vcsUpdates, err := externalRef2Svr.EncodeToGnmiVcs(
			elements.Vcs400, false, forDelete, externalRef2.Target(target),
			"/vcs")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiVcs() %s", err)
		}
		updates = append(updates, vcsUpdates...)
	}

	return updates, nil
}
