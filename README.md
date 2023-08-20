# CVES-Guard

cvesguard (Short form of Docker Scout CVES Guard) is a tool to police docker images cves vulnerabilities through a policy manifest in your pipeline. This tool helps CI/CD engineers to determine whether a Docker Image can be deployed after docker scout command execution.

## How to use cvesguard

cvesguard is based on docker scout command with sarif format to json. You can run the command as per example:


```bash
docker scout cves hello-world --format sarif --output cves-report.json
```
There're few commands available for cvesguard tool.

### Lint
Validates if the policy and cves files are compliant with the schema .

``` bash
cvesguard lint --policy policy.yaml --cves-report cves-report.json

or

cvesguard lint -p policy.yaml -r cves-report.json

```

### Exec policy
Checks the specified policy in the manifest.

```bash
cvesguard exec --policy policy.yaml --cves-report cves-report.json

or

cvesguard exec -p policy.yaml -r cves-report.json

```

## Configuration

The policy manifest file contains the configurations details of how apply command should be executed.

Example:

policy.yaml

```
---
version: v1.0.0 
spec:
  config:
    vulnerability:
      critical: 0
      high: 0
      medium: 100
      low: 2
    packages:
      - name: log4j
        action: ignore
        severity:
          - critical
          - high
```

## Contributing

Contributions are welcome! Please do not hesitate to submit a Pull Request.

## License

This project is licensed under the MIT License.