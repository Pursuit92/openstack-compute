package compute

import (
	"testing"
	"os"
)

func TestKeyPairs(t *testing.T) {
	cc,err := NewClient(os.Getenv("OS_AUTH_URL"))
	if err != nil {
		t.Error(err)
		return
	}
	cc.PasswordAuth(os.Getenv("OS_USERNAME"),os.Getenv("OS_PASSWORD"))
	cc.TenantName(os.Getenv("OS_TENANT_NAME"))
	err = cc.Authenticate()
	if err != nil {
		t.Error(err)
		return
	}

	_,err = cc.Keypairs()
	if err != nil {
		t.Error(err)
		return
	}

	key := &Keypair{Name: "testingKey"}

	err = cc.NewKeypair(key)
	if err != nil {
		t.Error(err)
		return
	}

	key,err = cc.GetKeypair(key.Name)
	if err != nil {
		t.Error(err)
		return
	}

	err = cc.DeleteKeypair(key)
	if err != nil {
		t.Error(err)
		return
	}
}



