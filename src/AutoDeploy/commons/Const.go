package commons

import "log"

const Html  = "<html>\n" +
"<head><title>welcome</title></head>\n" +
"<script type=\"text/javascript\" src=\"https://code.jquery.com/jquery-3.0.0.min.js\"></script>\n" +
"<style>\n" +
"    .black{  background-image: url(http://s5.51cto.com/wyfs02/M02/73/4B/wKiom1X5NxXxpjKBAADbGtrjmj4613.jpg);  height: 175px;  }\n" +
"    .listBlack{  background-image: url(http://i44.tinypic.com/33f5ev7.jpg);  }\n" +
"</style>\n" +
"<body>\n" +
"<div class=\"black\">\n" +
"<div name=\"t_name\" style=\"float: left;\" >\n" +
"       {{range $index,$v := .}}\n" +
"       <button name=\"span_name\" onclick=\"swich({{$v.Name}})\" id=\"name{{$v.Name}}\"  {{if  eq $index 0 }} style=\"background-color: brown;\" {{end}}>{{$v.Name}}</button>\n" +
"       &nbsp;\n" +
"       {{end}}\n" +
"\n" +
"</div>\n" +
"<div>\n" +
"    <button onclick=\"newJob()\" id=\"newJobButton\" style=\"background-color:greenyellow;\">new job</button>&nbsp;\n" +
"</div>\n" +
"</div>\n" +
"</br>\n" +
"<div id=\"jobList\" class=\"listBlack\">\n" +
"{{range $index,$v := .}}\n" +
"<div name=\"deploy\"  id=\"{{.Name}}\" {{if  gt $index 0 }} style=\"display:none;\" {{end}}>\n" +
" <form action=\"#\" method=\"post\" id=\"form_{{.Name}}\">\n" +
"         <div>\n" +
"             <input type=\"hidden\" name=\"deploy.Name\" value=\"{{.Name}}\" />\n" +
"             <h3>jobName：<label>{{.Name}}</label></h3>\n" +
"         </div>\n" +
"     <span>---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------</span>\n" +
"     </br>\n" +
"        <div>\n" +
"             <label>ssh.userName</label><input type=\"text\" name=\"config.User\" value=\"{{.Config.User}}\" />\n" +
"             <label>ssh.password</label><input type=\"password\" name=\"config.Password\" value=\"{{.Config.Password}}\" />\n" +
"             <label>ssh.ip</label><input type=\"text\" name=\"config.ip\" value=\"{{.Config.Ip}}\" />\n" +
"        </div>\n" +
"     <span>---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------</span>\n" +
"     </br>\n" +
"         <div>\n" +
"         本地命令\n" +
"         (多条换行 执行路径;命令;参数(参数用空格隔开) example:E:\\github\\ticket.h5\\web;mvn;clean package -Dmaven.test.skip=true -P artifactory,development -Dfile.encoding=UTF-8 ):\n" +
"         <textarea name=\"LocalBefore.Command\" class=\"n\" rows=\"3\" cols=\"200\">{{range .LocalBefore}}{{if ne .Path \"\"}}{{.Path}};{{.Command}};{{range .Args}}{{.}} {{end}}-n-{{end}}{{end}}</textarea>\n" +
"         </div>\n" +
"     <span>---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------</span>\n" +
"     </br>\n" +
"         <div>\n" +
"             远程命令(多条换行):<textarea name=\"RemoteBefore.Command\" class=\"n\" rows=\"3\" cols=\"180\">{{range .RemoteBefore}}{{.Command}}-n-{{end}}</textarea>\n" +
"         </div>\n" +
"     <span>---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------</span>\n" +
"     </br>\n" +
"         <div>\n" +
"             上传文件:\n" +
"             本地目录： <input name=\"Upload.Path\" type=\"text\" size=\"78\" value=\"{{.UploadJob.LocalPath}}\"/>\n" +
"             远程目录：<input name=\"Upload.RemotePath\" type=\"text\" size=\"78\" value=\"{{.UploadJob.RemotePath}}\"/>\n" +
"         </div>\n" +
"     <span>---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------</span>\n" +
"     </br>\n" +
"         <div>\n" +
"             本地命令\n" +
"             (多条换行 执行路径;命令;参数(参数用空格隔开) example:E:\\github\\ticket.h5\\web;mav;clean package -Dmaven.test.skip=true -P artifactory,development -Dfile.encoding=UTF-8 ):\n" +
"             <textarea name=\"LocalAfter.Command\" class=\"n\" rows=\"3\" cols=\"200\">{{range .LocalAfter}}{{if ne .Path \"\"}}{{.Path}};{{.Command}};{{range .Args}}{{.}} {{end}}-n-{{end}}{{end}}</textarea>\n" +
"         </div>\n" +
"     <span>---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------</span>\n" +
"     </br>\n" +
"         <div>\n" +
"             远程命令(多条换行):\n" +
"             <textarea id=\"remoteAfter{{$index}}\" class=\"n\" name=\"RemoteAfter.Command\" rows=\"3\" cols=\"180\">{{range .RemoteAfter}}{{.Command}}-n-{{end}}</textarea>\n" +
"         </div>\n" +
"     <span>---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------</span>\n" +
"     </br>\n" +
"     <button type=\"button\" onclick=\"deploy({{.Name}})\">deploy</button>&nbsp;<button type=\"button\" onclick=\"update({{.Name}})\">update</button>&nbsp;<button type=\"button\" onclick=\"deleteJob({{.Name}})\">delete</button>\n" +
" </form>\n" +
"</div>\n" +
"{{end}}\n" +
"</div>\n" +
"<div id=\"newJob\" style=\"display: none\" class=\"listBlack\">\n" +
"    <form action=\"saveOrUpdate\" method=\"post\">\n" +
"        <div>\n" +
"            <h3>jobName (不可重复)：<input name=\"deploy.Name\" type=\"text\"  value=\"\"/></h3>\n" +
"        </div>\n" +
"        <span>---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------</span>\n" +
"        </br>\n" +
"        <div>\n" +
"            <label>ssh.userName:</label><input type=\"text\" name=\"config.User\" />\n" +
"            <label>ssh.password:</label><input type=\"text\" name=\"config.Password\" />\n" +
"            <label>ssh.ip:</label><input type=\"text\" name=\"config.ip\" />\n" +
"        </div>\n" +
"        <span>---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------</span>\n" +
"        </br>\n" +
"        <div>\n" +
"            本地命令\n" +
"            (多条换行 执行路径;命令;参数(参数用空格隔开) example:E:\\github\\ticket.h5\\web;mvn;clean package -Dmaven.test.skip=true -P artifactory,development -Dfile.encoding=UTF-8 ):\n" +
"            <textarea name=\"LocalBefore.Command\" rows=\"3\" cols=\"200\"></textarea>\n" +
"        </div>\n" +
"        <span>---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------</span>\n" +
"        </br>\n" +
"        <div>\n" +
"            远程命令(多条换行):<textarea name=\"RemoteBefore.Command\" rows=\"3\" cols=\"180\"></textarea>\n" +
"        </div>\n" +
"        <span>---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------</span>\n" +
"        </br>\n" +
"        <div>\n" +
"            上传文件:\n" +
"            本地目录： <input name=\"Upload.Path\" type=\"text\" size=\"78\" />\n" +
"            远程目录： <input name=\"Upload.RemotePath\" type=\"text\" size=\"78\"/>\n" +
"        </div>\n" +
"        <span>---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------</span>\n" +
"        </br>\n" +
"        <div>\n" +
"            本地命令\n" +
"            (多条换行 执行路径;命令;参数(参数用空格隔开) example:E:\\github\\ticket.h5\\web;mav;clean package -Dmaven.test.skip=true -P artifactory,development -Dfile.encoding=UTF-8 ):\n" +
"            <textarea name=\"LocalAfter.Command\" rows=\"3\" cols=\"200\"></textarea>\n" +
"        </div>\n" +
"        <span>---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------</span>\n" +
"        </br>\n" +
"        <div>\n" +
"            远程命令(多条换行):\n" +
"            <textarea name=\"RemoteAfter.Command\" rows=\"3\" cols=\"180\"></textarea>\n" +
"        </div>\n" +
"        <span>---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------</span>\n" +
"        </br>\n" +
"        <button>save</button>&nbsp;\n" +
"    </form>\n" +
"</div>\n" +
"\n" +
"<script>\n" +
"    var url = window.location.href;\n" +
"    if(url.indexOf(\"index\")){\n" +
"\n" +
"    }\n" +
"$(document).ready(function() {\n" +
"    var text = $(\"textarea.n\")\n" +
"    for(var i=0;i<text.length;i++){\n" +
"        t = $(text[i]).html();\n" +
"        while (t.indexOf(\"-n-\") > 0){\n" +
"            t = t.replace(\"-n-\",\"\\r\\n\");\n" +
"        }\n" +
"        $(text[i]).html(t)\n" +
"    }\n" +
"\n" +
"});\n" +
"function swich(name) {\n" +
"    $(\"#jobList\").attr(\"style\",\"display:block;\")\n" +
"    $(\"#newJob\").attr(\"style\",\"display:none;\")\n" +
"    $(\"[name=span_name]\").attr(\"style\",\"background-color: none;\")\n" +
"    $(\"#name\" + name).attr(\"style\",\"background-color: brown;\");\n" +
"    $(\"[name=deploy]\").attr(\"style\",\"display:none;\")\n" +
"    $(\"#\" + name).attr(\"style\",\"display:block;\")\n" +
"}\n" +
"function newJob() {\n" +
"        $(\"#jobList\").attr(\"style\",\"display:none;\")\n" +
"        $(\"#newJob\").attr(\"style\",\"display:block;\")\n" +
"    }\n" +
"function deploy(jobName){\n" +
"    var form = $(\"#form_\" + jobName);\n" +
"    $(form).attr(\"action\",\"deploy\");\n" +
"    $(form).submit();\n" +
"}\n" +
"function update(jobName){\n" +
"    var form = $(\"#form_\" + jobName)\n" +
"    $(form).attr(\"action\",\"saveOrUpdate\");\n" +
"    $(form).submit();\n" +
"}function deleteJob(jobName){var form = $(\"#form_\" + jobName);$(form).attr(\"action\",\"delete\");$(form).submit();}</script></body></html>"

type RemoteOutPut struct{
	Name string
}

func (s *RemoteOutPut) Write(p []byte) (n int, err error)  {
	log.Printf("remote.out:",string(p))
	return len(p),nil
}
