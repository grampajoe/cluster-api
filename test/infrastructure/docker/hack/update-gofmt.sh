#!/usr/bin/env bash
# Copyright 2019 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# script to run gofmt over our code (not vendor)
set -o errexit
set -o nounset
set -o pipefail

# shellcheck source=./hack/utils.sh
source "$(git rev-parse --show-toplevel)/hack/utils.sh"

cd_capd_root_path

# update go fmt
go fmt ./...
