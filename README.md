# crypter-go
Simple encryption/decryption using the AES with 32 byte key.

**Example:**
```
key, _ := crypter.GenerateKey()

crypter, _ := crypter.New(key)

plaintext := []byte("My Super Secret Code Stuff")
fmt.Println(string(plaintext))

ciphertext, _ := crypter.Encrypt(plaintext)
fmt.Println(string(ciphertext))

plaintext, _ = crypter.Decrypt(ciphertext)
fmt.Println(string(plaintext))
```

**Output:**
```
My Super Secret Code Stuff
���#���f#���1=���O���&☻E����b��☼�♂G�@RO►"ɡ↨��
My Super Secret Code Stuff
```
