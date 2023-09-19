$user = $env:USERNAME
$binFolder = "C:\Users\$user\bin"
New-Item -ItemType Directory -Path $binFolder -Force
$currentPath = [System.Environment]::GetEnvironmentVariable("PATH", [System.EnvironmentVariableTarget]::User)
if ($currentPath -split ";" -contains $binFolder) {
    Write-Host "Path could not be updated"
} else {
    $newPath = "$currentPath;$binFolder"
    [System.Environment]::SetEnvironmentVariable("PATH", $newPath, [System.EnvironmentVariableTarget]::User)
    Write-Host "Path has been updated"
}
$destinationPath = Join-Path -Path $binFolder -ChildPath "raminfo.exe"
Invoke-WebRequest -Uri 'https://github.com/RealLava/raminfo/raw/main/raminfo.exe' -OutFile $destinationPath