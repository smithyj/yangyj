# 杀掉 air 进程
ps -ef | grep -w air | grep -v grep | sort -k 2rn | awk '{if (NR>1){print $2}}' | xargs kill -9
# 杀掉 serve 进程
lsof -i:8080 | grep serve | awk '{print $2}' | xargs kill -9
# 杀掉 dlv 进程
lsof -i:2345 | grep dlv | awk '{print $2}' | xargs kill -9
# debug
dlv debug --listen=:2345 --headless=true --api-version=2 --continue --accept-multiclient --output=./tmp/serve ./cmd/serve/main.go