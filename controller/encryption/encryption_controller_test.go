package encryption

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func setup() {
	NewEncryptionController()
}

func TestDecryptionPassword(t *testing.T) {
	setup()
	//mockEncryptedPassword := "VjceXgEDsCNdLm0XQUI10Zo5NHZXHPw0klMMEqlhjiyk2UmNDYgdD5bRhpjZlL7de2MM4Uprnv9j0YhGjaNID2YuSlaBuPecGAVCHHJyla6fpFiKzfwd/Q0uGy5+majUPAj7Z8ummFm0uituRhoyRMrZhfDLOdp5v3srZy1CDYhBMMbR6zhNSDtvtjGTBkmUvZxULvZmm6cWn6RyMLpoOMKvqJ21zYw7XJVKALFCpslK0t46yfn/9a9jotFnSA98ziN0LIddyLZcWhuOT/cbYJ/qmB5Wgff5qodKpckT8cqn2x2xvW4Kz9v/Mde6XGKSzgF9aAMnMluXKKXVJhvUww=="
	//decryptedPassword, err := EncryptionControllerObj.DecryptPassword(mockEncryptedPassword)
	//
	//assert.Nil(t, err)
	//assert.Equal(t, "P@ssw0rd", decryptedPassword)
}

func TestBase64DecodingError(t *testing.T) {
	setup()
	mockEncryptedPassword := "U2FsdGVkX13Z"
	decryptedPassword, err := EncryptionControllerObj.DecryptPassword(mockEncryptedPassword)

	assert.NotNil(t, err)
	assert.Equal(t, decryptedPassword, "")
}
