# Write in a file to-git-or-not-to-git.sh the command that isolates your Gitea id.
# Only the numbers will appear.
# Here is the base command that needs to be adapted with your username and more :

var=$(curl -s https://01.kood.tech/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"simone\"}}){id}}"}')
echo "$var" | tr -dc '0-9';