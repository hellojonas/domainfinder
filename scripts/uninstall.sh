#!/usr/bin/env bash

if [[ $(basename "$PWD") != "scripts" ]] 
then
  echo uninstallation script must be run within the scripts folder
  exit 1
fi

echo domainfinder uninstallation started
rm "$GOPATH"/bin/coolify
rm "$GOPATH"/bin/sprinkle
rm "$GOPATH"/bin/available
rm "$GOPATH"/bin/synonym
rm "$GOPATH"/bin/domainify
rm "$GOPATH"/bin/domainfinder
echo domainfinder uninstalled successfully

