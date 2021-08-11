# aura
*aura* is a command line utility for  non-interactive download of files from the Web.  It supports HTTP, HTTPS protocols, as well as retrieval through HTTP proxies.

<br>


## ✨ Features

- Just one command to rule them all.
- Written in Go, and works for SPEED
- Distraction Free Downloads, simple workflow
- A simpler replacement to wget

<br>

## ⚡️ Installation
### Method 1

**For Linux users:**

Execute the following command as **Root User** (sudo su):

```bash
# Bash Users

curl https://raw.githubusercontent.com/arghyagod-coder/aura/master/linux/bash_install.sh > aura_install.sh
chmod u+x ./aura_install.sh
bash ./aura_install.sh

# Fish Users

curl https://raw.githubusercontent.com/arghyagod-coder/aura/master/linux/fish_install.sh > aura_install.sh
chmod u+x ./aura_install.sh
fish ./aura_install.sh

# Zsh Users

curl https://raw.githubusercontent.com/arghyagod-coder/aura/master/linux/zsh_install.sh > aura_install.sh
chmod u+x ./aura_install.sh
sh ./aura_install.sh
```


**For MacOS users:**

Execute the following command as **Root User** (sudo su):

```bash
curl https://raw.githubusercontent.com/arghyagod-coder/aura/master/mac_install.sh > aura_install.sh
chmod +x ./aura_install.sh
bash ./aura_install.sh
```

**For Windows users:**

Open Powershell **as Admin** and execute the following command:
```powershell
Set-ExecutionPolicy Bypass -Scope Process -Force; (Invoke-WebRequest -Uri https://raw.githubusercontent.com/arghyagod-coder/aura/master/windows_install.ps1 -UseBasicParsing).Content | powershell -
```

To verify the installation of *aura*, open a new shell and execute `aura -v`. You should see output like this:
```
aura 1.2.0

Version: 1.2.0
```
If the output isn't something like this, you need to repeat the above steps carefully.

<br>

## 💡 Usage
`aura` usage is simple

```bash
aura {url} --fn {filename}
```

--fn flag stands for filename. If I want a file named foo from https://foo.bar/foo.sh, I can execute:

```bash
aura https://foo.bar/foo.sh --fn foo
```

<br>

