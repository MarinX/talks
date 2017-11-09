// Key is the type used for keys. This is the main structure used for encrypting/decrypting data.
type Key struct {
	KeyBase64    string            
	Key          []byte             
	DecryptedKey []byte             
	Context      map[string]*string 
	MasterID     string            
	Region       string            
}

func GenerateKey(masterID string, context map[string]*string) (*Key, error) {

	key, decryptedKey, err := generateNewDataKey(masterID, context)

	//error handling

	encoded := base64.StdEncoding.EncodeToString(key)

	return &Key{values_set}, nil

}
