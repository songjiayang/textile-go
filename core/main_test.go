package core_test

import (
	"crypto/rand"
	"github.com/op/go-logging"
	. "github.com/textileio/textile-go/core"
	"github.com/textileio/textile-go/keypair"
	libp2pc "gx/ipfs/Qme1knMqwt1hKZbc1BmQFmnm9f36nyQGwXxPGVpVJ9rMK5/go-libp2p-crypto"
	"os"
	"testing"
)

var repoPath = "testdata/.textile"
var node *Textile

func TestInitRepo(t *testing.T) {
	os.RemoveAll(repoPath)
	accnt := keypair.Random()
	if err := InitRepo(InitConfig{
		Account:  *accnt,
		RepoPath: repoPath,
		LogLevel: logging.DEBUG,
	}); err != nil {
		t.Errorf("init node failed: %s", err)
	}
}

func TestNewTextile(t *testing.T) {
	var err error
	node, err = NewTextile(RunConfig{
		RepoPath: repoPath,
		LogLevel: logging.DEBUG,
	})
	if err != nil {
		t.Errorf("create node failed: %s", err)
	}
}

func TestCore_Start(t *testing.T) {
	if err := node.Start(); err != nil {
		t.Errorf("start node failed: %s", err)
	}
	<-node.OnlineCh()
}

func TestCore_Started(t *testing.T) {
	if !node.Started() {
		t.Errorf("should report started")
	}
}

func TestCore_Online(t *testing.T) {
	if !node.Online() {
		t.Errorf("should report online")
	}
}

func TestCore_CafeRegister(t *testing.T) {
	// TODO
}

func TestCore_AddThread(t *testing.T) {
	sk, _, err := libp2pc.GenerateEd25519Key(rand.Reader)
	if err != nil {
		t.Error(err)
	}
	thrd, err := node.AddThread("test", sk, true)
	if err != nil {
		t.Errorf("add thread failed: %s", err)
		return
	}
	if thrd == nil {
		t.Error("add thread didn't return thread")
	}
}

func TestCore_AddPhoto(t *testing.T) {
	added, err := node.AddPhoto("../photo/testdata/image.jpg")
	if err != nil {
		t.Errorf("add photo failed: %s", err)
		return
	}
	if len(added.Id) == 0 {
		t.Errorf("add photo got bad id")
	}
	// test adding an image w/o the orientation tag
	added2, err := node.AddPhoto("../photo/testdata/image-no-orientation.jpg")
	if err != nil {
		t.Errorf("add photo w/o orientation tag failed: %s", err)
		return
	}
	if len(added2.Id) == 0 {
		t.Errorf("add photo w/o orientation tag got bad id")
	}
}

func TestCore_Stop(t *testing.T) {
	err := node.Stop()
	if err != nil {
		t.Errorf("stop node failed: %s", err)
	}
}

func TestCore_StartedAgain(t *testing.T) {
	if node.Started() {
		t.Errorf("should report stopped")
	}
}

func TestCore_OnlineAgain(t *testing.T) {
	if node.Online() {
		t.Errorf("should report offline")
	}
}

func TestCore_Teardown(t *testing.T) {
	os.RemoveAll(node.GetRepoPath())
	os.RemoveAll(node1.GetRepoPath())
	os.RemoveAll(node2.GetRepoPath())
}