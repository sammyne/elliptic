#!/bin/bash

set -e

remote="https://github.com/miracl/core"
commitID="6c7e3be46ca9c7cf600a83ded084c83b0822d61a"
workingDir=$PWD
outDir=$PWD/_out
goPkgDir=$outDir/go

rm -rf $outDir
git clone $remote $outDir
cd $outDir
git checkout $commitID

cd $goPkgDir
echo -e "28\n0\n" | python3 config32.py
cd -

cp $goPkgDir/core/*.go $goPkgDir/core/BN254/*.go $workingDir/

rm -rf $outDir
