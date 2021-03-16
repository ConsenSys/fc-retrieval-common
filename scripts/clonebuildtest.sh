#!/bin/bash
set -e
echo Clone, Build, and Test for branch: $1

if [ -z "$1" ]
  then
    echo "No argument supplied. Please specify branch to test"
    exit
fi

cd ../..
mkdir -p clonebuildtest
cd clonebuildtest
rm -rf $1
mkdir $1
cd $1
BRANCH=$1


clone_on_branch() {
  echo "Cloning repo: $1"
  git clone https://github.com/ConsenSys/fc-retrieval-$1.git
  cd fc-retrieval-$1
  git checkout $BRANCH
  cd ..
}



check_deps() {
  echo "Checking repo: $1"
  cd ../fc-retrieval-$1
  make useremote
  if [ -n "$(git status --porcelain)" ]; then 
    echo "Dependancies not up to date $1"
    # exit
  fi
}

build_tag() {
  cd ../fc-retrieval-$1
  make build tag
}


#git clone https://github.com/ConsenSys/fc-retrieval-client.git
#git clone https://github.com/ConsenSys/fc-retrieval-common.git
#git clone https://github.com/ConsenSys/fc-retrieval-gateway.git
#git clone https://github.com/ConsenSys/fc-retrieval-gateway-admin.git
#git clone https://github.com/ConsenSys/fc-retrieval-itest.git
#git clone https://github.com/ConsenSys/fc-retrieval-provider.git
#git clone https://github.com/ConsenSys/fc-retrieval-provider-admin.git
#git clone https://github.com/ConsenSys/fc-retrieval-register.git

clone_on_branch client
clone_on_branch common
clone_on_branch gateway
clone_on_branch gateway-admin
clone_on_branch itest
clone_on_branch provider
clone_on_branch provider-admin
clone_on_branch register

# Make sure the branch exists on itest.
cd fc-retrieval-itest
git checkout $BRANCH
ITEST_REPO_BRANCH=`git rev-parse --abbrev-ref HEAD`
if [ $ITEST_REPO_BRANCH != $BRANCH ]; then 
  echo "Branch $BRANCH does not exist on itest"
  exit
fi


cd ../fc-retrieval-common
bash scripts/dockerclean.sh
git checkout $BRANCH
#check_deps common
check_deps register
check_deps client
check_deps provider-admin
check_deps gateway-admin
check_deps gateway
check_deps provider

build_tag register
build_tag gateway
build_tag provider
build_tag itest

make itestlocal
