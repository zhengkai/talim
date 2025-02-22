#!/bin/bash

NAME="${1:-doll}"

cd "$(dirname "$(readlink -f "$0")")" || exit 1

mkdir -p output

# shellcheck disable=SC1090
# . "../env/${NAME}" || exit
. ./env || exit
if [ -z "$BRANCH" ]; then
	export BRANCH=master
fi

envsubst < ./tpl-install.sh > output/install.sh
chmod +x output/install.sh

cat output/install.sh

TARGET="${NAME}:/www/talim/docker/"
scp output/install.sh "$TARGET"
# rsync --partial -vzrtopg -e ssh "output/talim-image-${BRANCH}.tar" "$TARGET"
