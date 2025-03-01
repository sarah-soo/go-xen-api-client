//
// This file is generated. To change the content of this file, please do not
// apply the change to this file because it will get overwritten. Instead,
// change xenapi.go and execute 'go generate'.
//

package xenapi

import (
	"fmt"
	"github.com/amfranz/go-xmlrpc-client"
	"reflect"
	"strconv"
	"time"
)

var _ = fmt.Errorf
var _ = xmlrpc.NewClient
var _ = reflect.TypeOf
var _ = strconv.Atoi
var _ = time.UTC

type VifOperations string

const (
	// Attempting to attach this VIF to a VM
	VifOperationsAttach VifOperations = "attach"
	// Attempting to hotplug this VIF
	VifOperationsPlug VifOperations = "plug"
	// Attempting to hot unplug this VIF
	VifOperationsUnplug VifOperations = "unplug"
)

type VifLockingMode string

const (
	// No specific configuration set - default network policy applies
	VifLockingModeNetworkDefault VifLockingMode = "network_default"
	// Only traffic to a specific MAC and a list of IPv4 or IPv6 addresses is permitted
	VifLockingModeLocked VifLockingMode = "locked"
	// All traffic is permitted
	VifLockingModeUnlocked VifLockingMode = "unlocked"
	// No traffic is permitted
	VifLockingModeDisabled VifLockingMode = "disabled"
)

type VifIpv4ConfigurationMode string

const (
	// Follow the default IPv4 configuration of the guest (this is guest-dependent)
	VifIpv4ConfigurationModeNone VifIpv4ConfigurationMode = "None"
	// Static IPv4 address configuration
	VifIpv4ConfigurationModeStatic VifIpv4ConfigurationMode = "Static"
)

type VifIpv6ConfigurationMode string

const (
	// Follow the default IPv6 configuration of the guest (this is guest-dependent)
	VifIpv6ConfigurationModeNone VifIpv6ConfigurationMode = "None"
	// Static IPv6 address configuration
	VifIpv6ConfigurationModeStatic VifIpv6ConfigurationMode = "Static"
)

type VIFRecord struct {
	// Unique identifier/object reference
	UUID string
	// list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	AllowedOperations []VifOperations
	// links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	CurrentOperations map[string]VifOperations
	// order in which VIF backends are created by xapi
	Device string
	// virtual network to which this vif is connected
	Network NetworkRef
	// virtual machine to which this vif is connected
	VM VMRef
	// ethernet MAC address of virtual interface, as exposed to guest
	MAC string
	// MTU in octets
	MTU int
	// additional configuration
	OtherConfig map[string]string
	// is the device currently attached (erased on reboot)
	CurrentlyAttached bool
	// error/success code associated with last attach-operation (erased on reboot)
	StatusCode int
	// error/success information associated with last attach-operation status (erased on reboot)
	StatusDetail string
	// Device runtime properties
	RuntimeProperties map[string]string
	// QoS algorithm to use
	QosAlgorithmType string
	// parameters for chosen QoS algorithm
	QosAlgorithmParams map[string]string
	// supported QoS algorithms for this VIF
	QosSupportedAlgorithms []string
	// metrics associated with this VIF
	Metrics VIFMetricsRef
	// true if the MAC was autogenerated; false indicates it was set manually
	MACAutogenerated bool
	// current locking mode of the VIF
	LockingMode VifLockingMode
	// A list of IPv4 addresses which can be used to filter traffic passing through this VIF
	Ipv4Allowed []string
	// A list of IPv6 addresses which can be used to filter traffic passing through this VIF
	Ipv6Allowed []string
	// Determines whether IPv4 addresses are configured on the VIF
	Ipv4ConfigurationMode VifIpv4ConfigurationMode
	// IPv4 addresses in CIDR format
	Ipv4Addresses []string
	// IPv4 gateway (the empty string means that no gateway is set)
	Ipv4Gateway string
	// Determines whether IPv6 addresses are configured on the VIF
	Ipv6ConfigurationMode VifIpv6ConfigurationMode
	// IPv6 addresses in CIDR format
	Ipv6Addresses []string
	// IPv6 gateway (the empty string means that no gateway is set)
	Ipv6Gateway string
}

type VIFRef string

// A virtual network interface
type VIFClass struct {
	client *Client
}

// GetAllRecords Return a map of VIF references to VIF records for all VIFs known to the system.
func (_class VIFClass) GetAllRecords(sessionID SessionRef) (_retval map[VIFRef]VIFRecord, _err error) {
	_method := "VIF.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVIFRefToVIFRecordMapToGo(_method + " -> ", _result.Value)
	return
}

// GetAll Return a list of all the VIFs known to the system.
func (_class VIFClass) GetAll(sessionID SessionRef) (_retval []VIFRef, _err error) {
	_method := "VIF.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVIFRefSetToGo(_method + " -> ", _result.Value)
	return
}

// ConfigureIpv6 Configure IPv6 settings for this virtual interface
func (_class VIFClass) ConfigureIpv6(sessionID SessionRef, self VIFRef, mode VifIpv6ConfigurationMode, address string, gateway string) (_err error) {
	_method := "VIF.configure_ipv6"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_modeArg, _err := convertEnumVifIpv6ConfigurationModeToXen(fmt.Sprintf("%s(%s)", _method, "mode"), mode)
	if _err != nil {
		return
	}
	_addressArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "address"), address)
	if _err != nil {
		return
	}
	_gatewayArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "gateway"), gateway)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _modeArg, _addressArg, _gatewayArg)
	return
}

// ConfigureIpv4 Configure IPv4 settings for this virtual interface
func (_class VIFClass) ConfigureIpv4(sessionID SessionRef, self VIFRef, mode VifIpv4ConfigurationMode, address string, gateway string) (_err error) {
	_method := "VIF.configure_ipv4"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_modeArg, _err := convertEnumVifIpv4ConfigurationModeToXen(fmt.Sprintf("%s(%s)", _method, "mode"), mode)
	if _err != nil {
		return
	}
	_addressArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "address"), address)
	if _err != nil {
		return
	}
	_gatewayArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "gateway"), gateway)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _modeArg, _addressArg, _gatewayArg)
	return
}

// RemoveIpv6Allowed Removes an IPv6 address from this VIF
func (_class VIFClass) RemoveIpv6Allowed(sessionID SessionRef, self VIFRef, value string) (_err error) {
	_method := "VIF.remove_ipv6_allowed"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// AddIpv6Allowed Associates an IPv6 address with this VIF
func (_class VIFClass) AddIpv6Allowed(sessionID SessionRef, self VIFRef, value string) (_err error) {
	_method := "VIF.add_ipv6_allowed"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetIpv6Allowed Set the IPv6 addresses to which traffic on this VIF can be restricted
func (_class VIFClass) SetIpv6Allowed(sessionID SessionRef, self VIFRef, value []string) (_err error) {
	_method := "VIF.set_ipv6_allowed"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringSetToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// RemoveIpv4Allowed Removes an IPv4 address from this VIF
func (_class VIFClass) RemoveIpv4Allowed(sessionID SessionRef, self VIFRef, value string) (_err error) {
	_method := "VIF.remove_ipv4_allowed"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// AddIpv4Allowed Associates an IPv4 address with this VIF
func (_class VIFClass) AddIpv4Allowed(sessionID SessionRef, self VIFRef, value string) (_err error) {
	_method := "VIF.add_ipv4_allowed"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetIpv4Allowed Set the IPv4 addresses to which traffic on this VIF can be restricted
func (_class VIFClass) SetIpv4Allowed(sessionID SessionRef, self VIFRef, value []string) (_err error) {
	_method := "VIF.set_ipv4_allowed"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringSetToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetLockingMode Set the locking mode for this VIF
func (_class VIFClass) SetLockingMode(sessionID SessionRef, self VIFRef, value VifLockingMode) (_err error) {
	_method := "VIF.set_locking_mode"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumVifLockingModeToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Move Move the specified VIF to the specified network, even while the VM is running
func (_class VIFClass) Move(sessionID SessionRef, self VIFRef, network NetworkRef) (_err error) {
	_method := "VIF.move"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_networkArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "network"), network)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _networkArg)
	return
}

// UnplugForce Forcibly unplug the specified VIF
func (_class VIFClass) UnplugForce(sessionID SessionRef, self VIFRef) (_err error) {
	_method := "VIF.unplug_force"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// Unplug Hot-unplug the specified VIF, dynamically unattaching it from the running VM
func (_class VIFClass) Unplug(sessionID SessionRef, self VIFRef) (_err error) {
	_method := "VIF.unplug"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// Plug Hotplug the specified VIF, dynamically attaching it to the running VM
func (_class VIFClass) Plug(sessionID SessionRef, self VIFRef) (_err error) {
	_method := "VIF.plug"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// RemoveFromQosAlgorithmParams Remove the given key and its corresponding value from the qos/algorithm_params field of the given VIF.  If the key is not in that Map, then do nothing.
func (_class VIFClass) RemoveFromQosAlgorithmParams(sessionID SessionRef, self VIFRef, key string) (_err error) {
	_method := "VIF.remove_from_qos_algorithm_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg)
	return
}

// AddToQosAlgorithmParams Add the given key-value pair to the qos/algorithm_params field of the given VIF.
func (_class VIFClass) AddToQosAlgorithmParams(sessionID SessionRef, self VIFRef, key string, value string) (_err error) {
	_method := "VIF.add_to_qos_algorithm_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg, _valueArg)
	return
}

// SetQosAlgorithmParams Set the qos/algorithm_params field of the given VIF.
func (_class VIFClass) SetQosAlgorithmParams(sessionID SessionRef, self VIFRef, value map[string]string) (_err error) {
	_method := "VIF.set_qos_algorithm_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetQosAlgorithmType Set the qos/algorithm_type field of the given VIF.
func (_class VIFClass) SetQosAlgorithmType(sessionID SessionRef, self VIFRef, value string) (_err error) {
	_method := "VIF.set_qos_algorithm_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// RemoveFromOtherConfig Remove the given key and its corresponding value from the other_config field of the given VIF.  If the key is not in that Map, then do nothing.
func (_class VIFClass) RemoveFromOtherConfig(sessionID SessionRef, self VIFRef, key string) (_err error) {
	_method := "VIF.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg)
	return
}

// AddToOtherConfig Add the given key-value pair to the other_config field of the given VIF.
func (_class VIFClass) AddToOtherConfig(sessionID SessionRef, self VIFRef, key string, value string) (_err error) {
	_method := "VIF.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg, _valueArg)
	return
}

// SetOtherConfig Set the other_config field of the given VIF.
func (_class VIFClass) SetOtherConfig(sessionID SessionRef, self VIFRef, value map[string]string) (_err error) {
	_method := "VIF.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// GetIpv6Gateway Get the ipv6_gateway field of the given VIF.
func (_class VIFClass) GetIpv6Gateway(sessionID SessionRef, self VIFRef) (_retval string, _err error) {
	_method := "VIF.get_ipv6_gateway"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// GetIpv6Addresses Get the ipv6_addresses field of the given VIF.
func (_class VIFClass) GetIpv6Addresses(sessionID SessionRef, self VIFRef) (_retval []string, _err error) {
	_method := "VIF.get_ipv6_addresses"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringSetToGo(_method + " -> ", _result.Value)
	return
}

// GetIpv6ConfigurationMode Get the ipv6_configuration_mode field of the given VIF.
func (_class VIFClass) GetIpv6ConfigurationMode(sessionID SessionRef, self VIFRef) (_retval VifIpv6ConfigurationMode, _err error) {
	_method := "VIF.get_ipv6_configuration_mode"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVifIpv6ConfigurationModeToGo(_method + " -> ", _result.Value)
	return
}

// GetIpv4Gateway Get the ipv4_gateway field of the given VIF.
func (_class VIFClass) GetIpv4Gateway(sessionID SessionRef, self VIFRef) (_retval string, _err error) {
	_method := "VIF.get_ipv4_gateway"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// GetIpv4Addresses Get the ipv4_addresses field of the given VIF.
func (_class VIFClass) GetIpv4Addresses(sessionID SessionRef, self VIFRef) (_retval []string, _err error) {
	_method := "VIF.get_ipv4_addresses"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringSetToGo(_method + " -> ", _result.Value)
	return
}

// GetIpv4ConfigurationMode Get the ipv4_configuration_mode field of the given VIF.
func (_class VIFClass) GetIpv4ConfigurationMode(sessionID SessionRef, self VIFRef) (_retval VifIpv4ConfigurationMode, _err error) {
	_method := "VIF.get_ipv4_configuration_mode"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVifIpv4ConfigurationModeToGo(_method + " -> ", _result.Value)
	return
}

// GetIpv6Allowed Get the ipv6_allowed field of the given VIF.
func (_class VIFClass) GetIpv6Allowed(sessionID SessionRef, self VIFRef) (_retval []string, _err error) {
	_method := "VIF.get_ipv6_allowed"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringSetToGo(_method + " -> ", _result.Value)
	return
}

// GetIpv4Allowed Get the ipv4_allowed field of the given VIF.
func (_class VIFClass) GetIpv4Allowed(sessionID SessionRef, self VIFRef) (_retval []string, _err error) {
	_method := "VIF.get_ipv4_allowed"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringSetToGo(_method + " -> ", _result.Value)
	return
}

// GetLockingMode Get the locking_mode field of the given VIF.
func (_class VIFClass) GetLockingMode(sessionID SessionRef, self VIFRef) (_retval VifLockingMode, _err error) {
	_method := "VIF.get_locking_mode"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVifLockingModeToGo(_method + " -> ", _result.Value)
	return
}

// GetMACAutogenerated Get the MAC_autogenerated field of the given VIF.
func (_class VIFClass) GetMACAutogenerated(sessionID SessionRef, self VIFRef) (_retval bool, _err error) {
	_method := "VIF.get_MAC_autogenerated"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result.Value)
	return
}

// GetMetrics Get the metrics field of the given VIF.
func (_class VIFClass) GetMetrics(sessionID SessionRef, self VIFRef) (_retval VIFMetricsRef, _err error) {
	_method := "VIF.get_metrics"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVIFMetricsRefToGo(_method + " -> ", _result.Value)
	return
}

// GetQosSupportedAlgorithms Get the qos/supported_algorithms field of the given VIF.
func (_class VIFClass) GetQosSupportedAlgorithms(sessionID SessionRef, self VIFRef) (_retval []string, _err error) {
	_method := "VIF.get_qos_supported_algorithms"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringSetToGo(_method + " -> ", _result.Value)
	return
}

// GetQosAlgorithmParams Get the qos/algorithm_params field of the given VIF.
func (_class VIFClass) GetQosAlgorithmParams(sessionID SessionRef, self VIFRef) (_retval map[string]string, _err error) {
	_method := "VIF.get_qos_algorithm_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result.Value)
	return
}

// GetQosAlgorithmType Get the qos/algorithm_type field of the given VIF.
func (_class VIFClass) GetQosAlgorithmType(sessionID SessionRef, self VIFRef) (_retval string, _err error) {
	_method := "VIF.get_qos_algorithm_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// GetRuntimeProperties Get the runtime_properties field of the given VIF.
func (_class VIFClass) GetRuntimeProperties(sessionID SessionRef, self VIFRef) (_retval map[string]string, _err error) {
	_method := "VIF.get_runtime_properties"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result.Value)
	return
}

// GetStatusDetail Get the status_detail field of the given VIF.
func (_class VIFClass) GetStatusDetail(sessionID SessionRef, self VIFRef) (_retval string, _err error) {
	_method := "VIF.get_status_detail"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// GetStatusCode Get the status_code field of the given VIF.
func (_class VIFClass) GetStatusCode(sessionID SessionRef, self VIFRef) (_retval int, _err error) {
	_method := "VIF.get_status_code"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}

// GetCurrentlyAttached Get the currently_attached field of the given VIF.
func (_class VIFClass) GetCurrentlyAttached(sessionID SessionRef, self VIFRef) (_retval bool, _err error) {
	_method := "VIF.get_currently_attached"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result.Value)
	return
}

// GetOtherConfig Get the other_config field of the given VIF.
func (_class VIFClass) GetOtherConfig(sessionID SessionRef, self VIFRef) (_retval map[string]string, _err error) {
	_method := "VIF.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result.Value)
	return
}

// GetMTU Get the MTU field of the given VIF.
func (_class VIFClass) GetMTU(sessionID SessionRef, self VIFRef) (_retval int, _err error) {
	_method := "VIF.get_MTU"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}

// GetMAC Get the MAC field of the given VIF.
func (_class VIFClass) GetMAC(sessionID SessionRef, self VIFRef) (_retval string, _err error) {
	_method := "VIF.get_MAC"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// GetVM Get the VM field of the given VIF.
func (_class VIFClass) GetVM(sessionID SessionRef, self VIFRef) (_retval VMRef, _err error) {
	_method := "VIF.get_VM"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}

// GetNetwork Get the network field of the given VIF.
func (_class VIFClass) GetNetwork(sessionID SessionRef, self VIFRef) (_retval NetworkRef, _err error) {
	_method := "VIF.get_network"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertNetworkRefToGo(_method + " -> ", _result.Value)
	return
}

// GetDevice Get the device field of the given VIF.
func (_class VIFClass) GetDevice(sessionID SessionRef, self VIFRef) (_retval string, _err error) {
	_method := "VIF.get_device"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// GetCurrentOperations Get the current_operations field of the given VIF.
func (_class VIFClass) GetCurrentOperations(sessionID SessionRef, self VIFRef) (_retval map[string]VifOperations, _err error) {
	_method := "VIF.get_current_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToEnumVifOperationsMapToGo(_method + " -> ", _result.Value)
	return
}

// GetAllowedOperations Get the allowed_operations field of the given VIF.
func (_class VIFClass) GetAllowedOperations(sessionID SessionRef, self VIFRef) (_retval []VifOperations, _err error) {
	_method := "VIF.get_allowed_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVifOperationsSetToGo(_method + " -> ", _result.Value)
	return
}

// GetUUID Get the uuid field of the given VIF.
func (_class VIFClass) GetUUID(sessionID SessionRef, self VIFRef) (_retval string, _err error) {
	_method := "VIF.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// Destroy Destroy the specified VIF instance.
func (_class VIFClass) Destroy(sessionID SessionRef, self VIFRef) (_err error) {
	_method := "VIF.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// Create Create a new VIF instance, and return its handle. The constructor args are: device*, network*, VM*, MAC*, MTU*, other_config*, currently_attached, qos_algorithm_type*, qos_algorithm_params*, locking_mode, ipv4_allowed, ipv6_allowed (* = non-optional).
func (_class VIFClass) Create(sessionID SessionRef, args VIFRecord) (_retval VIFRef, _err error) {
	_method := "VIF.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_argsArg, _err := convertVIFRecordToXen(fmt.Sprintf("%s(%s)", _method, "args"), args)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _argsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVIFRefToGo(_method + " -> ", _result.Value)
	return
}

// GetByUUID Get a reference to the VIF instance with the specified UUID.
func (_class VIFClass) GetByUUID(sessionID SessionRef, uuid string) (_retval VIFRef, _err error) {
	_method := "VIF.get_by_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_uuidArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "uuid"), uuid)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _uuidArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVIFRefToGo(_method + " -> ", _result.Value)
	return
}

// GetRecord Get a record containing the current state of the given VIF.
func (_class VIFClass) GetRecord(sessionID SessionRef, self VIFRef) (_retval VIFRecord, _err error) {
	_method := "VIF.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVIFRecordToGo(_method + " -> ", _result.Value)
	return
}
