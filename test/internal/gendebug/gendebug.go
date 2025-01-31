package gendebug

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/seq"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

type InitiatingMessageValue struct {
	Present                int
	Opt                    fp.Option[string]
	AMFConfigurationUpdate *AMFConfigurationUpdate `aper:"valueExt,referenceFieldValue:0,criticality:reject"`
	HandoverCancel         *HandoverCancel         `aper:"valueExt,referenceFieldValue:10,criticality:reject"`
	HandoverRequired       *HandoverRequired       `aper:"valueExt,referenceFieldValue:12,criticality:reject"`
	HandoverRequest        *HandoverRequest        `aper:"valueExt,referenceFieldValue:13,criticality:reject"`
}

// @fp.Generate
var _ = genfp.GenerateFromStructs{
	File:    "gendebug_generated.go",
	Imports: seq.Of(genfp.ImportPackage{Package: "fmt", Name: "fmt"}),
	List:    seq.Of(genfp.TypeOf[InitiatingMessageValue]()),
	Template: `
		func Hello{{.N}}(arg {{.N}}) {
			{{$st := .N}}
			{{range .N.Fields }}
				fmt.Printf("{{$st}}.{{.Type.Name}} = %v\n", arg.{{.Name}} )
				fmt.Printf("elem type = {{.ElemType}}")

			{{end}}
		}
	`,
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
