# 🦎 FORGEZERO INSTALLATION ON YOUR SYSTEM — USER MANUAL

**For installation on various systems, please either use the pre-built binaries available in the releases, or build from source.**

The releases are always built for multiple architectures. However, due to technical constraints, some functionality may be limited in certain builds, as the project was originally written for and primarily developed on Linux. Nevertheless, you will always receive battle-tested core functionality. Please note that filesystem security features, sealing, and SBOM support may not work reliably in some environments. We apologize for the inconvenience!

## 🐧 For Arch Linux users, the process is simpler — install via:

```bash
yay -S forgezero-git
```

> Note: This installs a development version built from the source code on GitHub, tracking the main branch. This means you get the latest released version plus unreleased features that are still under active development.

## 🛠 To build from source manually:

```bash
git clone https://github.com/forgezero-cli/forgezero.git
cd forgezero
bash build.sh && sudo mv fz /usr/local/bin
```

## 📒 Additional installation option for ForgeZero

You may use the provided installer script, which was specifically written to eliminate the routine effort of compiling from source:

```bash
curl -fsSL https://raw.githubusercontent.com/forgezero-cli/ForgeZero/main/install.sh | sh
```

> Note: Ensure that the `curl` utility is installed on your system prior to running this command.

📥 If you encounter any errors during installation (which should not occur under normal circumstances), please report them to us via Issues: https://github.com/forgezero-cli/ForgeZero/issues

_If you are a contributor and intend to work with ForgeZero, please ensure that the following utilities are installed on your system for successful test execution (e.g., go test -v ./... and related commands): gcc, clang, nasm, fasm, and zig._

In the event that these utilities are not available, some tests may fail. If your goal is to test only a specific added module or component, you may limit testing to that particular package rather than running the full test suite.
