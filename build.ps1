#Requires -Version 5.1
<#
.SYNOPSIS
    Build SpotiFLACNeo for Windows (.exe) via Wails.

.PARAMETER SkipFrontend
    Skip npm/pnpm install (only run wails build).

.PARAMETER Compress
    Compress the exe with UPX when available.

.PARAMETER Dev
    Run `wails dev` instead of a release build.

.EXAMPLE
    .\build.bat
    .\build.bat -Compress

NOTE
    If "running scripts is disabled", use build.bat or:
    powershell -ExecutionPolicy Bypass -File .\build.ps1
#>
[CmdletBinding()]
param(
    [switch]$SkipFrontend,
    [switch]$Compress,
    [switch]$Dev
)

$ErrorActionPreference = "Stop"
$Root = $PSScriptRoot
$Frontend = Join-Path $Root "frontend"
$OutputExe = Join-Path $Root "build\bin\SpotiFLACNeo.exe"

function Write-Step([string]$Message) {
    Write-Host ""
    Write-Host "==> $Message" -ForegroundColor Cyan
}

function Test-Command([string]$Name) {
    return [bool](Get-Command $Name -ErrorAction SilentlyContinue)
}

function Find-GoExecutable {
    $cmd = Get-Command go -ErrorAction SilentlyContinue
    if ($cmd) {
        return $cmd.Source
    }

    $candidates = @(
        if ($env:GOROOT) { Join-Path $env:GOROOT "bin\go.exe" }
        "C:\Program Files\Go\bin\go.exe"
        "$env:LOCALAPPDATA\Programs\Go\bin\go.exe"
        "$env:USERPROFILE\sdk\go*\bin\go.exe"
    ) | Where-Object { $_ }

    foreach ($path in $candidates) {
        $resolved = $null
        if ($path -like '*`**') {
            $resolved = Get-Item $path -ErrorAction SilentlyContinue | Select-Object -First 1 -ExpandProperty FullName
        }
        elseif (Test-Path $path) {
            $resolved = (Resolve-Path $path).Path
        }
        if ($resolved) {
            return $resolved
        }
    }
    return $null
}

function Ensure-GoInPath {
    $goExe = Find-GoExecutable
    if (-not $goExe) {
        Write-Host ""
        Write-Host "Go was not found in PATH." -ForegroundColor Red
        Write-Host ""
        Write-Host "Install Go 1.26+ (required by this project), then reopen PowerShell:" -ForegroundColor Yellow
        Write-Host "  winget install GoLang.Go" -ForegroundColor White
        Write-Host "  https://go.dev/dl/  (Windows amd64 MSI)" -ForegroundColor White
        Write-Host ""
        Write-Host "After install, verify:" -ForegroundColor Yellow
        Write-Host "  go version" -ForegroundColor White
        throw "Go is required (1.26+). https://go.dev/dl/"
    }

    $goBin = Split-Path $goExe -Parent
    if ($env:Path -notlike "*$goBin*") {
        $env:Path = "$goBin;$env:Path"
    }

    $versionLine = & $goExe version 2>&1
    Write-Host "Using: $versionLine" -ForegroundColor DarkGray

    if ($versionLine -match 'go(\d+)\.(\d+)') {
        $major = [int]$Matches[1]
        $minor = [int]$Matches[2]
        if ($major -lt 1 -or ($major -eq 1 -and $minor -lt 26)) {
            Write-Warning "Project needs Go 1.26+. Installed: go$major.$minor"
        }
    }
}

function Ensure-GoModule {
    Push-Location $Root
    try {
        if (-not (Test-Path (Join-Path $Root "go.sum"))) {
            Write-Step "go mod download"
            go mod download
        }
    }
    finally {
        Pop-Location
    }
}

function Ensure-Wails {
    if (-not (Test-Command "wails")) {
        Write-Step "Installing Wails CLI"
        go install github.com/wailsapp/wails/v2/cmd/wails@latest
        $goBin = Join-Path (go env GOPATH) "bin"
        if ($env:Path -notlike "*$goBin*") {
            $env:Path = "$goBin;$env:Path"
        }
    }
    if (-not (Test-Command "wails")) {
        throw "Wails CLI not found. Add GOPATH\bin to PATH and retry."
    }
}

function Ensure-Frontend {
    if (-not (Test-Command "node")) {
        throw "Node.js is required (20+, 24 LTS recommended). https://nodejs.org/"
    }
    Push-Location $Frontend
    try {
        if (-not $SkipFrontend) {
            Write-Step "Installing frontend dependencies"
            if (Test-Command "pnpm") {
                pnpm install
                pnpm run generate-icon
            }
            else {
                npm install
                npm run generate-icon
            }
        }
        if (-not (Test-Path "dist\index.html")) {
            Write-Step "Building frontend (vite)"
            if (Test-Command "pnpm") { pnpm run build } else { npm run build }
        }
    }
    finally {
        Pop-Location
    }
}

Write-Host "SpotiFLACNeo - Windows build" -ForegroundColor Green
Write-Host "Root: $Root"

Ensure-GoInPath
Ensure-GoModule
Ensure-Wails
Ensure-Frontend

Push-Location $Root
try {
    if ($Dev) {
        Write-Step "Starting wails dev"
        wails dev
        return
    }

    Write-Step "Building application (windows/amd64)"
    wails build -platform windows/amd64 -clean

    if (-not (Test-Path $OutputExe)) {
        throw "Build finished but output not found: $OutputExe"
    }

    if ($Compress) {
        if (Test-Command "upx") {
            Write-Step "Compressing with UPX"
            upx --best --lzma $OutputExe
        }
        else {
            Write-Warning "UPX not found. Install with: choco install upx"
        }
    }

    $sizeMb = [math]::Round((Get-Item $OutputExe).Length / 1MB, 2)
    Write-Host ""
    Write-Host "Done: $OutputExe ($sizeMb MB)" -ForegroundColor Green
}
finally {
    Pop-Location
}
