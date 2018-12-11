BASE_PATH="${GOPATH}/src/bitbucket.org/${BITBUCKET_REPO_OWNER}"
mkdir -pv ${BASE_PATH}
export PKG_PATH="bitbucket.org/rhagenson/${BITBUCKET_REPO_SLUG}"

git config --global url."git@bitbucket.org:".insteadOf "https://bitbucket.org/"