#!/bin/bash
#
# name:
#   inspect_ipa_infoplist.sh
#
# description:
#   Find and Print all Info.plist files.
#
# author:
#   Toshihiro Suzuki (@toshi0383)
#
# license:
#   MIT
#
# parameters:
#   - 1 IPA_PATH: path to ipa
#   - 2 APP_ONLY (optional): If set, prints only an App's main Info.plist.
#

set -eo pipefail

main() {
    IPA_PATH=${1:?ipa path is missing}
    APP_ONLY=${2}
    if [ ! -f "$IPA_PATH" ];then
        echo "No such file ${IPA_PATH}"
        exit 1
    fi
    tmp=`mktemp`
    tmppayload=`mktemp -d`

    if [ $APP_ONLY ];
    then
        unzip "${IPA_PATH}" Payload/*.app/Info.plist -d $tmppayload > /dev/null
    else
        unzip "${IPA_PATH}" Payload/**/Info.plist -x *storyboardc/* -d $tmppayload > /dev/null
    fi

    find $tmppayload -name "*Info.plist" | \
        grep -v storyboardc > $tmp

    while read line;
    do
        echo "### `echo ${line} | sed -e 's?.*Payload/??g'`"
        plutil -p $line 2> /dev/null
        echo ""
    done < $tmp
    rm $tmp
    rm -rf $tmppayload
}
