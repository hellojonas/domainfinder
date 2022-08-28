#!/usr/bin/env bash

if [[ $(basename "$PWD") != "scripts" ]] 
then
  echo installation script must be run within the scripts folder
  exit 1
fi

echo installation started
cd ../cmd > /dev/null

echo installing sprinkle...
go install ./sprinkle/sprinkle.go

echo installing coolify...
go install ./coolify/coolify.go

echo installing synonym...
go install ./synonym/synonym.go

echo installing domainify...
go install ./domainify/domainify.go

echo installing available...
go install ./available/available.go

echo installing domainfinder...
go install ./domainfinder/domainfinder.go

echo installation success
cd - > /dev/null
