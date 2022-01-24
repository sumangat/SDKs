#!/bin/sh
#
# Usage example: /bin/bash git_push_sdk.sh java-msx-sdk v1.0.10 `cat ~/.ssh/personal_access_token`

msx_sdk=$1
release_tag=$2
git_token=$3

if [ "$release_tag" = "" ]; then
    echo "[INFO] No command line input provided. Set \$release_tag to $release_tag"
    exit
fi

if [ "$git_token" = "" ]; then
    echo "[INFO] No command line input provided. Set \$git_token to $git_token"
    exit
fi

# Adds the files in the local repository and stages them for commit.
echo "Git add"
git add .

# Commits the tracked changes and prepares them to be pushed to a remote repository.
echo "Git commit"
git commit -m "$release_tag"

# Pushes (Forces) the changes in the local repository up to the remote repository
echo "Git pushing $msx_sdk to main"
git push origin main -v 2>&1 | grep -v 'To https'

echo "Git tag $msx_sdk with $release_tag"
git tag -a $release_tag -m "Auto-tagged"
git push --tags

JSON_FMT="{\"tag_name\": \"%s\",\"target_commitish\": \"main\",\"name\": \"%s\",\"body\": \"Auto-released\",\"draft\": false,\"prerelease\": false}"
CURL_DATA=$(printf "$JSON_FMT" "$release_tag" "$release_tag")

# TODO: Need to change this to the fake destination.
CURL_URL="https://api.github.com/repos/sumangat/$msx_sdk"

echo "Git release $msx_sdk version $release_tag"
curl -g -H "Authorization: token $git_token" -H "Content-type:application/json" -d "$CURL_DATA" -L $CURL_URL
