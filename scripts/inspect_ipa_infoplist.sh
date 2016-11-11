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
#

IPA_PATH=${1:?ipa path is missing}
if [ ! -f "$IPA_PATH" ];then
    echo "No such file ${IPA_PATH}"
    exit 1
fi
tmp=`mktemp`
tmppayload=`mktemp -d`

unzip "${IPA_PATH}" Payload/**/*Info.plist -d $tmppayload > /dev/null

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
