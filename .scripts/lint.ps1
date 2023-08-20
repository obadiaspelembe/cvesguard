$goExecutable = "go.exe"
$goFile = "internal\main.go"

$command = "$goExecutable run $goFile lint -r cves-report.example.json -p policy.example.yaml"

Invoke-Expression $command