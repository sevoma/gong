package gong

import (
	"fmt"
	. "gopkg.in/check.v1"
	"testing"
)

func TestClient(t *testing.T) { TestingT(t) }

type ClientSuite struct{}

var _ = Suite(&ClientSuite{})

type FakeClient struct{}

func (f *FakeClient) FormatField(fieldName string, value string) string {
	return ""
}

func (receiver *FakeClient) GetAuthFields() map[string]bool {
	return map[string]bool{}
}

func (c *FakeClient) GetName() string {
	return "fakeclient"
}

func (c *FakeClient) Authenticate(field map[string]string) bool {
	return false
}

func (c *FakeClient) Start(issueType, issueId string) (string, error) {
	return fmt.Sprintf("%s/%s", issueType, issueId), nil
}

func (s *ClientSuite) TestClientStartIssue(c *C) {
	client := &FakeClient{}
	branchName, _ := Start(client, "feature", "111")

	c.Assert(branchName, Equals, "feature/111")
}