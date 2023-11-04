# SyCrypt
It's a password hashing package to create a strong military grade passwords.


## Motivation
To stop using ridiculous & outdated libraries for password hashing with salt and crap!


## How-to Guide
You can instantiate a creation of new `Credentials` data object using:

```go
package main

import (
	"encoding/json"

	"github.com/syniol/sycrypt"
)

func main() {
	credentials, err := sycrypt.NewCredential("johnspassword1")
	if err != nil {
		panic(err)
	}

	// prints hashed password with public and private key
	result, err := json.Marshal(credentials)
	if err != nil {
		panic(err)
	}

	println(string(result))
}
```

This will execution will output a credentials JSON object where public, private key, and hashed 
password is stored in `base64` before `hex` encode.

```json
{
  "key":      "MmQyZDJkMmQyZDQyNDU0NzQ5NGUyMDUwNTU0MjRjNDk0MzIwNGI0NTU5MmQyZDJkMmQyZDBhNGQ0MzZmNzc0MjUxNTk0NDRiMzI1Njc3NDE3OTQ1NDE2MTVhNTc2NjZlNmUzNzQxNmU0YzQ0NDY2MTRmMzM0NzMwNDQ2OTdhNzA1NzRhNDY2MjdhNzI3NjcyNDIzMTRhNjE3YTc1NzM0NjY5MzQ2ODY4Mzg2NzNkMGEyZDJkMmQyZDJkNDU0ZTQ0MjA1MDU1NDI0YzQ5NDMyMDRiNDU1OTJkMmQyZDJkMmQwYQ==",
  "hashedPassword": "ZDRkM2QzMWQyN2JjZTYyZTRjODI2MTFkYmZjMzk0YmIzNTI4MmRhODMwYTBhMWI3NjBiZjhkZjQzOGZjZDViOTViMGI4ZDBjMTY5ZjlhMzAxNGIwMGY4ZDVlYTMyMWE5MDAzNzVhNGE0MWZhMTFhZDViNjEwYTg0YTk2ZTAyMDI="
}
```

This could be stored in database for verification. Please see code below for verification.

```go

package main

import (
	"github.com/syniol/sycrypt"
)

func main() {
	// populate hashed password, public key and private keys from initial creation
	// this could be stored in database to use in verification process 'VerifyPassword'
	credentials := &sycrypt.Credential{
		Key:      "MmQyZDJkMmQyZDQyNDU0NzQ5NGUyMDUwNTU0MjRjNDk0MzIwNGI0NTU5MmQyZDJkMmQyZDBhNGQ0MzZmNzc0MjUxNTk0NDRiMzI1Njc3NDE3OTQ1NDE2MTVhNTc2NjZlNmUzNzQxNmU0YzQ0NDY2MTRmMzM0NzMwNDQ2OTdhNzA1NzRhNDY2MjdhNzI3NjcyNDIzMTRhNjE3YTc1NzM0NjY5MzQ2ODY4Mzg2NzNkMGEyZDJkMmQyZDJkNDU0ZTQ0MjA1MDU1NDI0YzQ5NDMyMDRiNDU1OTJkMmQyZDJkMmQwYQ==",
		HashedPassword: "ZDRkM2QzMWQyN2JjZTYyZTRjODI2MTFkYmZjMzk0YmIzNTI4MmRhODMwYTBhMWI3NjBiZjhkZjQzOGZjZDViOTViMGI4ZDBjMTY5ZjlhMzAxNGIwMGY4ZDVlYTMyMWE5MDAzNzVhNGE0MWZhMTFhZDViNjEwYTg0YTk2ZTAyMDI=",
	}

	isVerified := credentials.VerifyPassword("johnspassword1")
	
	// prints if password is verified
	if isVerified == true {
		println("verified")
		
		return
    }
	
	println("not verified")
}
```

> __Please refer at test file for more examples and test cases.__


#### Credits
Engineered by [Hadi Tajallaei](mailto:hadi@syniol.com) in London.

Copyright &copy; 2023 Syniol Limited. All rights Reserved.
