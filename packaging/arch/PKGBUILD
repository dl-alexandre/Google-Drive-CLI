pkgname=gdrv-git
pkgver=0.1.0.3.g38517bd
pkgrel=1
pkgdesc="Google Drive CLI"
arch=('x86_64' 'aarch64')
url="https://github.com/dl-alexandre/Google-Drive-CLI"
license=('MIT')
depends=('libsecret')
makedepends=('git' 'go')
provides=('gdrv')
conflicts=('gdrv')
source=("git+https://github.com/dl-alexandre/Google-Drive-CLI.git")
sha256sums=('SKIP')

pkgver() {
	cd "$srcdir/Google-Drive-CLI"
	git describe --tags --long --match 'v*' | sed 's/^v//;s/-/./g'
}

build() {
	cd "$srcdir/Google-Drive-CLI"
	local ldflags="-s -w -buildid= -X github.com/dl-alexandre/gdrv/internal/cli.version=${pkgver}"
	go build -trimpath -buildmode=pie -mod=readonly -ldflags "$ldflags" -o gdrv ./cmd/gdrv
}

package() {
	cd "$srcdir/Google-Drive-CLI"
	install -Dm755 gdrv "$pkgdir/usr/bin/gdrv"
	install -Dm644 LICENSE "$pkgdir/usr/share/licenses/$pkgname/LICENSE"
}
