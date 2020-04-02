# 获取最新代码
git fetch --all
git reset --hard origin/dev
git checkout dev
git pull
# 编译go文件
go build -o ./output/yhy-run
chmod 777 ./output/yhy-run
echo "编译完成"
# 关闭原先的go进程
ps -aux|grep yhy-run| grep -v grep | awk '{print $2}' | xargs -r kill -9
echo "关闭进程"
# 重新启动
nohup ./output/yhy-run
echo "重启进程"
echo "正常退出"
exit
