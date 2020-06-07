set -ex

cd ~/isucon9-qualify/webapp/go/ && rm -rf isucari && make
cd ~/isucon9-qualify/ && sudo systemctl restart isucari
