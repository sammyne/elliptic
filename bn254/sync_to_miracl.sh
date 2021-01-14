#!/bin/bash

set -e

remote="https://github.com/miracl/core"
commitID="6c7e3be46ca9c7cf600a83ded084c83b0822d61a"
workingDir=$PWD
localDir=$PWD/internal/miracl
outDir=$PWD/_out
goPkgDir=$outDir/go

rm -rf $localDir
mkdir -p $localDir
rm -rf $outDir

git clone $remote $outDir
cd $outDir
git checkout $commitID

cd $goPkgDir
rm -rf core
echo -e "28\n0\n" | python3 config32.py
cp -r core/* $localDir/
cd -

srcs=$(find $localDir -type f -name "*.go")

for v in ${srcs[@]}; do
  sed -i '' 's!package core!package miracl!g' $v
  sed -i '' 's!"github.com/miracl/core/go/core"!core "github.com/sammyne/elliptic/bn254/internal/miracl"!g' $v
done

#cp $goPkgDir/core/*.go $goPkgDir/core/BN254/*.go $workingDir/
