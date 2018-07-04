# OpenFaaS Bitbar Plugin

[![Build Status](https://travis-ci.org/johnmccabe/openfaas-bitbar.svg?branch=master)](https://travis-ci.org/johnmccabe/openfaas-bitbar)
[![Go Report Card](https://goreportcard.com/badge/github.com/johnmccabe/openfaas-bitbar)](https://goreportcard.com/report/github.com/johnmccabe/openfaas-bitbar)
[![OpenFaaS](https://img.shields.io/badge/openfaas-serverless-blue.svg)](https://www.openfaas.com)

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

The `openfaas-bitbar` plugin looks for the presense of a `~/.openfaas/config.yml` file, and if present it will use that to decide what gateways to connect to and use their configured auth. To ensure you can see/login to your gateway, make sure that you've run the `faas-cli login` command for each gateway you wish to see in the menubar.

## Updating

    brew upgrade openfaas-bitbar