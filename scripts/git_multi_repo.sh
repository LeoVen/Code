# run multiple git commands on multiple repos starting from a parent directory
# this assumes that git repos all have a .git directory

# git pull on all git repositories inside a folder
find . -name .git -type d -prune | xargs dirname | xargs -I {} bash -c "cd {} && git pull"

# another example
# checkout to main
find . -name .git -type d -prune | xargs dirname | xargs -I {} bash -c "cd {} && git checkout main"
