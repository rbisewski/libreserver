# libreserver

Simple open source webserver, written in golang. The purpose of this repo
is to be as plain as possible and stick to the barebones of concept.

# Building

Run the make as you would a typical project.

```
make
```

After doing so, consider reading the usage or installation section.

# Installation

To install the server, run the following command:

```
make install
```

To remove the server from the system, use the following command:

```
make uninstall
```

Afterwards the binary and any associated logs will be deleted.

# Usage

[insert content here]

# TODOs

* Support more addresses than just localhost / 127.0.0.1
* Add functions to properly handle certificates.
* Get this to serve HTTPS and HTTP2 content.
* Add unit tests
