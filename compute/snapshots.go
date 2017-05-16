package compute

// SnapshotsClient is a client for the Snapshot functions of the Compute API.
type SnapshotsClient struct {
	ResourceClient
}

// Snapshots obtains an SnapshotsClient which can be used to access to the
// Snapshot functions of the Compute API
func (c *Client) Snapshots() *SnapshotsClient {
	return &SnapshotsClient{
		ResourceClient: ResourceClient{
			Client:              c,
			ResourceDescription: "Snapshot",
			ContainerPath:       "/snapshot/",
			ResourceRootPath:    "/snapshot",
		}}
}

type SnapshotDelay string

const (
	SnapshotDelayShutdown SnapshotDelay = "shutdown"
)

// SnapshotInfo describes an existing Snapshot.
type Snapshot struct {
	// Shows the default account for your identity domain.
	Account string `json:"account"`
	// Timestamp when this request was created.
	CreationTime string `json:"creation_time"`
	// Snapshot of the instance is not taken immediately.
	Delay SnapshotDelay `json:"delay"`
	// A description of the reason this request entered "error" state.
	ErrorReason string `json:"error_reason"`
	// Name of the instance
	Instance string `json:"instance"`
	// Name of the machine image generated from the instance snapshot request.
	MachineImage string `json:"machineimage"`
	// Name of the instance snapshot request.
	Name string `json:"name"`
	// Not used
	Quota string `json:"quota"`
	// The state of the request.
	State string `json:"state"`
	// Uniform Resource Identifier
	URI string `json:"uri"`
}

// CreateSnapshotInput defines an Snapshot to be created.
type CreateSnapshotInput struct {
	//The name of the account that contains the credentials and access details of
	// Oracle Storage Cloud Service. The machine image file is uploaded to the Oracle
	// Storage Cloud Service account that you specify.
	// Optional
	Account string `json:"account,omitempty"`
	// Use this option when you want to preserve the custom changes you have made
	// to an instance before deleting the instance. The only permitted value is shutdown.
	// Snapshot of the instance is not taken immediately. It creates a machine image which
	// preserves the changes you have made to the instance, and then the instance is deleted.
	// Note: This option has no effect if you shutdown the instance from inside it. Any pending
	// snapshot request on that instance goes into error state. You must delete the instance
	// (DELETE /instance/{name}).
	// Optional
	Delay string `json:"delay,omitempty"`
	// Name of the instance that you want to clone.
	// Required
	Instance string `json:"instance"`
	// Specify the name of the machine image created by the snapshot request.
	// Object names can contain only alphanumeric characters, hyphens, underscores, and periods.
	// Object names are case-sensitive.
	// If you don't specify a name for this object, then the name is generated automatically.
	// Optional
	MachineImage string `json:"machineimage,omitempty"`
}

// CreateSnapshot creates a new Snapshot
func (c *SnapshotsClient) CreateSnapshot(createInput *CreateSnapshotInput) (*Snapshot, error) {
	createInput.Account = c.getQualifiedACMEName(createInput.Account)
	createInput.Instance = c.getQualifiedName(createInput.Instance)
	createInput.MachineImage = c.getQualifiedName(createInput.MachineImage)

	var snapshotInfo Snapshot
	if err := c.createResource(&createInput, &snapshotInfo); err != nil {
		return nil, err
	}

	return c.success(&snapshotInfo)
}

// GetSnapshotInput describes the snapshot to get
type GetSnapshotInput struct {
	// The name of the Snapshot
	Name string `json:name`
}

// GetSnapshot retrieves the Snapshot with the given name.
func (c *SnapshotsClient) GetSnapshot(getInput *GetSnapshotInput) (*Snapshot, error) {
	getInput.Name = c.getQualifiedName(getInput.Name)
	var snapshotInfo Snapshot
	if err := c.getResource(getInput.Name, &snapshotInfo); err != nil {
		return nil, err
	}

	return c.success(&snapshotInfo)
}

// DeleteSnapshotInput describes the snapshot to delete
type DeleteSnapshotInput struct {
	// The name of the Snapshot
	Name string `json:name`
}

// DeleteSnapshot deletes the Snapshot with the given name.
func (c *SnapshotsClient) DeleteSnapshot(deleteInput *DeleteSnapshotInput) error {
	return c.deleteResource(deleteInput.Name)
}

func (c *SnapshotsClient) success(snapshotInfo *Snapshot) (*Snapshot, error) {
	c.unqualify(&snapshotInfo.Account)
	c.unqualify(&snapshotInfo.Instance)
	c.unqualify(&snapshotInfo.MachineImage)
	c.unqualify(&snapshotInfo.Name)
	return snapshotInfo, nil
}
