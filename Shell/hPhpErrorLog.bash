# /bin/bash
# PHP error_log 工具函数

# 定义目录
hPhpErrorLog='/tmp/hPhpErrorLog.php'

# 验证文件是否存在
if [[ -f $hPhpErrorLog ]]; then
   rm $hPhpErrorLog
fi

# 新建文件
touch $hPhpErrorLog

# 写入工具函数
echo '
<?php
    if (function_exists('hlog')) {
        die("hlog 函数已经被定义.");
    }
    else {
        function hlog($source) {

            $scriptPath = debug_backtrace()[0];

            error_log("----- start ----");

            if (is_string($source)) {
                error_log($source);
            }
            else {
                error_log(print_r($source, true));
            }

            error_log("-----  end  ----");
            error_log("log路径为: " . $scriptPath['file']);
        }
    }
' >> $hPhpErrorLog

# 循环项目目录
path='/home/homework/webroot'
for file in `ls $path`
do
    if [[ ! -d $path/$file ]]; then
        continue
    fi

    if [[ ! -f $path/$file/index.php ]]; then
        continue
    fi

    para1=`grep -rn 'hPhpErrorLog' $path/$file/index.php`
    if [ -n "$para1" ]; then
      continue
    fi

    echo "require '$hPhpErrorLog';" >> $path/$file/index.php

done

echo -e "\033[0;32mDone...\033[0m"


