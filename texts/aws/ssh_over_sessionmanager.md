# SSH over Session Manager

Connect to private instances over `ssh` with AWS Session Manager.

1. [Install session manager plugin](https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-working-with-install-plugin.html)
2. Add the following to the `ssh` config file (usually `~/.ssh/config`)
```
# SSH over Session Manager
Host i-* mi-*
  ProxyCommand sh -c "aws ssm start-session --target %h --document-name AWS-StartSSHSession --parameters 'portNumber=%p'"
```

Example with Leapp setup:

```
AWS_PROFILE=project-stage ssh -i /path/to/file.pem ec2-user@i-0b22a22eec53b9321
```
