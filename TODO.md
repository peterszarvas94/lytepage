- dev mode
- embed lt2-test to lt2
- init command for lt2 cli
- replace lytepage repo on github with lt2

when running "go install" on root, I need "go.mod.tmpl" in the embedded folder
when developing root, I need "go.mod.tmpl" in the embedded folder

when developing embedded folder, I need "go.mod"

when running the cli "lt2 init", I need to copy embedded folder, and rename "go.mod.tmpl" to "go.mod"
