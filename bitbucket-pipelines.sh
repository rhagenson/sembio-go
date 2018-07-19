BASE_PATH="${GOPATH}/src/bitbucket.org/${BITBUCKET_REPO_OWNER}"
PKG_PATH="${BASE_PATH}/${BITBUCKET_REPO_SLUG}"
mkdir -pv ${PKG_PATH}
cd ${PKG_PATH}

git config --global url."git@bitbucket.org:".insteadOf "https://bitbucket.org/"
