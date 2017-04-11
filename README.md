# ipanema 🏖

ipanema is a great partner when you're debugging ipa device install errors.
ipanema analyzes and prints useful information from *.ipa file.

# Usage

```
Usage: ipanema [OPTIONS] path-to.ipa

OPTIONS:
  -A    Do both what -E and -I would do.
  -E    Find, Decrypt and Print all embedded.mobileprovision files.
  -I    Find and Print all Info.plist files.
  -v    Print version number
```

## embedded.mobileprovision files. (-E)
```
$ ipanema -E path/to/your.ipa | head
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>AppIDName</key>
    <string>XC your app id</string>
    <key>ApplicationIdentifierPrefix</key>
    <array>
    <string>XXXXXXXXXX</string>
    </array>
```

## Info.plist (-I)
```
$ ipanema -I path/to/your.ipa | head

### App.app/Frameworks/APIKit.framework/Info.plist
{
  "CFBundleName" => "APIKit"
  "DTXcode" => "0730"
  "DTSDKName" => "appletvos9.2"
  "NSHumanReadableCopyright" => "Copyright © 2015 Yosuke Ishikawa. All rights reserved."
  "DTSDKBuild" => "13Y227"
  "CFBundleDevelopmentRegion" => "en"
  "CFBundleVersion" => "1"
```

# Prequisite

- OSX/macOS

# Install

- Download the latest binary from [releases page](https://github.com/toshi0383/ipanema/releases).
- `chmod +x ~/Download/ipanema`
- `cp ~/Download/ipanema where-ever-you-want/`

# License
MIT
