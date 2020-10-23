package outscale

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/antihax/optional"
	oscgo "github.com/marinsalinas/osc-sdk-go"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func attrLBListenerRules() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"listener_rules": {
			Type:     schema.TypeList,
			Computed: true,
			ForceNew: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"action": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"host_name_pattern": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"listener_rule_name": {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},
					"path_pattern": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"priority": {
						Type:     schema.TypeInt,
						Computed: true,
					},
					"vm_ids": {
						Type:     schema.TypeSet,
						Computed: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},
					"request_id": {
						Type:     schema.TypeString,
						Computed: true,
					},
				},
			},
		},
	}
}

func dataSourceOutscaleOAPILoadBalancerLDRules() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceOutscaleOAPILoadBalancerLDRulesRead,
		Schema: getDataSourceSchemas(attrLBListenerRules()),
	}
}

func dataSourceOutscaleOAPILoadBalancerLDRulesRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*OutscaleClient).OSCAPI

	lrNamei, nameOk := d.GetOk("listener_rule_name")
	filters, filtersOk := d.GetOk("filter")
	filter := &oscgo.FiltersListenerRule{}

	if !nameOk && !filtersOk {
		return fmt.Errorf("listener_rule_name must be assigned")
	}

	if filtersOk {
		set := filters.(*schema.Set)

		if set.Len() < 1 {
			return fmt.Errorf("filter can't be empty")
		}
		for _, v := range set.List() {
			m := v.(map[string]interface{})
			filterValues := make([]string, 0)
			for _, e := range m["values"].([]interface{}) {
				filterValues = append(filterValues, e.(string))
			}

			switch name := m["name"].(string); name {
			case "listener_rule_name":
				filter.ListenerRuleNames = &filterValues
			default:
				filter.ListenerRuleNames = &filterValues
				log.Printf("[Debug] Unknown Filter Name: %s. default to 'load_balancer_name'", name)
			}
		}
	} else {
		filter = &oscgo.FiltersListenerRule{
			ListenerRuleNames: &[]string{lrNamei.(string)},
		}
	}

	req := oscgo.ReadListenerRulesRequest{
		Filters: filter,
	}

	describeElbOpts := &oscgo.ReadListenerRulesOpts{
		ReadListenerRulesRequest: optional.NewInterface(req),
	}

	var resp oscgo.ReadListenerRulesResponse
	var err error
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		resp, _, err = conn.ListenerApi.ReadListenerRules(
			context.Background(),
			describeElbOpts)
		if err != nil {
			if strings.Contains(fmt.Sprint(err), "Throttling:") {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})

	if err != nil {
		return err
	}

	lrs := *resp.ListenerRules
	lrs_len := len(lrs)
	if lrs_len < 1 {
		return fmt.Errorf("can't find listener rule")
	}

	lrs_ret := make([]map[string]interface{}, lrs_len)
	for k, lr := range lrs {
		l := make(map[string]interface{})
		if lr.Action != nil {
			l["action"] = lr.Action
		}
		if lr.HostNamePattern != nil {
			l["host_name_pattern"] = lr.HostNamePattern
		}
		if lr.ListenerRuleName != nil {
			l["listener_rule_name"] = lr.ListenerRuleName
		}
		if lr.PathPattern != nil {
			l["path_pattern"] = lr.PathPattern
		}

		if lr.Priority != nil {
			l["priority"] = lr.Priority
		} else {
			fmt.Errorf("Malformated listener rule")
		}

		if lr.VmIds != nil {
			l["vm_ids"] = flattenStringList(lr.VmIds)
		} else {
			fmt.Errorf("Malformated listener rule")
		}
		lrs_ret[k] = l
	}

	d.Set("listener_rules", lrs_ret)
	d.Set("request_id", resp.ResponseContext.RequestId)
	d.SetId(resource.UniqueId())

	return nil
}
