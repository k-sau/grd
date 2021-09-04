# grd

### Get root domains
This uses *https://github.com/joeguo/tldextract* go module.
It takes subdomains from stdin and prints the root domains.

### Installation
```
GO111MODULE=on go get github.com/k-sau/grd
```

### Usage
```
bbrf scope in -p verizonmedia | grd

yahoo.com
flurry.com
makers.com
buildseries.com
builtbygirls.com
ryot.org
engadget.com
techcrunch.com
aol.com
rivals.com
protrade.com
verizondigitalmedia.com
vdms.io
aolpublishers.com
netscape.com
compuserve.com
vzbuilders.com
oath.cloud
yahoo.cloud
downlynk.com
uplynk.net
uplynk.com

```
