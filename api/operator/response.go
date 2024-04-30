//go:generate easyjson -all -no_std_marshalers

package operator

type ServiceStatusType string

const (
	ServiceStatusTypeActive ServiceStatusType = "Up"
	ServiceStatusTypeDown   ServiceStatusType = "Down"
)

type EigenNodeResponse struct {
	NodeName    string `json:"node_name"`
	SpecVersion string `json:"spec_version"`
	NodeVersion string `json:"node_version"`
}

type ErrorMessageResponse struct {
	Message string `json:"message"`
}

type EigenNodeService struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Status      ServiceStatusType `json:"status"`
}

type EigenNodeServicesResponse struct {
	Services []EigenNodeService `json:"services"`
}
