
# 获取最新代码
gitSynchronize(){
    git fetch --all
    git reset --hard origin/master
    git checkout master
    git pull
}

# 编译go文件
runBuild(){
    go build -o ./output/yhy-run
    chmod 777 ./output/yhy-run
}
# 运行程序
runProgram(){
    # 关闭原先的go进程
    ps -aux|grep yhy-run| grep -v grep | awk '{print $2}' | xargs -r kill -9
    # 重新启动
    nohup ./output/yhy-run >/var/log/yhy.log 2>&1 &
    try_num=0
    while([ ! -n "$(ps -aux|grep yhy-run| grep -v grep | awk '{print $2}')" ])
    do
        sleep 1
        let try_num+=1
        if [ "$try_num" -lt 30 ]
        then
            echo "重启进程中...($try_num)"
        else
            echo "重启进程超时"
            exit 1
        fi
    done
    sleep 3
}

export GIN_MODE=release

echo "-----同步代码-----"
gitSynchronize
echo "-----同步代码完毕-----"

echo "-----编译程序-----"
runBuild
echo "-----编译程序完毕-----"

echo "-----启动程序-----"
runProgram
echo "-----启动程序完毕-----"

exit 0
