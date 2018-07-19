BASE_PATH="${GOPATH}/src/bitbucket.org/${BITBUCKET_REPO_OWNER}"
mdir -pv ${BASE_PATH}
export PKG_PATH="bitbucket.org/${BITBUCKET_REPO_OWNER}/${BITBUCKET_REPO_SLUG}"

git config --global url."git@bitbucket.org:".insteadOf "https://bitbucket.org/"
