## deliverables
go-command

## summary
Using cobra, which is often used to create a CLI in go　and using viper around the settings.

## add a command
Implement each command as a subcommand

### example

### user
`1. INSERT: bookCmd book --mode insert --table account_user --accountname Kayle --mailaddress kayle@gmail.com`

`2. UPDATE: bookCmd book --mode update --table account_user --accountid 10 --accountname Kayle --mailaddress kayle1013@gmail.com`

`3. DELETE: bookCmd book --mode delete --table account_user --accountid 10`

`4. DELETE ALL: bookCmd book --mode deleteall --table account_user`
`
### book
`1. INSERT: bookCmd book --mode insert --table book --accountid 10 --title Harry --author JKR --status 1 --note so much fun!!`

`2. UPDATE: bookCmd book --mode update --table book --accountid 10 --title HarryPotter --author JKR --status 2 --note so much fun!!`

`3. DELETE: bookCmd book --mode delete --table book --accountid 10 --title HarryPotter`

`4. DELETE ALL: bookCmd book --mode deleteall --table book`


### reference
cobra document : https://umarcor.github.io/cobra/# go-command
