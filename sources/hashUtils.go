package sources

import "golang.org/x/crypto/bcrypt"

// func DeriveKey(password, salt []byte) ([]byte, []byte, error) {
// 	if salt == nil {
// 		salt = make([]byte, 32)
// 		if _, err := rand.Read(salt); err != nil {
// 			return nil, nil, err
// 		}
// 	}

// 	key, err := scrypt.Key(password, salt, 1048576, 8, 1, 32)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	return key, salt, nil
// }

//deriveKey generates a new NaCl key from a passphrase and salt.
const SECRET_KEYSIZE = 32
const SALTSIZE = 32

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// func DeriveKey(pass, salt []byte) (*[SECRET_KEYSIZE]byte, error) {
// 	var naclKey = new([SECRET_KEYSIZE]byte)
// 	key, err := scrypt.Key(pass, salt, 1048576, 8, 1, SECRET_KEYSIZE)
// 	if err != nil {
// 		return nil, err
// 	}

// 	copy(naclKey[:], key)
// 	util.Zero(key)
// 	return naclKey, nil
// }

// func Encrypt(pass, message []byte) ([]byte, error) {
// 	// salt, err := util.RandBytes(SALTSIZE)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	salt := make([]byte, SALTSIZE)
// 	_, err := io.ReadFull(rand.Reader, salt)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	key, err := DeriveKey(pass, salt)
// 	if err != nil {
// 		return nil, err
// 	}

// 	out, err := secret.Encrypt(key, message)
// 	util.Zero(key[:]) // Zero key immediately after
// 	if err != nil {
// 		return nil, err
// 	}

// 	out = append(salt, out...)
// 	return out, nil
// }

// const Overhead = SALTSIZE + secretbox.Overhead + secret.NonceSize

// func Decrypt(pass, message []byte) ([]byte, error) {
// 	if len(message) < Overhead {
// 		return nil, err
// 	}

// 	key, err := DeriveKey(pass, message[:SALTSIZE])
// 	if err != nil {
// 		return nil, err
// 	}

// 	out, err := secret.Decrypt(key, message[SALTSIZE:])
// 	util.Zero(key[:]) // Zero key immediately after
// 	if err != nil {
// 		return nil, err
// 	}

// 	return out, nil
// }
