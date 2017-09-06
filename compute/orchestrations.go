// Manages orchestration resources
// For clarity in the following resource methods, Orc is used
// as an abbreviation for Orchestration(s). Otherwise, these method
// and function headers will get quite ugly.
// This resource solely supports v1 of the orchestrations resource.
// A future resource will likely need to be added to support v2 of
// the orchestrations API.

package compute

const (
	OrcDescription  = "Orchestrations"
	OrcContainerPath = "/orchestration/"
	OrcRootPath = "/orchestration"
)

// Allows management of orchestration resources
type OrcClient struct {
	ResourceClient
}

// Returns an orchestration client from an existing compute client
func (c *ComputeClient) Orchestrations() *OrcClient {
	return &OrcClient{
		ResourceClient: ResourceClient{
			ComputeClient: c,
			ResourceDescription: OrcDescription,
			ContainerPath: OrcContainerPath,
			ResourceRootPath: OrcRootPath,
		},
	}
}

// Holds all the information about an orchestration resource
type OrcInfo struct {
	// Shows the default account for your domain
	Account string `json:"account"`
	// Description of the plan
	Description string `json:"description"`
	// TODO: Add Info, once determined what's actually in the nested response
	// Info []InfoObject `json:"info"`
	// Name of the object
	Name string `json:"name"`
	// User-written Plan for Orchestration
	OPlans string `json:"oplans"`
	// Relationships between objects
	Relationships []Relationship `json:"relationships"`
	// Schedules
	Schedule Schedule `json:"schedule"`
	// Status of the orchestration
	Status string `json:"status"`
	// Status timestamp, time the current view was generated
	StatusTimestamp string `json:"status_timestamp"`
	// URI
	Uri string `json:"uri"`
	// Two part name of the user who has most recently taken an action
	// on the orchestration
	User string `json:"user"`
}

// Relationship nested struct
type Relationship struct {
	// TODO: Determine what fields are present here.
}

// Schedule nested struct
type Schedule struct {
	// Start Time(?)
	// Stop Time(?)
	// TODO: Determine what fields are present here.
}

// All of the possible fields used to create an orchestrations resource
type CreateOrcInput struct {
	// Description of the orchestration
	Description string `json:"description"`
	// Name of the Orchestration
	Name string `json:"name"`
	// User input of the orchestration plan - Must pass JSON parsing check
	OPlans string `json:"oplans"`
	// Relationships between object plans
	// (See: http://www.oracle.com/pls/topic/lookup?ctx=cloud&id=STCSG-GUID-1896C799-49A6-42B8-9813-7DE5695267FE__RELATIONSHIPS-58824D2E)
	Relationships []Relationship `json:"relationships"`
	// Schedule
	Schedule Schedule `json:"schedule"`
}

// Creates an orchestration, and returns the information about the orchestration
func (c *OrcClient) CreateOrc(input *CreateOrcInput) (*OrcInfo, error) {

}

// Input structure for qualification information to lookup
// the desired Orchestration object
type GetOrcInput struct {
	// Name of the orchestration
	// Required
	Name string `json:"name"`
}

// Get details about an orchestration, return an info struct
func (c *OrcClient) GetOrc(input *GetOrcInput) (*OrcInfo, error) {

}

// Update information to update an orchestration
type UpdateOrcInput struct {
	// TODO
}

// Update an orchestration, return an information struct
func (c *OrcClient) UpdateOrc(input *UpdateOrcInput) (*OrcInfo, error) {

}

// Delete an orchestration input
type DeleteOrcInput struct {
	// Name of the orchestration
	Name string `json:"name"`
}

func (c *OrcClient) DeleteOrc(input *DeleteOrcInput) (*OrcInfo, error) {

}
