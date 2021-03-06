#!/usr/bin/env bash

# Copyright 2020 Rafael Fernández López <ereslibre@ereslibre.es>
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

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

BOILERPLATE="$(cat ${SCRIPT_DIR}/../hack/boilerplate.go.txt)"

RELEASE_DATA=$(cat ${SCRIPT_DIR}/../RELEASE)

cat <<EOF > ${SCRIPT_DIR}/../internal/pkg/constants/zz_generated.constants.go
$BOILERPLATE

package constants

// Code generated by release-gen script. DO NOT EDIT.

const (
	// RawReleaseData represents the supported versions for this release
	RawReleaseData = \`$RELEASE_DATA\`
)
EOF
