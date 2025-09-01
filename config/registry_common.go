// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package config

import (
	_ "embed"
)

var (
	//go:embed schema.json
	providerSchema string

	//go:embed provider-metadata.yaml
	providerMetadata []byte
)
