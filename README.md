# Command Line Password Generator In Golang


## Build
1. First you have to install the golang on the system.
```bash
go version # check the system has golang 
```

2. Clone the project.
```bash
git clone github.com/hashen47/go_password_generator
```

3. change directory to project folder and then run `make build` .
```bash
cd go_password_generator
make build
```

4. Then bin folder have to binary file called *password_gen*.
```bash
./bin/password_gen -h # show the help
```

## How to
- List the help menu.
```bash
password_gen -h
# generate custom password according to the given options
# 
# Usage:
#   passwd_gen [len] [-D digit] [-L lower] [-U upper] [-S special] [flags]
# 
# Flags:
#   -D, --digit string     minimum digit character count
#   -h, --help             help for passwd_gen
#       --len string       password length
#   -L, --lower string     minimum lowercase character count
#   -S, --special string   minimum specialcase character count
#   -U, --upper string     minimum uppercase character count
```

- Generate password with default options (--len 8 -L 1 -U 1 -D 1 -S 1).
```bash
password_gen
# m-aZB_u1 # example password 
```

- You can give password length, lowercase, uppercase, digits and special character minimum counts in that password.
```bash
password_gen --len 16 -U 5 -L 2 -D 5 -S 3  # generate a password length is 16, uppercase minimum letter count 5, lowercase 2, digits 5 and special character count 3.
# X2k]4Zf5$5#nI0TD # example password
```
