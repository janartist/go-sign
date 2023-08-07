package sign_test

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"math/big"
	"testing"

	"github.com/janartist/go-sign/sign"
	"github.com/stretchr/testify/assert"
)

var data = []byte("我是内容12138")

var rsaPrivateKey = &rsa.PrivateKey{
	PublicKey: rsa.PublicKey{
		N: fromBase10("9353930466774385905609975137998169297361893554149986716853295022578535724979677252958524466350471210367835187480748268864277464700638583474144061408845077"),
		E: 65537,
	},
	D: fromBase10("7266398431328116344057699379749222532279343923819063639497049039389899328538543087657733766554155839834519529439851673014800261285757759040931985506583861"),
	Primes: []*big.Int{
		fromBase10("98920366548084643601728869055592650835572950932266967461790948584315647051443"),
		fromBase10("94560208308847015747498523884063394671606671904944666360068158221458669711639"),
	},
}

func TestHMACSigner(t *testing.T) {
	secretKey := []byte("se1")
	secretKey2 := []byte("se1x")
	t.Run("sha256", func(t *testing.T) {
		MACSigner := sign.NewHMACSigner(sha256.New, secretKey)
		MACSigner2 := sign.NewHMACSigner(sha256.New, secretKey2)
		signature, err := MACSigner.Sign(data)
		assert.Nil(t, err)
		ok, err := MACSigner.Verify(data, signature)
		assert.Nil(t, err)
		assert.Equal(t, ok, true)
		ok, err = MACSigner2.Verify(data, signature)
		assert.Nil(t, err)
		assert.Equal(t, ok, false)
	})
	t.Run("md5", func(t *testing.T) {
		MACSigner := sign.NewHMACSigner(md5.New, secretKey)
		MACSigner2 := sign.NewHMACSigner(sha256.New, secretKey2)
		signature, err := MACSigner.Sign(data)
		assert.Nil(t, err)
		ok, err := MACSigner.Verify(data, signature)
		assert.Nil(t, err)
		assert.Equal(t, ok, true)
		ok, err = MACSigner2.Verify(data, signature)
		assert.Nil(t, err)
		assert.Equal(t, ok, false)
	})
}

func TestRSASigner(t *testing.T) {

	RSASignerPrivateKey := sign.NewRSASignerPrivateKey(rsaPrivateKey, crypto.SHA256)
	signature, err := RSASignerPrivateKey.Sign(data)
	assert.Nil(t, err)
	ok, err := RSASignerPrivateKey.Verify(data, signature)
	assert.Nil(t, err)
	assert.Equal(t, ok, true)

	RSASignerPublicKey := sign.NewRSASignerPublicKey(&rsaPrivateKey.PublicKey, crypto.SHA256)
	ok, err = RSASignerPublicKey.Verify(data, signature)
	assert.Nil(t, err)
	assert.Equal(t, ok, true)
}

func TestECDSASigner(t *testing.T) {
	ecdsaPrivateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if !assert.Nil(t, err) {
		return
	}

	ECDSASigner := sign.NewECDSASigner(ecdsaPrivateKey)
	signature, err := ECDSASigner.Sign(data)
	assert.Nil(t, err)
	ok, err := ECDSASigner.Verify(data, signature)
	assert.Nil(t, err)
	assert.Equal(t, ok, true)
}

func fromBase10(base10 string) *big.Int {
	i, ok := new(big.Int).SetString(base10, 10)
	if !ok {
		panic("bad number: " + base10)
	}
	return i
}
