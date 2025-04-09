# unicode-tools

```fish
function fe
        cat $HOME/src/github.com/wayneashleyberry/unicode-tools/release/latest-en.txt | fzf --header='Unicode search! <ctrl-o> to copy the selection. ' --wrap --bind 'ctrl-o:execute-silent(echo {} | cut -f1 | pbcopy)' | tee /dev/tty | cut -f1 | pbcopy
end
```
