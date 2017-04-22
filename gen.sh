#/bin/sh

set -ex

if [ ! -z "$1" ]; then GOROOT=$1; fi

cp -f $GOROOT/src/runtime/go_tls.h internal/HEAD/go_tls.h
WORK=$(go build -a -x -work runtime 2>&1|sed -e 's/^WORK=//g; t; d;')
cp $WORK/runtime/_obj/go_asm.h internal/HEAD/go_asm.h
rm -rf $WORK
cat <<EOF > HEAD.go
package panick

import _ "github.com/t-yuki/panick/internal/HEAD"
EOF


