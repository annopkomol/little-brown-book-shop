migrate -version >/dev/null 2>&1
if [ $? -ne 0 ]; then
  printf "You have not install golang-migrate yet.\nPlease visit this site to install migration tool first: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate\n"
  exit
fi
script_path="$(
  cd "$(dirname "$0")" >/dev/null 2>&1
  pwd -P
)"
migrate_path=$(dirname "$script_path")
migrate_path="$migrate_path/migration"
echo "Is this your migration path ?"
read -r -p "'$migrate_path' [Y/n]: " response
if [[ "$response" =~ ^([yY][eE][sS]|[yY])$ ]]; then
  echo "Please provide the name of your migration script file (ex: created_user_table)"
  read migration_name
  migrate create -ext sql -dir $migrate_path -seq $migration_name
else
  echo "Please put this script in the correct path '{your_root_dir}/db/script/' "
fi
