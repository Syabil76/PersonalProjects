# File Encryption
  My second golang CLI tool made to encrypt and decrypt files given a password
  The logic is very simple: get file --> validate --get password --> double entry --> encrypt
  vice versa

## Function
  There are only 3 current commands with that being of 
```
  -en: ENCRYPTS a file using sha1 and aes algorithms built in go
  -d: DECRYPTS a file using key
  --h: SENDS help file
```
## Known Issues
  As of right now, the password function doesn't work, the program will encrypt and decrypt files without authentication.
  This is potenntally due to the **term** package imported which is supposed to read the user input.
  insufficient comments added as of 31/7/24- makes code hard to read.

## Potential updates
  Add more functions such as:
  ```
    -ca: CHOOSES ALGORITHM to ENCRYPT file
    e.g -en img.png -ca sha256
  ```
  
