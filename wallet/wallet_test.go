package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"io/fs"
	"reflect"
	"testing"
)

const (
	testKey     = "30770201010420699d3796268174e6c35204029ccb7c1a9e0576b691b767598649beee964b35a1a00a06082a8648ce3d030107a1440342000469e6cb9f1f93ca48e1660c05160b9f3df72840eb4c2b72a8344d7728ca65ad742a53be9fc1cf2a06b815cb12612d2965cc3b20b168042049d32b99e333476086"
	testPayload = "0054ff88200ccf7f640131fed68d231574a18db7b08954aff7a7df3ef5944d2f"
	testSig     = "41e5246987d7147644ed8b05fbc84bd9ca384c4bb285f7870af439ea397f08d82d1301cb0f339fa10f416bf4d59c1222dc2b3ad5f905575bd8fa6803fc4a63dd"
)

type fakeLayer struct {
	fakeHasWalletFile func() bool
}

func (f fakeLayer) hasWalletFile() bool {
	return f.fakeHasWalletFile()

}

func (fakeLayer) writeFile(name string, data []byte, perm fs.FileMode) error {
	return nil
}

func (fakeLayer) readFile(name string) ([]byte, error) {
	return x509.MarshalECPrivateKey(makeTestWallet().privateKey)
}

func TestWallet(t *testing.T) {
	t.Run("New Wallet is created", func(t *testing.T) {
		files = fakeLayer{
			fakeHasWalletFile: func() bool { return false },
		}
		w := Wallet()
		if reflect.TypeOf(w) != reflect.TypeOf(&wallet{}) {
			t.Error("New Wallet should return a new wallet instance")
		}
	})
	t.Run("Wallet is restrored", func(t *testing.T) {
		files = fakeLayer{
			fakeHasWalletFile: func() bool { return true },
		}
		w = nil
		w := Wallet()
		if reflect.TypeOf(w) != reflect.TypeOf(&wallet{}) {
			t.Error("New Wallet should return a new wallet instance")
		}
	})
}

func makeTestWallet() *wallet {
	w := &wallet{}
	b, _ := hex.DecodeString(testKey)
	key, _ := x509.ParseECPrivateKey(b)
	w.privateKey = key
	w.Address = aFromK(key)
	return w
}

func TestSign(t *testing.T) {
	s := Sign(testPayload, makeTestWallet())
	_, err := hex.DecodeString(s)
	if err != nil {
		t.Errorf("Sign() should return a hex encoded string, got %s", s)
	}
}

func TestVerify(t *testing.T) {
	type test struct {
		input string
		ok    bool
	}
	tests := []test{
		{testPayload, true},
		{"0154ff88200ccf7f640131fed68d231574a18db7b08954aff7a7df3ef5944d2f", false},
	}
	for _, tc := range tests {
		w := makeTestWallet()
		ok := Verify(testSig, tc.input, w.Address)
		if ok != tc.ok {
			t.Error("Verify() could not verify testSignature and Pauload")
		}
	}
}

func TestRestoreBigInts(t *testing.T) {
	_, _, err := restoreBigInts("xx")
	if err == nil {
		t.Error("restoreBigInts should return error when payload is not hex.")
	}
}
