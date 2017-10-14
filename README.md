# OpenFaaS Bitbar Plugin

[![Build Status](https://travis-ci.org/johnmccabe/openfaas-bitbar.svg?branch=master)](https://travis-ci.org/johnmccabe/openfaas-bitbar)
[![Go Report Card](https://goreportcard.com/badge/github.com/johnmccabe/openfaas-bitbar)](https://goreportcard.com/report/github.com/johnmccabe/openfaas-bitbar)

Coming soon to a Mac near you...

![OpenFaaS Bitbar Plugin Gif](assets/screencap.gif)

## Trying out the plugin

You can install and use the plugin as follows (be aware that this is still under development so expect frequent changes, and please do raise issues):

First, follow the Bitbar Getting Started instructions to get BitBar running on your system. Make sure you create a new plugins directory when first installing, don't just use the Documents directory itself - try creating ~/Documents/bitbar_plugins.
 - https://github.com/matryer/bitbar

When BitBar is running you can now install the OpenFaaS plugin.
```
brew tap johnmccabe/openfaas-bitbar
brew install openfaas-bitbar
```

Next create a config file (this step will be automated in future releases).
```
mkdir -p ~/.openfaas
curl -o ~/.openfaas/config.toml https://raw.githubusercontent.com/johnmccabe/openfaas-bitbar/master/.openfaas/config.toml
```
This default config file points at OpenFaaS/Prometheus instances running on localhost. Update the endpoints here to point at your own OpenFaaS instance (support for auth coming soon).

Now you can make the plugin available to BitBar.
```
openfaas-bitbar install
```
Recommend you use the default refresh rate of `30s`.

Now you either restart BitBar, or select `Prefernces/Refresh all` from its dropdowns.