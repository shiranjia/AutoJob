#AutoJob
an auto job tool  自动部署工具</br>
新版本技术使用 [vue2](https://github.com/shiranjia/AutoJob-vue "AutoJob-vue") + go rest api   访问地址：http://127.0.0.1</br>
vue生成的js，css放在七牛存储</br>
旧版本技术使用 jquery + go server   访问地址：http://127.0.0.1/old</br>
任意添加本地命令
任意添加远程命令
任意上传文件或目录</br>
######获取方式： go get github.com/shiranjia/AutoJob</br>
######将项目路径添加到GOPATH</br>
######编译：AutoJob/src go install AutoJob</br>
######../bin目录下生成 AutoJob 可执行文件，直接运行，浏览器访问127.0.0.1即可</br>
[下载地址](https://github.com/shiranjia/AutoJob/releases  "releases")</br></br>
新版本使用：![image](https://github.com/shiranjia/AutoJob/blob/master/resources/20170105150404.png)</br>
旧版本使用： </br>
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

####命令行单独执行任务：AutoDeploy.exe -j jobName -p dataPath
