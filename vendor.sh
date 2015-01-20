#!/usr/bin/env bash
set -e

cd "$(dirname "$BASH_SOURCE")"

# Downloads dependencies into vendor/ directory
mkdir -p vendor
cd vendor

clone() {
	vcs=$1
	pkg=$2
	rev=$3

	pkg_url=https://$pkg
	target_dir=src/$pkg

	echo "$pkg @ $rev: "

	if [ -d $target_dir ]; then
		echo "rm old, $pkg"
		rm -fr $target_dir
	fi

	echo "clone, $pkg"
	case $vcs in
		git)
			git clone --quiet --no-checkout $pkg_url $target_dir
			( cd $target_dir && git reset --quiet --hard $rev )
			;;
		hg)
			hg clone --quiet --updaterev $rev $pkg_url $target_dir
			;;
	esac

	echo "rm VCS, $vcs"
	( cd $target_dir && rm -rf .{git,hg} )

	echo "done"
}

clone git github.com/docker/docker v1.4.1
clone git github.com/docker/libcontainer 73ba097bf596249068513559225d6e18c1767b47

echo "don't forget to add vendor folder to your GOPATH (export GOPATH=\$GOPATH:\`pwd\`/vendor)"
