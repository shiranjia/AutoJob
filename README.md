#AutoDeploy
an auto deploy tool  自动部署工具

任意添加本地命令
任意添加远程命令
任意上传文件或目录
获取方式：go get github.com/shiranjia/AutoDeploy
编译    ：AutoDeploy/src> go install AutoDeploy
../bin目录下生成 AutoDeploy 可执行文件，直接运行，浏览器访问127.0.0.1即可
例子： </br>
######ssh.userName: root ssh.password: ••••••••  ssh.ip:  192.168.104.141

--------------------------------------------------------------------------------------------------------------------------------------

######本地命令 (多条换行 执行路径;命令;参数(参数用空格隔开)
######example:E:\github\web;mvn;clean package -Dmaven.test.skip=true -P artifactory,development -Dfile.encoding=UTF-8 ): </br>
--------------------------------------------------------------------------------------------------------------------------------------
######E:\github\web;mvn;clean package -Dmaven.test.skip=true -P artifactory,development -Dfile.encoding=UTF-8 

--------------------------------------------------------------------------------------------------------------------------------------
######远程命令(多条换行):
######rm -rf /path/*

--------------------------------------------------------------------------------------------------------------------------------------

######上传文件: 本地目录： E:/github/web/target/web 远程目录：/home

--------------------------------------------------------------------------------------------------------------------------------------

#######本地命令 (多条换行 执行路径;命令;参数(参数用空格隔开)
#######example:E:\github\web;mav;clean package -Dmaven.test.skip=true -P artifactory,development -Dfile.encoding=UTF-8 ): 

--------------------------------------------------------------------------------------------------------------------------------------

######远程命令(多条换行):  
######mv /home/web/* /path/
######sh /shellPath/restart

--------------------------------------------------------------------------------------------------------------------------------------

######命令行单独执行job:AutoDeploy.exe -j jobName -p dataPath
