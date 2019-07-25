# /bin/bash
# 同步项目目录

# 判断是否为项目目录
if [ ! -d $(pwd)"/.git/" ];then
    echo "请到正确的项目目录执行！"
    exit
fi

echo -e "\033[31m 密码：Dove@2018 \033[0m"
scp -P 8885 -r $(pwd) homework@192.168.240.255:/home/homework/app/ 1>> /dev/null 2>> /dev/null

echo -e "\033[31m done. \033[0m"