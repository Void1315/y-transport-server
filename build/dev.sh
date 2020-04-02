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
screen -dm "$(./output/yhy-run)"
let try_num=0
while([ ! -n "$(ps -aux|grep yhy-run| grep -v grep | awk '{print $2}')" ])
do
    sleep 1
    try_num+=1
    if [ "$try_num" -lt 60 ]
    then
        echo "重启进程中...($try_num)"
    else
        echo "重启进程超时"
        exit 1
    fi
done
sleep 3
screen -dR
echo "启动成功"
exit 0
