$goExecutable = "go.exe"
$goFile = "internal\main.go"

$command = "$goExecutable run $goFile exec -r cves-report.example.json -p policy.example.yaml"

Invoke-Expression $command