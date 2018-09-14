#!/bin/bash

readonly BASEDIR="$(cd "$(dirname "$0")"/.. && pwd)"

main(){
  run_tests
}

run_tests() {
  pushd "$BASEDIR" || exit 1
    ginkgo -r -keepGoing --skipPackage=launcher,packs,integration
  popd || exit 1
}

main "$@"
