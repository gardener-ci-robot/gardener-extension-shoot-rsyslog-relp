// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RsyslogRelpConfig configuration resource.
type RsyslogRelpConfig struct {
	metav1.TypeMeta `json:",inline"`

	// Target is the target server to connect to via relp.
	Target string `json:"target"`
	// Port is the TCP port to use when connecting to the target server.
	Port int `json:"port"`
	// LoggingRules contain a list of LoggingRules that are used to determine which logs are
	// sent to the target server by the the rsyslog relp action.
	LoggingRules []LoggingRule `json:"loggingRules,omitempty"`
	// TLS hods the TLS config.
	// +optional
	TLS *TLS `json:"tls,omitempty"`
	// RebindInterval is the rebind interval for the rsyslog relp action.
	// +optional
	RebindInterval *int `json:"rebindInterval,omitempty"`
	// Timeout is the connection timeout for the rsyslog relp action.
	// +optional
	Timeout *int `json:"timeout,omitempty"`
	// ResumeRetryCount is the resume retry count for the rsyslog relp action.
	// +optional
	ResumeRetryCount *int `json:"resumeRetryCount,omitempty"`
	// ReportSuspensionContinuation determines whether suspension continuation in the relp action
	// should be reported.
	// +optional
	ReportSuspensionContinuation *bool `json:"reportSuspensionContinuation,omitempty"`
}

// TLS contains options for the tls connection to the target server.
type TLS struct {
	// Enabled determines whether TLS encryption should be used for the connection
	// to the target server.
	Enabled bool `json:"enabled"`
	// SecretReferenceName is the name of the reference for the secret
	// containing the certificates for the TLS connection when encryption is enabled.
	// +optional
	SecretReferenceName *string `json:"secretReferenceName,omitempty"`
	// PermittedPeer is the name of the rsyslog relp permitted peer.
	// Only peers which have been listed in this parameter may be connected to.
	// +optional
	PermittedPeer []string `json:"permittedPeer,omitempty"`
	// AuthMode is the mode used for mutual authentication.
	// Possible values are "fingerprint" or "name".
	// +optional
	AuthMode *AuthMode `json:"authMode,omitempty"`
}

// LoggingRule contains options that determines which logs are sent to the target server.
type LoggingRule struct {
	// ProgramNames are the names of the programs for which logs are sent to the target server.
	// +optional
	ProgramNames []string `json:"programNames,omitempty"`
	// Severity determines which logs are sent to the target server based on their severity.
	Severity int `json:"severity"`
}

// AuthMode is the type of authentication mode that can be used for the rsyslog relp connection to the target server.
type AuthMode string

const (
	// AuthModeName specifies the rsyslog name authentication mode.
	AuthModeName AuthMode = "name"
	// AuthModeFingerPrint specifies the rsyslog fingerprint authentication mode.
	AuthModeFingerPrint AuthMode = "fingerprint"
)
