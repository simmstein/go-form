go doc -all -short form.Form | sed 's/^func \(.*\)$/### \0\n\n\`\`\`golang\n\0\n\`\`\`/g' | sed 's/^    \(.*\)/\n\1/g' | sed 's/^### func .* \(.*\)\((.*)\).*/### \1/' | xclip -sel c
