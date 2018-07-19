BASE_PATH="${GOPATH}/src/bitbucket.org/${BITBUCKET_REPO_OWNER}"
mkdir -p ${BASE_PATH}
export PKG_PATH="${BASE_PATH}/${BITBUCKET_REPO_SLUG}"

git config --global url."git@bitbucket.org:".insteadOf "https://bitbucket.org/"
