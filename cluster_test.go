package ovirtapi_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/emsl-msc/ovirtapi"
)

func TestCluster(t *testing.T) {
	username := os.Getenv("OVIRT_USERNAME")
	if username == "" {
		t.Error("OVIRT_USERNAME is not set")
	}
	password := os.Getenv("OVIRT_PASSWORD")
	if password == "" {
		t.Error("OVIRT_PASSWORD is not set")
	}
	url := os.Getenv("OVIRT_URL")
	if url == "" {
		t.Error("OVIRT_URL is not set")
	}
	con, err := ovirtapi.NewConnection(url, username, password)
	con.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG_TRANSPORT"))
	if err != nil {
		t.Error("error creating con connection", err)
		return
	}
	newCluster := con.NewCluster()
	newCluster.Name = "test-cluster"
	newCluster.CPU = &ovirtapi.CPU{Type: "Intel Haswell-noTSX Family"}
	newCluster.DataCenter = &ovirtapi.Link{ID: "00000001-0001-0001-0001-000000000311"}
	err = newCluster.Save()
	if err != nil {
		t.Fatal("Error creating new cluster", err)
	}
	retrievedCluster, err := con.GetCluster(newCluster.ID)
	if err != nil {
		t.Fatal("Error retrieving cluster", err)
	}
	retrievedCluster.Description = "about to delete"
	err = retrievedCluster.Save()
	if err != nil {
		t.Fatal("Error updating cluster", err)
	}
	err = retrievedCluster.Delete()
	if err != nil {
		t.Fatal("Error Deleting cluster", err)
	}
}
