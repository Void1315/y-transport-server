# 获取最新代码
git fetch --all
git reset --hard origin/master
git checkout master
git pull
# 编译go文件
go build -o ./output/yhy-run
chmod 777 ./output/yhy-run

# 关闭原先的go进程
ps -aux|grep yhy-run| grep -v grep | awk '{print $2}' | xargs -r kill -9

# 设置环境变量
export GIN_MODE=release

# 重新启动
nohup ./output/run &