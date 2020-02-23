## 0.17.0

* Bind .terraformrc file when running Terraform

## 0.16.0

* Go Modules, and Go 1.13

## 0.15.1

* The `--entrypoint`/`-e` flag for `lucli generic` should actually, ya know, DO SOMETHING

## 0.15.0

* Generic container runner, for quick $PWD:/tmp/workdir

## 0.14.1

* Fix panic when running `lucli terraform` with no args

## 0.14.0

* Initial support for custom Terraform plugins

## 0.13.1

* Use mapcrafter/mapcrafter:113

## 0.13.0

* Mapcrafter

## 0.12.0

* S3 Explorer, from https://github.com/tinyzimmer/s3explorer
* Global -p / --aws-profile flag

## 0.11.1

* Sportsball, but with UK times

## 0.11.0

* SPORTSBALL!

## 0.10.0

* Dry - https://github.com/moncho/dry

## 0.9.0

* AWS CLI / AWS Shell, with lucli awscli

## 0.8.0

* Initial version of keybase for AWS credentials

## 0.7.2

* Build with latest Cali, 0.2.0

## 0.7.1

* Correctly use vault-version flag

## 0.7.0

* Initial PoC Vault

## 0.6.2

* lucli version --check-update=false, for just displaying current version
(Because Travis is inconsistent with whether it can do this or not)

## 0.6.1

* lucli version now says if you're already on the latest version

## 0.6.0

* lucli update

## 0.5.0

* Build with Go 1.10
* lucli github-release, borrowed from staticli, to release new versions of lucli to github
* Switched to Github for releases. Sorry, BinTray
* Terraform now supports arbitrary versions, defaulting to latest
* Other misc stuff

## 0.4.0

* Generic wrapper around images which expect an X $DISPLAY

## 0.3.1

* Version will only show commit (and now repo!) when in verbose mode

## 0.3.0

* (vim) UID and GID mapping from the host's user to the container's user

## 0.2.1

* Darwin build is now amd64, rather than 386. Added a linux amd64 build too.

## 0.2.0

* Versioning: lucli version command now displays which version you're running.
  Precursor to auto-updating

* Minor bugfixes & tweaks to skybet/cali

* Bump Terraform

## 0.1.2

* If XQuartz has alredy started, no need to start it again

## 0.1.1

* Iceweasel --> Firefox

## 0.1.0

* Actually started caring about Changelog
* Initial version of Iceweasel PoC
