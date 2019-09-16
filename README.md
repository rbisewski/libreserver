# libreserver

Simple open source webserver, written in golang. The purpose of this repo
is to be as plain as possible and stick to the barebones of concept.

# Building

Run the make as you would a typical project.

```bash
make
```

After doing so, consider reading the usage or installation section.

# Installation

To install the server, run the following command:

```bash
make install
```

To remove the server from the system, use the following command:

```bash
make uninstall
```

Afterwards the binary and any associated logs will be deleted.

# Usage

Once installed, the binary itself can be ran via commandline like so:

```bash
libreserver
```

Alternatively, consider the include systemd service, which can be started and enabled like so:

```bash
sudo systemctl enable libreserver
sudo systemctl start libreserver
```

# TODOs

* Support more addresses than just localhost / 127.0.0.1
* Add functions to properly handle certificates.
* Get this to serve HTTPS and HTTP2 content.
* Add unit tests
