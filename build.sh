#!/usr/bin/env bash

package="github.com/ericogr/hdce"
package_split=(${package//\// })
package_name=${package_split[-1]}
platforms=("windows/amd64" "linux/amd64")
output_dir=bin
ldflags="-s -w"
upxflags="-q"

#install common dependencies

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=$package_name'-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    echo -n "building $output_dir/$output_name: "
    env GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="$ldflags" -o $output_dir/$output_name $package
    if [ $? -ne 0 ]; then
        echo 'An error has occurred (build)! Aborting the script execution...'
        exit 1
    fi
    echo "ok"

    echo -n "compressing $output_dir/$output_name: "
    upx $upxflags $output_dir/$output_name>/dev/null
    if [ $? -ne 0 ]; then
        echo 'An error has occurred (upx)! Aborting the script execution...'
        echo '>>> Is upx installed? <<<'
        exit 1
    fi
    echo "ok"
    echo
done
