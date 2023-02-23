# TODO:

- All messages in array and CLI arg to print.
- Change setting about min/max length data type.
- Verify syntax and spell check for all messages.
- Indexes lint.
- Foreign Key lint.
- Verify cardinality with MySQL and data.
- List all codes and checks.
- Allow to connect to MySQL for lint tables and analize cardinality.
- Put sumary when it finish, how many errors? It seems everything is right.
- Sync code version with tag.
- install via linux and homebrew.

tener el modo para hacer post de comment en github

      - name: Post Infracost comment
        run: |
          infracost comment github --path=/tmp/infracost.json \
                                   --repo=$GITHUB_REPOSITORY \
                                   --github-token=${{github.token}} \
                                   --pull-request=${{github.event.pull_request.number}} \
                                   --behavior=update
