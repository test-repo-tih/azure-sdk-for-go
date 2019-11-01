## Breaking Changes

### Removed Funcs

1. PossibleJSONWebKeyCurveNameValues() []JSONWebKeyCurveName
1. PossibleJSONWebKeyEncryptionAlgorithmValues() []JSONWebKeyEncryptionAlgorithm
1. PossibleJSONWebKeyOperationValues() []JSONWebKeyOperation
1. PossibleJSONWebKeySignatureAlgorithmValues() []JSONWebKeySignatureAlgorithm
1. PossibleJSONWebKeyTypeValues() []JSONWebKeyType

## Signature Changes

### Const Types

1. Decrypt changed type from JSONWebKeyOperation to KeyOperation
1. EC changed type from JSONWebKeyType to KeyType
1. ECHSM changed type from JSONWebKeyType to KeyType
1. ES256 changed type from JSONWebKeySignatureAlgorithm to SignatureAlgorithm
1. ES256K changed type from JSONWebKeySignatureAlgorithm to SignatureAlgorithm
1. ES384 changed type from JSONWebKeySignatureAlgorithm to SignatureAlgorithm
1. ES512 changed type from JSONWebKeySignatureAlgorithm to SignatureAlgorithm
1. Encrypt changed type from JSONWebKeyOperation to KeyOperation
1. Oct changed type from JSONWebKeyType to KeyType
1. P256 changed type from JSONWebKeyCurveName to KeyCurveName
1. P256K changed type from JSONWebKeyCurveName to KeyCurveName
1. P384 changed type from JSONWebKeyCurveName to KeyCurveName
1. P521 changed type from JSONWebKeyCurveName to KeyCurveName
1. PS256 changed type from JSONWebKeySignatureAlgorithm to SignatureAlgorithm
1. PS384 changed type from JSONWebKeySignatureAlgorithm to SignatureAlgorithm
1. PS512 changed type from JSONWebKeySignatureAlgorithm to SignatureAlgorithm
1. RS256 changed type from JSONWebKeySignatureAlgorithm to SignatureAlgorithm
1. RS384 changed type from JSONWebKeySignatureAlgorithm to SignatureAlgorithm
1. RS512 changed type from JSONWebKeySignatureAlgorithm to SignatureAlgorithm
1. RSA changed type from JSONWebKeyType to KeyType
1. RSA15 changed type from JSONWebKeyEncryptionAlgorithm to EncryptionAlgorithm
1. RSAHSM changed type from JSONWebKeyType to KeyType
1. RSAOAEP changed type from JSONWebKeyEncryptionAlgorithm to EncryptionAlgorithm
1. RSAOAEP256 changed type from JSONWebKeyEncryptionAlgorithm to EncryptionAlgorithm
1. RSNULL changed type from JSONWebKeySignatureAlgorithm to SignatureAlgorithm
1. Sign changed type from JSONWebKeyOperation to KeyOperation
1. UnwrapKey changed type from JSONWebKeyOperation to KeyOperation
1. Verify changed type from JSONWebKeyOperation to KeyOperation
1. WrapKey changed type from JSONWebKeyOperation to KeyOperation

### Struct Fields

1. JSONWebKey.Crv changed type from JSONWebKeyCurveName to KeyCurveName
1. JSONWebKey.Kty changed type from JSONWebKeyType to KeyType
1. KeyCreateParameters.Curve changed type from JSONWebKeyCurveName to KeyCurveName
1. KeyCreateParameters.KeyOps changed type from *[]JSONWebKeyOperation to *[]KeyOperation
1. KeyCreateParameters.Kty changed type from JSONWebKeyType to KeyType
1. KeyOperationsParameters.Algorithm changed type from JSONWebKeyEncryptionAlgorithm to EncryptionAlgorithm
1. KeyProperties.Curve changed type from JSONWebKeyCurveName to KeyCurveName
1. KeyProperties.KeyType changed type from JSONWebKeyType to KeyType
1. KeySignParameters.Algorithm changed type from JSONWebKeySignatureAlgorithm to SignatureAlgorithm
1. KeyUpdateParameters.KeyOps changed type from *[]JSONWebKeyOperation to *[]KeyOperation
1. KeyVerifyParameters.Algorithm changed type from JSONWebKeySignatureAlgorithm to SignatureAlgorithm

## New Content

### New Funcs

1. PossibleEncryptionAlgorithmValues() []EncryptionAlgorithm
1. PossibleKeyCurveNameValues() []KeyCurveName
1. PossibleKeyOperationValues() []KeyOperation
1. PossibleKeyTypeValues() []KeyType
1. PossibleSignatureAlgorithmValues() []SignatureAlgorithm
