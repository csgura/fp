package gendebug

import (
	"github.com/csgura/fp/mshow"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.Derive(recursive=true)
var _ mshow.Derives[mshow.Show[InitiatingMessageValue]]

type InitiatingMessageValue struct {
	Present                               int
	AMFConfigurationUpdate                *AMFConfigurationUpdate                `aper:"valueExt,referenceFieldValue:0,criticality:reject"`
	HandoverCancel                        *HandoverCancel                        `aper:"valueExt,referenceFieldValue:10,criticality:reject"`
	HandoverRequired                      *HandoverRequired                      `aper:"valueExt,referenceFieldValue:12,criticality:reject"`
	HandoverRequest                       *HandoverRequest                       `aper:"valueExt,referenceFieldValue:13,criticality:reject"`
	InitialContextSetupRequest            *InitialContextSetupRequest            `aper:"valueExt,referenceFieldValue:14,criticality:reject"`
	NGReset                               *NGReset                               `aper:"valueExt,referenceFieldValue:20,criticality:reject"`
	NGSetupRequest                        *NGSetupRequest                        `aper:"valueExt,referenceFieldValue:21,criticality:reject"`
	PathSwitchRequest                     *PathSwitchRequest                     `aper:"valueExt,referenceFieldValue:25,criticality:reject"`
	PDUSessionResourceModifyRequest       *PDUSessionResourceModifyRequest       `aper:"valueExt,referenceFieldValue:26,criticality:reject"`
	PDUSessionResourceModifyIndication    *PDUSessionResourceModifyIndication    `aper:"valueExt,referenceFieldValue:27,criticality:reject"`
	PDUSessionResourceReleaseCommand      *PDUSessionResourceReleaseCommand      `aper:"valueExt,referenceFieldValue:28,criticality:reject"`
	PDUSessionResourceSetupRequest        *PDUSessionResourceSetupRequest        `aper:"valueExt,referenceFieldValue:29,criticality:reject"`
	PWSCancelRequest                      *PWSCancelRequest                      `aper:"valueExt,referenceFieldValue:32,criticality:reject"`
	RANConfigurationUpdate                *RANConfigurationUpdate                `aper:"valueExt,referenceFieldValue:35,criticality:reject"`
	UEContextModificationRequest          *UEContextModificationRequest          `aper:"valueExt,referenceFieldValue:40,criticality:reject"`
	UEContextReleaseCommand               *UEContextReleaseCommand               `aper:"valueExt,referenceFieldValue:41,criticality:reject"`
	UEContextResumeRequest                *UEContextResumeRequest                `aper:"valueExt,referenceFieldValue:58,criticality:reject"`
	UEContextSuspendRequest               *UEContextSuspendRequest               `aper:"valueExt,referenceFieldValue:59,criticality:reject"`
	UERadioCapabilityCheckRequest         *UERadioCapabilityCheckRequest         `aper:"valueExt,referenceFieldValue:43,criticality:reject"`
	UERadioCapabilityIDMappingRequest     *UERadioCapabilityIDMappingRequest     `aper:"valueExt,referenceFieldValue:60,criticality:reject"`
	WriteReplaceWarningRequest            *WriteReplaceWarningRequest            `aper:"valueExt,referenceFieldValue:51,criticality:reject"`
	AMFCPRelocationIndication             *AMFCPRelocationIndication             `aper:"valueExt,referenceFieldValue:64,criticality:reject"`
	AMFStatusIndication                   *AMFStatusIndication                   `aper:"valueExt,referenceFieldValue:1,criticality:ignore"`
	CellTrafficTrace                      *CellTrafficTrace                      `aper:"valueExt,referenceFieldValue:2,criticality:ignore"`
	ConnectionEstablishmentIndication     *ConnectionEstablishmentIndication     `aper:"valueExt,referenceFieldValue:65,criticality:reject"`
	DeactivateTrace                       *DeactivateTrace                       `aper:"valueExt,referenceFieldValue:3,criticality:ignore"`
	DownlinkNASTransport                  *DownlinkNASTransport                  `aper:"valueExt,referenceFieldValue:4,criticality:ignore"`
	DownlinkNonUEAssociatedNRPPaTransport *DownlinkNonUEAssociatedNRPPaTransport `aper:"valueExt,referenceFieldValue:5,criticality:ignore"`
	DownlinkRANConfigurationTransfer      *DownlinkRANConfigurationTransfer      `aper:"valueExt,referenceFieldValue:6,criticality:ignore"`
	DownlinkRANEarlyStatusTransfer        *DownlinkRANEarlyStatusTransfer        `aper:"valueExt,referenceFieldValue:63,criticality:ignore"`
	DownlinkRANStatusTransfer             *DownlinkRANStatusTransfer             `aper:"valueExt,referenceFieldValue:7,criticality:ignore"`
	DownlinkRIMInformationTransfer        *DownlinkRIMInformationTransfer        `aper:"valueExt,referenceFieldValue:54,criticality:ignore"`
	DownlinkUEAssociatedNRPPaTransport    *DownlinkUEAssociatedNRPPaTransport    `aper:"valueExt,referenceFieldValue:8,criticality:ignore"`
	ErrorIndication                       *ErrorIndication                       `aper:"valueExt,referenceFieldValue:9,criticality:ignore"`
	HandoverNotify                        *HandoverNotify                        `aper:"valueExt,referenceFieldValue:11,criticality:ignore"`
	HandoverSuccess                       *HandoverSuccess                       `aper:"valueExt,referenceFieldValue:61,criticality:ignore"`
	InitialUEMessage                      *InitialUEMessage                      `aper:"valueExt,referenceFieldValue:15,criticality:ignore"`
	LocationReport                        *LocationReport                        `aper:"valueExt,referenceFieldValue:18,criticality:ignore"`
	LocationReportingControl              *LocationReportingControl              `aper:"valueExt,referenceFieldValue:16,criticality:ignore"`
	LocationReportingFailureIndication    *LocationReportingFailureIndication    `aper:"valueExt,referenceFieldValue:17,criticality:ignore"`
	NASNonDeliveryIndication              *NASNonDeliveryIndication              `aper:"valueExt,referenceFieldValue:19,criticality:ignore"`
	OverloadStart                         *OverloadStart                         `aper:"valueExt,referenceFieldValue:22,criticality:ignore"`
	OverloadStop                          *OverloadStop                          `aper:"valueExt,referenceFieldValue:23,criticality:reject"`
	Paging                                *Paging                                `aper:"valueExt,referenceFieldValue:24,criticality:ignore"`
	PDUSessionResourceNotify              *PDUSessionResourceNotify              `aper:"valueExt,referenceFieldValue:30,criticality:ignore"`
	// PrivateMessage                        *PrivateMessage                        `aper:"valueExt,referenceFieldValue:31,criticality:ignore"`
	// PWSFailureIndication                  *PWSFailureIndication                  `aper:"valueExt,referenceFieldValue:33,criticality:ignore"`
	// PWSRestartIndication                  *PWSRestartIndication                  `aper:"valueExt,referenceFieldValue:34,criticality:ignore"`
	// RANCPRelocationIndication             *RANCPRelocationIndication             `aper:"valueExt,referenceFieldValue:57,criticality:reject"`
	// RerouteNASRequest                     *RerouteNASRequest                     `aper:"valueExt,referenceFieldValue:36,criticality:reject"`
	// RetrieveUEInformation                 *RetrieveUEInformation                 `aper:"valueExt,referenceFieldValue:55,criticality:reject"`
	// RRCInactiveTransitionReport           *RRCInactiveTransitionReport           `aper:"valueExt,referenceFieldValue:37,criticality:ignore"`
	// SecondaryRATDataUsageReport           *SecondaryRATDataUsageReport           `aper:"valueExt,referenceFieldValue:52,criticality:ignore"`
	// TraceFailureIndication                *TraceFailureIndication                `aper:"valueExt,referenceFieldValue:38,criticality:ignore"`
	// TraceStart                            *TraceStart                            `aper:"valueExt,referenceFieldValue:39,criticality:ignore"`
	// UEContextReleaseRequest               *UEContextReleaseRequest               `aper:"valueExt,referenceFieldValue:42,criticality:ignore"`
	// UEInformationTransfer                 *UEInformationTransfer                 `aper:"valueExt,referenceFieldValue:56,criticality:reject"`
	// UERadioCapabilityInfoIndication       *UERadioCapabilityInfoIndication       `aper:"valueExt,referenceFieldValue:44,criticality:ignore"`
	// UETNLABindingReleaseRequest           *UETNLABindingReleaseRequest           `aper:"valueExt,referenceFieldValue:45,criticality:ignore"`
	// UplinkNASTransport                    *UplinkNASTransport                    `aper:"valueExt,referenceFieldValue:46,criticality:ignore"`
	// UplinkNonUEAssociatedNRPPaTransport   *UplinkNonUEAssociatedNRPPaTransport   `aper:"valueExt,referenceFieldValue:47,criticality:ignore"`
	// UplinkRANConfigurationTransfer        *UplinkRANConfigurationTransfer        `aper:"valueExt,referenceFieldValue:48,criticality:ignore"`
	// UplinkRANEarlyStatusTransfer          *UplinkRANEarlyStatusTransfer          `aper:"valueExt,referenceFieldValue:62,criticality:reject"`
	// UplinkRANStatusTransfer               *UplinkRANStatusTransfer               `aper:"valueExt,referenceFieldValue:49,criticality:ignore"`
	// UplinkRIMInformationTransfer          *UplinkRIMInformationTransfer          `aper:"valueExt,referenceFieldValue:53,criticality:ignore"`
	// UplinkUEAssociatedNRPPaTransport      *UplinkUEAssociatedNRPPaTransport      `aper:"valueExt,referenceFieldValue:50,criticality:ignore"`
}

type AMFConfigurationUpdate struct {
}
type HandoverCancel struct {
}
type HandoverRequired struct {
}
type HandoverRequest struct {
}
type InitialContextSetupRequest struct {
}
type NGReset struct {
}
type NGSetupRequest struct {
}
type PathSwitchRequest struct {
}
type PDUSessionResourceModifyRequest struct {
}
type PDUSessionResourceModifyIndication struct {
}
type PDUSessionResourceReleaseCommand struct {
}
type PDUSessionResourceSetupRequest struct {
}
type PWSCancelRequest struct {
}
type RANConfigurationUpdate struct {
}
type UEContextModificationRequest struct {
}
type UEContextReleaseCommand struct {
}
type UEContextResumeRequest struct {
}
type UEContextSuspendRequest struct {
}
type UERadioCapabilityCheckRequest struct {
}
type UERadioCapabilityIDMappingRequest struct {
}
type WriteReplaceWarningRequest struct {
}
type AMFCPRelocationIndication struct {
}
type AMFStatusIndication struct {
}
type CellTrafficTrace struct {
}
type ConnectionEstablishmentIndication struct {
}
type DeactivateTrace struct {
}
type DownlinkNASTransport struct {
}
type DownlinkNonUEAssociatedNRPPaTransport struct {
}
type DownlinkRANConfigurationTransfer struct {
}
type DownlinkRANEarlyStatusTransfer struct {
}
type DownlinkRANStatusTransfer struct {
}
type DownlinkRIMInformationTransfer struct {
}
type DownlinkUEAssociatedNRPPaTransport struct {
}
type ErrorIndication struct {
}
type HandoverNotify struct {
}
type HandoverSuccess struct {
}
type InitialUEMessage struct {
}
type LocationReport struct {
}
type LocationReportingControl struct {
}
type LocationReportingFailureIndication struct {
}
type NASNonDeliveryIndication struct {
}
type OverloadStart struct {
}
type OverloadStop struct {
}
type Paging struct {
}
type PDUSessionResourceNotify struct {
}
type PrivateMessage struct {
}
type PWSFailureIndication struct {
}
type PWSRestartIndication struct {
}
type RANCPRelocationIndication struct {
}
type RerouteNASRequest struct {
}
type RetrieveUEInformation struct {
}
type RRCInactiveTransitionReport struct {
}
type SecondaryRATDataUsageReport struct {
}
type TraceFailureIndication struct {
}
type TraceStart struct {
}
type UEContextReleaseRequest struct {
}
type UEInformationTransfer struct {
}
type UERadioCapabilityInfoIndication struct {
}
type UETNLABindingReleaseRequest struct {
}
type UplinkNASTransport struct {
}
type UplinkNonUEAssociatedNRPPaTransport struct {
}
type UplinkRANConfigurationTransfer struct {
}
type UplinkRANEarlyStatusTransfer struct {
}
type UplinkRANStatusTransfer struct {
}
type UplinkRIMInformationTransfer struct {
}
type UplinkUEAssociatedNRPPaTransport struct {
}
