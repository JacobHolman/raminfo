# raminfo
Linux program to display RAM info; written in Go and now also available on Windows.

<img src="https://raw.githubusercontent.com/JacobHolman/raminfo/main/preview.png">

## Install

**Linux:**

```bash
bash <(curl -s https://raw.githubusercontent.com/JacobHolman/raminfo/main/install.sh)
```

**Windows:**

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -Command "iwr 'https://raw.githubusercontent.com/JacobHolman/raminfo/refs/heads/main/install.ps1' -UseBasicParsing | iex"
```

## Uninstall

**Linux:**

```bash
sudo rm /usr/bin/raminfo
```

**Windows:**

```powershell
del "%USERPROFILE%\bin\raminfo.exe"
```
