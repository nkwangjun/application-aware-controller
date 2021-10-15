#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..

# 安装k8s.io/code-generator
[[ -d $GOPATH/src/k8s.io/code-generator ]] || go get -u k8s.io/code-generator/...

# 执行代码自动生成, 其中pkg/client是生成目标目录, pkg/apis是类型定义目录
bash $GOPATH/src/k8s.io/code-generator/generate-groups.sh all \
  k8s.io/application-aware-controller/pkg/generated k8s.io/application-aware-controller/pkg/apis \
  aacontroller:v1 \
