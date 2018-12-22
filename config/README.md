## Configuration Files

The default location for the configuration file is $GGEROOT/gge.json. If you want to change this location, you will need to change the code in this package.

By default, the code looks for an environmental variable named $GGEROOT. If it is not there it is assigned the default of '/home/user/.config/gge'.

If $GGEROOT/gge.json does not exist, configuration defaults are set by ggconf.defaultConfiguration().

In the future I will put more work into this package so that it can verify which settings are actually set and which are not in the config, and assign defaults to those that are blank. For now, if the config file is proper json, it assumes it is good.

The example configuration file in this directory will show not only my default expectations, but also what you will get if you fall back to defaultConfiguration().
