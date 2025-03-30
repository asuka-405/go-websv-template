package libsailpoint

import (
	"encoding/json"
	"time"
)

type Source struct {
	Description               string                 `json:"description"`
	Owner                     Identity               `json:"owner"`
	Cluster                   Cluster                `json:"cluster"`
	AccountCorrelationConfig  Entity                 `json:"accountCorrelationConfig"`
	AccountCorrelationRule    *Entity                `json:"accountCorrelationRule"`
	ManagerCorrelationMapping *Entity                `json:"managerCorrelationMapping"`
	ManagerCorrelationRule    *Entity                `json:"managerCorrelationRule"`
	Schemas                   []Schema               `json:"schemas"`
	PasswordPolicies          interface{}            `json:"passwordPolicies"`
	Features                  []string               `json:"features"`
	Type                      string                 `json:"type"`
	Connector                 string                 `json:"connector"`
	ConnectorClass            string                 `json:"connectorClass"`
	ConnectorAttributes       map[string]interface{} `json:"connectorAttributes"`
	DeleteThreshold           int                    `json:"deleteThreshold"`
	Authoritative             bool                   `json:"authoritative"`
	Healthy                   bool                   `json:"healthy"`
	Status                    string                 `json:"status"`
	Since                     string                 `json:"since"`
	ConnectorID               string                 `json:"connectorId"`
	ConnectorName             string                 `json:"connectorName"`
	ConnectionType            string                 `json:"connectionType"`
	ConnectorImplementationID string                 `json:"connectorImplementationId"`
	ManagementWorkgroup       interface{}            `json:"managementWorkgroup"`
	CredentialProviderEnabled bool                   `json:"credentialProviderEnabled"`
	Category                  interface{}            `json:"category"`
	AccountsFile              interface{}            `json:"accountsFile"`
	ID                        string                 `json:"id"`
	Name                      string                 `json:"name"`
	Created                   string                 `json:"created"`
	Modified                  string                 `json:"modified"`
}

// Identity, Cluster, Entity, and Schema struct definitions
type Identity struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Cluster struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Entity struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Schema struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	Name string `json:"name"`
}

type EntitlementSource struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type EntitlementAttributes struct {
	SecurityRoleAccessToSensitiveData bool   `json:"securityRoleAccessToSensitiveData"`
	LegalEntityID                     string `json:"legalEntityId"`
	LegalEntityName                   string `json:"legalEntityName"`
	OrganizationRoleName              string `json:"organizationRoleName"`
	SecurityRoleUserLicenseType       string `json:"securityRoleUserLicenseType"`
	SecurityRoleName                  string `json:"securityRoleName"`
	SecurityRoleIdentifier            string `json:"securityRoleIdentifier"`
}

type EntitlementManuallyUpdatedFields struct {
	DisplayName bool `json:"DISPLAY_NAME"`
}

type EntitlementAccessModelMetadata struct {
	Attributes []interface{} `json:"attributes"`
}

type Entitlement struct {
	Attribute              string                           `json:"attribute"`
	Value                  string                           `json:"value"`
	Description            *string                          `json:"description"`
	SourceSchemaObjectType string                           `json:"sourceSchemaObjectType"`
	Privileged             bool                             `json:"privileged"`
	CloudGoverned          bool                             `json:"cloudGoverned"`
	Requestable            bool                             `json:"requestable"`
	Attributes             EntitlementAttributes            `json:"attributes"`
	Source                 EntitlementSource                `json:"source"`
	Owner                  json.RawMessage                  `json:"owner"`
	DirectPermissions      []interface{}                    `json:"directPermissions"`
	Segments               []interface{}                    `json:"segments"`
	ManuallyUpdatedFields  EntitlementManuallyUpdatedFields `json:"manuallyUpdatedFields"`
	AccessModelMetadata    EntitlementAccessModelMetadata   `json:"accessModelMetadata"`
	Modified               time.Time                        `json:"modified"`
	Created                time.Time                        `json:"created"`
	ID                     string                           `json:"id"`
	Name                   string                           `json:"name"`
}

type AccessProfileReq struct {
	Name         string             `json:"name"`
	Owner        *APReqOwner        `json:"owner"`
	Source       *APReqSource       `json:"source"`
	Entitlements []APReqEntitlement `json:"entitlements,omitempty"`
	Description  *string            `json:"description,omitempty"`
	Enabled      bool               `json:"enabled,omitempty"`
	Requestable  bool               `json:"requestable,omitempty"`
	Segments     []string           `json:"segments,omitempty"`
	Provisioning *APReqProvisioning `json:"provisioningCriteria,omitempty"`
}

type APReqOwner struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

type APReqSource struct {
	ID   string  `json:"id"`
	Name string  `json:"name"`
	Type *string `json:"type,omitempty"`
}

type APReqEntitlement struct {
	ID   string  `json:"id"`
	Name string  `json:"name"`
	Type *string `json:"type,omitempty"`
}

type APReqProvisioning struct {
	Operation string              `json:"operation"`
	Children  []APReqProvisioning `json:"children,omitempty"`
	Attribute *string             `json:"attribute,omitempty"`
	Value     *string             `json:"value,omitempty"`
}
