# Get the version.
version=`git describe --tags --long`
# Write out the package.
cat << EOF > version_linux.go
package main

//go:generate bash ./get_version_linux.sh
var version = "$version"
EOF
