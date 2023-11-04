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

	// prints hashed password
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
  "publicKey":      "MmQyZDJkMmQyZDQyNDU0NzQ5NGUyMDUwNTU0MjRjNDk0MzIwNGI0NTU5MmQyZDJkMmQyZDBhNGQ0MzZmNzc0MjUxNTk0NDRiMzI1Njc3NDE3OTQ1NDE2MTVhNTc2NjZlNmUzNzQxNmU0YzQ0NDY2MTRmMzM0NzMwNDQ2OTdhNzA1NzRhNDY2MjdhNzI3NjcyNDIzMTRhNjE3YTc1NzM0NjY5MzQ2ODY4Mzg2NzNkMGEyZDJkMmQyZDJkNDU0ZTQ0MjA1MDU1NDI0YzQ5NDMyMDRiNDU1OTJkMmQyZDJkMmQwYQ==",
  "privateKey":     "MmQyZDJkMmQyZDQyNDU0NzQ5NGUyMDUwNTI0OTU2NDE1NDQ1MjA0YjQ1NTkyZDJkMmQyZDJkMGE0ZDQzMzQ0MzQxNTE0MTc3NDI1MTU5NDQ0YjMyNTY3NzQyNDM0OTQ1NDk0ZTMzNDY2NTU3NzM0ZjMzNTg1MDM2MzkzMzZmMzk0ZTY0NDE0NDM5NDMzODMwMzM2YzY1Njk3MTZjMmI2MTU4NTk3ODU0NDM2MjQyNDg1MDJmNzQ2YjBhMmQyZDJkMmQyZDQ1NGU0NDIwNTA1MjQ5NTY0MTU0NDUyMDRiNDU1OTJkMmQyZDJkMmQwYQ==",
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
	credentials := &sycrypt.Credential{
		PublicKey:      "MmQyZDJkMmQyZDQyNDU0NzQ5NGUyMDUwNTU0MjRjNDk0MzIwNGI0NTU5MmQyZDJkMmQyZDBhNGQ0MzZmNzc0MjUxNTk0NDRiMzI1Njc3NDE3OTQ1NDE2MTVhNTc2NjZlNmUzNzQxNmU0YzQ0NDY2MTRmMzM0NzMwNDQ2OTdhNzA1NzRhNDY2MjdhNzI3NjcyNDIzMTRhNjE3YTc1NzM0NjY5MzQ2ODY4Mzg2NzNkMGEyZDJkMmQyZDJkNDU0ZTQ0MjA1MDU1NDI0YzQ5NDMyMDRiNDU1OTJkMmQyZDJkMmQwYQ==",
		PrivateKey:     "MmQyZDJkMmQyZDQyNDU0NzQ5NGUyMDUwNTI0OTU2NDE1NDQ1MjA0YjQ1NTkyZDJkMmQyZDJkMGE0ZDQzMzQ0MzQxNTE0MTc3NDI1MTU5NDQ0YjMyNTY3NzQyNDM0OTQ1NDk0ZTMzNDY2NTU3NzM0ZjMzNTg1MDM2MzkzMzZmMzk0ZTY0NDE0NDM5NDMzODMwMzM2YzY1Njk3MTZjMmI2MTU4NTk3ODU0NDM2MjQyNDg1MDJmNzQ2YjBhMmQyZDJkMmQyZDQ1NGU0NDIwNTA1MjQ5NTY0MTU0NDUyMDRiNDU1OTJkMmQyZDJkMmQwYQ==",
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


## Golang 1.11.0 with Docker
You could experiment with this package using Docker and Golang 1.11.0 image.

```bash
  docker run -it --name sycrypt-golang110 -v $(pwd):/var/local/sycrypt -w /var/local/sycrypt  --rm golang:1.11.0-alpine sh
```


#### Credits
Copyright &copy; 2023 Syniol Limited. All rights Reserved.

Engineered by [Hadi Tajallaei](mailto:hadi@syniol.com) with ☕ & 💛 in London.
