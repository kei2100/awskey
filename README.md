awskey
====

Print AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY

that searches environment variables, shared credential file, and EC2 Instance Roles

### Install
[release page](https://github.com/kei2100/awskey/releases)

or 

```
go get -u github.com/kei2100/awskey
```

### Example
```
$ awskey -n 
AWS_ACCESS_KEY_ID=... AWS_SECRET_ACCESS_KEY=...

# specify flags
$ docker run $(awskey --flag=-e) awscli s3 ls
$ circleci local execute $(awskey --flag=-e) --job=test
```
