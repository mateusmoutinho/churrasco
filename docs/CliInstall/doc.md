# CliInstall

`churrasco` is a single static binary: no runtime, no dependencies. Pick your platform,
paste the block, done. Go 1.25+ is needed only to build it from source.

**macOS (Apple Silicon)**

```bash
curl -sL https://github.com/mateusmoutinho/churrasco/releases/latest/download/macarm64.bin -o churrasco && chmod +x churrasco && sudo mv churrasco /usr/local/bin/ && churrasco version
```

**macOS (Intel)**

```bash
curl -sL https://github.com/mateusmoutinho/churrasco/releases/latest/download/mac86.bin -o churrasco && chmod +x churrasco && sudo mv churrasco /usr/local/bin/ && churrasco version
```

**Linux (amd64)**

```bash
curl -sL https://github.com/mateusmoutinho/churrasco/releases/latest/download/linux86.out -o churrasco && chmod +x churrasco && sudo mv churrasco /usr/local/bin/ && churrasco version
```

**Linux (arm64)**

```bash
curl -sL https://github.com/mateusmoutinho/churrasco/releases/latest/download/linuxarm64.out -o churrasco && chmod +x churrasco && sudo mv churrasco /usr/local/bin/ && churrasco version
```

**Linux (32-bit)**

```bash
curl -sL https://github.com/mateusmoutinho/churrasco/releases/latest/download/linuxi32.out -o churrasco && chmod +x churrasco && sudo mv churrasco /usr/local/bin/ && churrasco version
```

**Windows (64-bit)** — PowerShell:

```powershell
$dir="$HOME\.local\bin"; New-Item -ItemType Directory -Force -Path $dir | Out-Null; curl.exe -sL https://github.com/mateusmoutinho/churrasco/releases/latest/download/windows86.exe -o "$dir\churrasco.exe"; [Environment]::SetEnvironmentVariable('PATH', [Environment]::GetEnvironmentVariable('PATH','User') + ";$dir", 'User')
```

**Windows (32-bit)** — PowerShell:

```powershell
$dir="$HOME\.local\bin"; New-Item -ItemType Directory -Force -Path $dir | Out-Null; curl.exe -sL https://github.com/mateusmoutinho/churrasco/releases/latest/download/windowsi32.exe -o "$dir\churrasco.exe"; [Environment]::SetEnvironmentVariable('PATH', [Environment]::GetEnvironmentVariable('PATH','User') + ";$dir", 'User')
```

**From a checkout** — needs Go 1.25+:

```bash
go build -o churrasco ./cmd/main && sudo mv churrasco /usr/local/bin/ && churrasco version
```

The released binaries are the ones `agnos compile --target all` builds and `agnos publish`
uploads. `churrasco version` prints the `version` of `AgnosConfig/project.yaml`,
`churrasco help` every command — each one is listed in [Commands](../Commands/doc.md).
