# OpenFaaS Bitbar Plugin

[![Build Status](https://travis-ci.org/johnmccabe/openfaas-bitbar.svg?branch=master)](https://travis-ci.org/johnmccabe/openfaas-bitbar)
[![Go Report Card](https://goreportcard.com/badge/github.com/johnmccabe/openfaas-bitbar)](https://goreportcard.com/report/github.com/johnmccabe/openfaas-bitbar)

![OpenFaaS Bitbar Plugin Gif](https://user-images.githubusercontent.com/83862/31579104-8fc7fc54-b126-11e7-9d50-000069534d16.gif)

## Installing

You can install and use the plugin as follows:

### BitBar

First, follow the [Bitbar Getting Started instructions](https://github.com/matryer/bitbar) to get BitBar running on your system. **Note:** make sure you create a new plugins directory when first installing, don't just use the Documents directory itself - try creating it in `~/Documents/bitbar_plugins`.

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