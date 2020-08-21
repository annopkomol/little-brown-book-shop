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

db_conn=$(cat $script_path/db_conn_url.txt)
if [ $? -ne 0 ]; then
  echo "please check README.md for migration instruction"
  exit
fi
if [[ "$response" =~ ^([yY][eE][sS]|[yY])$ ]]; then
  PS3='Please specify your command ([1] for Up /[2] for Down): '
  options=("Up" "Down")
  select opt in "${options[@]}"; do
    case $opt in
    "Up")
      MIGRATE_COMMAND=up
      break
      ;;
    "Down")
      MIGRATE_COMMAND=down
      break
      ;;
    *) echo "Invalid option $REPLY" ;;
    esac
  done
  read -r -p "Apply all or N $MIGRATE_COMMAND migrations (all or (1,2,3,...,N): " N_TIME
  if [[ "$N_TIME" =~ ^([aA][lL][lL])$ ]]; then
    N_TIME=""
  fi
  echo "Running migrate -verbose -path $migrate_path -database $db_conn $MIGRATE_COMMAND $N_TIME"
  migrate -verbose -path $migrate_path -database $db_conn $MIGRATE_COMMAND $N_TIME
else
  echo "Please put this script in the correct path '{your_root_dir}/db/script/' "
  exit
fi
