// Package embedded makes available the "Mozilla Included CA Certificate List"
// for TLS usage.
//
// This package has no side-effects unlike [github.com/breml/rootcerts].
package embedded

// MozillaCACertificatesPEM returns the "Mozilla Included CA Certificate List"
// (https://wiki.mozilla.org/CA/Included_Certificates) for TLS usage in PEM
// format.
//
// Use of these certificates is governed by Mozilla Public License 2.0
// that can be found in the LICENSE.certificates file.
func MozillaCACertificatesPEM() string {
	return data
}
