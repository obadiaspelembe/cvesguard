
$reportFileName = "report.json"

$dockerImageName = "hello-world"

$command = "docker scout cves $dockerImageName --format sarif --output $reportFileName"

Invoke-Expression $command