
Write-Host "Downloading aura..."

$url = "https://github.com/arghyagod-coder/aura/releases/latest/download/aura-windows-amd64.exe"

$dir = $env:USERPROFILE + "\.aura"
$filepath = $env:USERPROFILE + "\.aura\aura.exe"

[System.IO.Directory]::CreateDirectory($dir)
(Invoke-WebRequest -Uri $url -OutFile $filepath)

Write-Host "Adding aura to PATH..."
[Environment]::SetEnvironmentVariable(
    "Path",
    [Environment]::GetEnvironmentVariable("Path", [EnvironmentVariableTarget]::Machine) + ";"+$dir,
    [EnvironmentVariableTarget]::Machine)

Write-Host 'aura installation is successfull!'
Write-Host "You need to restart your shell to use aura."
