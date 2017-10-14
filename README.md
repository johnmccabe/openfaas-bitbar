# EARLY PRE-RELEASE

# OpenFaaS Bitbar Plugin

[![Build Status](https://travis-ci.org/johnmccabe/openfaas-bitbar.svg?branch=master)](https://travis-ci.org/johnmccabe/openfaas-bitbar)
[![Go Report Card](https://goreportcard.com/badge/github.com/johnmccabe/openfaas-bitbar)](https://goreportcard.com/report/github.com/johnmccabe/openfaas-bitbar)

Coming soon to a Mac near you...

![OpenFaaS Bitbar Plugin Gif](assets/screencap.gif)

## Installing

You can install and use the plugin as follows (be aware that this is still under development so expect frequent changes, and please do raise issues):

### BitBar

First, follow the Bitbar Getting Started instructions to get BitBar running on your system. **Note:** make sure you create a new plugins directory when first installing, don't just use the Documents directory itself - try creating ~/Documents/bitbar_plugins.
 - https://github.com/matryer/bitbar


### OpenFaaS Plugin

Install using the provider Homebrew tap.

    $ brew tap johnmccabe/openfaas-bitbar
    $ brew install openfaas-bitbar


Install a symlink with the specified refresh interval to the plugin directory.

    $ openfaas-bitbar install

**Note:** When installing for the fist time you will need to manually restart the BitBar App, or select `Preferences/Refresh all` from its dropdowns. 

## Configuring

By default the plugin will point at an OpenFaaS stack running on `localhost`, if you wish to point at a different location you can create/update the config using the CLI or editing the file yourself:

**CLI**

    $ openfaas-bitbar config

**File**

The config, if not using the default, is stored in `~/.openfaas/config.yaml`, here is the YAML corresponding to the CLI example above:

```yaml
stacks:
- name: Civo OpenFaaS
  gateway: http://xxx.xxx.xxx.xxx:8080/
  prometheus: http://xxx.xxx.xxx.xxx:9090/
```

## Updating

    brew upgrade openfaas-bitbar