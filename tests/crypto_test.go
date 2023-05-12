package tests

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"xs/utils"
)

func TestHash(t *testing.T) {
	hash := utils.GenHash("bla2", "bla1")
	assert.Equal(t, "e64938fc6124b4dfa8a2f225cc4998df473cbd6710c364684a1f42f6257d8f8c", hash)
}

func TestSeedGeneration(t *testing.T) {
	seed := utils.GenMnemonic()
	worlds := strings.Split(seed, " ")
	assert.True(t, 10 < len(worlds))
	assert.True(t, 100 < len(seed))
}

/*
it('should generate private key', async () => {
        const ct= new CryptoTools();
        const privateKey = await ct.privateKeyFromSeed("bachelor spy list giggle velvet adjust impulse weasel blush grant hole concert");
        expect(privateKey.buffer.byteLength).toBe(32);
        const hex = Buffer.from(privateKey).toString('hex');
        expect(hex).toBe("c6e35b65ae3a4b385fbc4583d2704c5e0f0bbde105c83850b8dd79e6229ebf6b");
    });
*/

func TestPrivateKeyGeneration(t *testing.T) {
	pk := utils.PrivateKeyFromSeed("bachelor spy list giggle velvet adjust impulse weasel blush grant hole concert")
	println(pk)
	//bytes := pk.Bytes()
	//assert.Equal(t, 32, len(bytes))
	//hexString := hex.EncodeToString(bytes)
	//assert.Equal(t, "c6e35b65ae3a4b385fbc4583d2704c5e0f0bbde105c83850b8dd79e6229ebf6b", hexString)
}

/*

describe('Tokens Tools', () => {




    it('sing and check', async () => {
        const ct = new CryptoTools();
        const privateKey = await ct.privateKeyFromSeed("bachelor spy list giggle velvet adjust impulse weasel blush grant hole concert");
        const publicKey = await ct.publicKeyFromPrivateKey(privateKey)
        const dataExpired = moment().add(14, "day").valueOf();
        const token = createToken({test: "test", expired: "" + dataExpired.valueOf()}, privateKey);
        const body = readToken(token, publicKey);
        console.log("GENERATED", body);
    });

});
*/

//func TestSingAndVerify(t *testing.T) {
//	privateKey := utils.PrivateKeyFromSeed("bachelor spy list giggle velvet adjust impulse weasel blush grant hole concert")
//	publicKey := privateKey.PubKey()
//	//dataExpired := time.Now().Add(14 * 24 * time.Hour).Unix()
//	//token := utils.CreateToken(map[string]string{"test": "test", "expired": string(dataExpired) }, privateKey)
//	//body := utils.ReadToken(token, publicKey)
//	//assert.Equal(t, "test", body["test"])
//}
