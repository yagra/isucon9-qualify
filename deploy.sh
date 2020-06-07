set -ex

cd ~/isucon9-qualify/webapp/go/ && rm -rf isucari && make
cd ~/isucon9-qualify/ && sudo systemctl restart isucari && ./bin/benchmarker -target-url http://ec2-54-199-15-35.ap-northeast-1.compute.amazonaws.com/
