
git submodule update --remote --merge

git push --recurse-submodules=check

git push --recurse-submodules=on-demand

git submodule foreach 'git checkout -b featureA'

$ git config alias.sdiff '!'"git diff && git submodule foreach 'git diff'"
$ git config alias.spush 'push --recurse-submodules=on-demand'
$ git config alias.supdate 'submodule update --remote --merge'


git clone --recursive https://github.com/chaconinc/MainProject

--