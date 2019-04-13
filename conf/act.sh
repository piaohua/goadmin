#!/bin/bash

set -e

usage() {
    echo " sh ./act.sh start|stop|open|close|create|update "
    echo " sh ./act.sh start host path name port "
    echo " sh ./act.sh stop host path name "
    echo " sh ./act.sh open host path name port "
    echo " sh ./act.sh close host path name port "
    echo " sh ./act.sh create host path name "
    echo " sh ./act.sh update host path name "
}

BINPATH=/mnt/bin
NGINXPATH=/usr/local/nginx
NGINXBIN=${NGINXPATH}/sbin/nginx

#BINPATH=~/go/src/mytest/bin
#NGINXPATH=~/go/src/mytest/bin/nginx
#NGINXBIN=${NGINXPATH}/sbin/nginx

get_inet_addr() {
    echo `ifconfig eth0 | grep "inet addr" | awk '{ print $2}' | awk -F: '{print $2}'`
    #echo "127.0.0.1"
}

is_remote() {
    host=$1
    inet=`get_inet_addr`
    if [[ ${host} == ${inet} ]]; then
        echo "false"
    else
        echo "true"
    fi
}

remote_exec() {
    host=$1
    cmd=$2
    if [[ ${host} == '' ]]; then
        echo "host not exist"
        exit 1
    fi
    if [[ "${cmd}" == '' ]]; then
        echo "cmd Can't be empty"
        exit 1
    fi
    ssh -i ~/.ssh/id_rsa root@${host} "${cmd}"
}

remote_scp() {
    host=$1
    file=$2
    remotepath=$3
    if [[ ${host} == '' ]]; then
        echo "host Can't be empty"
        exit 1
    fi
    if [[ ! -e ${file} ]]; then
        echo "${file} file not exist"
        exit 1
    fi
    if [[ ${remotepath} == '' ]]; then
        echo "path Can't be empty"
        exit 1
    fi
    scp -i ~/.ssh/id_rsa ${file} root@${host}:${remotepath}
}

show_status() {
    result=$1
    echo `echo "${result}" | awk '{ print $10}'`
}

get_status() {
    result=$1
    echo `echo "${result}" | awk '{ print $4}'`
}

pm2show() {
    name=$1
    num=`ps aux | grep -i "${name}-bin" | grep -v 'grep' | wc -l`
    if [ ${num} -eq 0 ]; then
        echo ""
    else
        result=`pm2 show ${name} | grep -i "status"`
        echo `get_status "${result}"`
    fi
}

pm2show_remote() {
    host=$1
    name=$2
    cmd="ps aux | grep -i '${name}-bin' | grep -v 'grep' | wc -l"
    num=`remote_exec ${host} "${cmd}"`
    if [ ${num} -eq 0 ]; then
        echo ""
    else
        cmd="pm2 show ${name} | grep -i 'status'"
        result=`remote_exec ${host} "${cmd}"`
        echo `get_status "${result}"`
    fi
}

check() {
    host=$1
    path=$2
    name=$3
    if [[ ${host} == '' ]]; then
        echo "host Can't be empty"
        exit 1
    fi
    if [[ ${path} == '' ]]; then
        echo "path Can't be empty"
        exit 1
    fi
    if [[ ${name} == '' ]]; then
        echo "name Can't be empty"
        exit 1
    fi
}

check_file() {
    path=$1
    jsonFile=$2
    shFile=$3
    if [[ ! -d ${path} ]]; then
        echo "${path} path not exist"
        exit 1
    fi
    if [[ ! -e ${jsonFile} ]]; then
        echo "${jsonFile} json file not exist"
        exit 1
    fi
    if [[ ! -e ${shFile} ]]; then
        echo "${shFile} sh file not exist"
        exit 1
    fi
}

check_file_remote() {
    host=$1
    path=$2
    jsonFile=$3
    shFile=$4
    remote_exec ${host} "test -d ${path}"
    if [ $(echo $?) -ne 0 ]; then
        echo "${path} remote path not exist"
        exit 1
    fi
    remote_exec ${host} "test -e ${jsonFile}"
    if [ $(echo $?) -ne 0 ]; then
        echo "${jsonFile} remote json file not exist"
        exit 1
    fi
    remote_exec ${host} "test -e ${shFile}"
    if [ $(echo $?) -ne 0 ]; then
        echo "${shFile} remote sh file not exist"
        exit 1
    fi
}

check_port() {
    port=$1
    if [[ ${port} == '' ]]; then
        echo "port Can't be empty"
        exit 1
    fi
    ss -tnlp | grep -i ":${port}" | wc -l
}

check_port_remote() {
    host=$1
    port=$2
    if [[ ${host} == '' ]]; then
        echo "host Can't be empty"
        exit 1
    fi
    if [[ ${port} == '' ]]; then
        echo "port Can't be empty"
        exit 1
    fi
    echo `remote_exec ${host} "ss -tnlp | grep -i ':${port}' | wc -l"`
}

check_bin_file() {
    binPath=$1
    if [[ ! -e ${binPath} ]]; then
        echo "${binPath} bin file not exist"
        exit 1
    fi
}

start() {
    host=$1
    path=$2
    name=$3
    port=$4
    check ${host} ${path} ${name}
    echo "启动主机${host}中名为${name}的服务,端口${port},路径{$path}"
    jsonFile=${path}/pm2/${name}.json
    shFile=${path}/pm2/${name}.sh
    isRemote=`is_remote ${host}`
    if [[ ${isRemote} == "true" ]]; then
        check_file_remote ${host} ${path} ${jsonFile} ${shFile}
        status=`pm2show_remote ${host} ${name}`
        if [[ ${status} == 'online' ]]; then
            echo "${name} is already running"
            exit 1
        fi
        result=`check_port_remote ${host} ${port}`
        if [ $result -ne 0 ]; then
            echo "${port} is already listen"
            exit 1
        fi
        result=`remote_exec ${host} "pm2 start ${jsonFile} | grep -i '${name}' | grep -i 'fork'"`
        echo `show_status "${result}"`
    else
        check_file ${path} ${jsonFile} ${shFile}
        status=`pm2show ${name}`
        if [[ ${status} == 'online' ]]; then
            echo "${name} is already running"
            exit 1
        fi
        result=`check_port ${port}`
        if [ $result -ne 0 ]; then
            echo "${port} is already listen"
            exit 1
        fi
        result=`pm2 start ${jsonFile} | grep -i "${name}" | grep -i "fork"`
        # TODO ./ctrl start
        echo `show_status "${result}"`
    fi
    echo "服务${name}启动成功"
}

stop() {
    host=$1
    path=$2
    name=$3
    check ${host} ${path} ${name}
    echo "停止主机${host}上服务${name},路径${path}"
    jsonFile=${path}/pm2/${name}.json
    shFile=${path}/pm2/${name}.sh
    isRemote=`is_remote ${host}`
    if [[ ${isRemote} == "true" ]]; then
        check_file_remote ${host} ${path} ${jsonFile} ${shFile}
        status=`pm2show_remote ${host} ${name}`
        if [[ ${status} == 'stopped' ]]; then
            echo "${name} is not running"
            exit 1
        fi
        result=`remote_exec ${host} "pm2 stop ${jsonFile} | grep -i '${name}' | grep -i 'fork'"`
        echo `show_status "${result}"`
    else
        check_file ${path} ${jsonFile} ${shFile}
        status=`pm2show ${name}`
        if [[ ${status} == 'stopped' ]]; then
            echo "${name} is not running"
            exit 1
        fi
        result=`pm2 stop ${jsonFile} | grep -i "${name}" | grep -i 'fork'`
        echo `show_status "${result}"`
    fi
    echo "服务${name}停止成功"
}

update() {
    host=$1
    path=$2
    name=$3
    check ${host} ${path} ${name}
    echo "更新主机${host}上服务${name}的bin文件,路径${path}"
    jsonFile=${path}/pm2/${name}.json
    shFile=${path}/pm2/${name}.sh
    binFile=${name}-bin
    n=${#name}
    appName=${name:0:n-1}
    appBinFile=${appName}-bin
    binPath=${BINPATH}/${appBinFile}
    check_bin_file ${binPath}
    isRemote=`is_remote ${host}`
    if [[ ${isRemote} == "true" ]]; then
        check_file_remote ${host} ${path} ${jsonFile} ${shFile}
        status=`pm2show_remote ${host} ${name}`
        if [[ ${status} == 'online' ]]; then
            echo "${name} is already running"
            exit 1
        fi
        remote_scp ${host} ${binPath} ${path}/${binFile}
    else
        check_file ${path} ${jsonFile} ${shFile}
        status=`pm2show ${name}`
        if [[ ${status} == 'online' ]]; then
            echo "${name} is already running"
            exit 1
        fi
        cp -f ${binPath} ${path}/${binFile}
    fi
    echo "服务${name}更新成功"
}

check_src_file() {
    jsonFile=$1
    shFile=$2
    appBinPath=$3
    switchJson=$4
    codeJson=$5
    datx=$6
    yaml=$7
    if [[ ! -e ${jsonFile} ]]; then
        echo "${jsonFile} json file not exist"
        exit 1
    fi
    if [[ ! -e ${shFile} ]]; then
        echo "${shFile} sh file not exist"
        exit 1
    fi
    if [[ ! -e ${appBinPath} ]]; then
        echo "${appBinPath} bin file not exist"
        exit 1
    fi
    if [[ ! -e ${switchJson} ]]; then
        echo "${switchJson} switch json file not exist"
        exit 1
    fi
    if [[ ! -e ${codeJson} ]]; then
        echo "${codeJson} code json file not exist"
        exit 1
    fi
    if [[ ! -e ${datx} ]]; then
        echo "${datx} datx file not exist"
        exit 1
    fi
    if [[ ! -e ${yaml} ]]; then
        echo "${yaml} yaml file not exist"
        exit 1
    fi
}

create_dir_remote() {
    host=$1
    path=$2
    if remote_exec ${host} "test -d ${path}"; then
    #if [ $(echo $?) -ne 0 ]; then
        remote_exec ${host} "mkdir -p ${path}"
        remote_exec ${host} "mkdir -p ${path}/pids/"
        remote_exec ${host} "mkdir -p ${path}/logs/"
        remote_exec ${host} "mkdir -p ${path}/assets/static/"
        remote_exec ${host} "mkdir -p ${path}/config/"
        remote_exec ${host} "mkdir -p ${path}/pm2/"
    fi
}

create_dir() {
    path=$1
    if [[ ! -d ${path} ]]; then
        mkdir -p ${path}
        mkdir -p ${path}/pids/
        mkdir -p ${path}/logs/
        mkdir -p ${path}/assets/static/
        mkdir -p ${path}/config/
        mkdir -p ${path}/pm2/
    fi
}

cp_src_file_remote() {
    host=$1
    path=$2
    jsonFile=$3
    shFile=$4
    appBinPath=$5
    binPath=$6
    switchJson=$7
    codeJson=$8
    datx=$9
    yaml=${10}
    remote_scp ${host} ${jsonFile} ${path}/pm2/
    remote_scp ${host} ${shFile} ${path}/pm2/
    remote_scp ${host} ${appBinPath} ${binPath}
    remote_scp ${host} ${switchJson} ${path}/assets/static/
    remote_scp ${host} ${codeJson} ${path}/assets/static/
    remote_scp ${host} ${datx} ${path}/config/
    remote_scp ${host} ${yaml} ${path}/config/
}

cp_src_file() {
    path=$1
    jsonFile=$2
    shFile=$3
    appBinPath=$4
    binPath=$5
    switchJson=$6
    codeJson=$7
    datx=$8
    yaml=$9
    cp -f ${jsonFile} ${path}/pm2/
    cp -f ${shFile} ${path}/pm2/
    cp -f ${appBinPath} ${binPath}
    cp -f ${switchJson} ${path}/assets/static/
    cp -f ${codeJson} ${path}/assets/static/
    cp -f ${datx} ${path}/config/
    cp -f ${yaml} ${path}/config/
}

create() {
    host=$1
    path=$2
    name=$3
    check ${host} ${path} ${name}
    echo "在主机${host}中创建新的服务名${name}路径${path}"
    jsonFile=${BINPATH}/pm2/${name}.json
    shFile=${BINPATH}/pm2/${name}.sh
    binFile=${name}-bin
    binPath=${path}/${binFile}
    n=${#name}
    appName=${name:0:n-1}
    appBinFile=${appName}-bin
    appBinPath=${BINPATH}/${appBinFile}
    switchJson=${BINPATH}/assets/static/switch.json
    codeJson=${BINPATH}/assets/static/code.json
    datx=${BINPATH}/config/17monipdb.datx
    yaml=${BINPATH}/config/config.${name}.yaml
    check_src_file ${jsonFile} ${shFile} ${appBinPath} ${switchJson} ${codeJson} ${datx} ${yaml}
    isRemote=`is_remote ${host}`
    if [[ ${isRemote} == "true" ]]; then
        status=`pm2show_remote ${host} ${name}`
        if [[ ${status} == 'online' ]]; then
            echo "${name} is already running"
            exit 1
        fi
        create_dir_remote ${host} ${path}
        cp_src_file_remote ${host} ${path} ${jsonFile} ${shFile} ${appBinPath} ${binPath} ${switchJson} ${codeJson} ${datx} ${yaml}
    else
        status=`pm2show ${name}`
        if [[ ${status} == 'online' ]]; then
            echo "${name} is already running"
            exit 1
        fi
        create_dir ${path}
        cp_src_file ${path} ${jsonFile} ${shFile} ${appBinPath} ${binPath} ${switchJson} ${codeJson} ${datx} ${yaml}
    fi
    echo "服务${name}创建成功"
}

check_nginx_conf() {
    nginxFile=$1
    if [[ ! -e ${nginxFile} ]]; then
        echo "${nginxFile} nginx conf not exist"
        exit 1
    fi
}

check_nginx_conf_remote() {
    host=$1
    nginxFile=$2
    remote_exec ${host} "test -e ${nginxFile}"
    if [ $(echo $?) -ne 0 ]; then
        echo "${nginxFile} remote nginx conf not exist"
        exit 1
    fi
}

check_nginx() {
    result=`service nginx status | awk '{ print $5}'`
    if [[ ${result} != "running..." ]]; then
        echo "nginx service stoped"
        exit 1
    fi
}

check_nginx_remote() {
    host=$1
    cmd="service nginx status"
    result=`remote_exec ${host} "${cmd}"`
    status=`echo "${result}" | awk '{ print $5}'`
    if [[ ${status} != "running..." ]]; then
        #echo "remote nginx service stoped"
        #exit 1
        check_nginx_remote7 ${host}
    fi
}

check_nginx_remote7() {
    host=$1
    cmd="systemctl status nginx.service"
    result=`remote_exec ${host} "${cmd}"`
    status=`echo "${result}" | grep "Active" | awk '{ print $3}'`
    if [[ ${status} != "(running)" ]]; then
        echo "remote nginx service stoped"
        exit 1
    fi
}

reload_nginx() {
    ${NGINXBIN} -t
    if [ $(echo $?) -ne 0 ]; then
        echo "nginx conf fail"
        exit 1
    fi
    ${NGINXBIN} -s reload
    if [ $(echo $?) -ne 0 ]; then
        echo "nginx reload fail"
        exit 1
    fi
}

reload_nginx_remote() {
    host=$1
    remote_exec ${host} "${NGINXBIN} -t"
    if [ $(echo $?) -ne 0 ]; then
        echo "remote nginx conf fail"
        exit 1
    fi
    remote_exec ${host} "${NGINXBIN} -s reload"
    if [ $(echo $?) -ne 0 ]; then
        echo "remote nginx reload fail"
        exit 1
    fi
}

nginx_conf_open() {
    port=$1
    appName=$2
    conf=$3
    addr="server 127.0.0.1:${port};"
    num=`sed -n "/^[[:space:]]*${addr}$/=" ${conf}`
    if [[ ${num} != '' ]]; then
        echo "${appName} ${port} is already open"
        exit 1
    fi
    num=`sed -n "/^[[:space:]]*#${addr}$/=" ${conf}`
    if [[ ${num} == '' ]]; then
        num=0
    fi
    if [ ${num} -eq 0 ]; then
        num=`sed -n "/^upstream ${appName}_upstream {$/=" ${conf}`
        if [ ${num} -eq 0 ]; then
            echo "${appName} upstream is not exist"
            exit 1
        fi
        #num=$[$num+1]
        sed -i "${num}a ${addr}" ${conf}
    elif [ ${num} -gt 0 ]; then
        sed -i "/^[[:space:]]*#${addr}$/s/#//g" ${conf}
    else
        echo "${appName} ${num} is error"
        exit 1
    fi
}

nginx_conf_open_remote() {
    host=$1
    port=$2
    appName=$3
    conf=$4
    addr="server 127.0.0.1:${port};"
    num=`remote_exec ${host} "sed -n '/^[[:space:]]*${addr}$/=' ${conf}"`
    if [[ ${num} != '' ]]; then
        echo "${appName} ${port} is already open"
        exit 1
    fi
    num=`remote_exec ${host} "sed -n '/^[[:space:]]*#${addr}$/=' ${conf}"`
    if [[ ${num} == '' ]]; then
        num=0
    fi
    if [ $num -eq 0 ]; then
        num=`remote_exec ${host} "sed -n '/^upstream ${appName}_upstream {$/=' ${conf}"`
        if [ $num -eq 0 ]; then
            echo "${appName} upstream is not exist"
            exit 1
        fi
        #num=$[$num+1]
        remote_exec ${host} "sed -i '${num}a ${addr}' ${conf}"
    elif [ ${num} -gt 0 ]; then
        remote_exec ${host} "sed -i '/^[[:space:]]*#${addr}$/s/#//g' ${conf}"
    else
        echo "${appName} ${num} is error"
        exit 1
    fi
}

open() {
    host=$1
    path=$2
    name=$3
    port=$4
    check ${host} ${path} ${name}
    echo "开放主机${host}上服务${name}访问"
    n=${#name}
    appName=${name:0:n-1}
    nginxName=${appName}.conf
    conf=${NGINXPATH}/conf/vhost/${nginxName}
    isRemote=`is_remote ${host}`
    if [[ ${isRemote} == "true" ]]; then
        check_nginx_conf_remote ${host} ${conf}
        check_nginx_remote ${host}
        result=`check_port_remote ${host} ${port}`
        if [ $result -eq 0 ]; then
            echo "${port} is not listen"
            exit 1
        fi
        nginx_conf_open_remote ${host} ${port} ${appName} ${conf}
        if [ $(echo $?) -ne 0 ]; then
            echo "${name} remote open fail"
            exit 1
        fi
        reload_nginx_remote ${host}
    else
        check_nginx_conf ${conf}
        check_nginx
        result=`check_port ${port}`
        if [ $result -eq 0 ]; then
            echo "${port} is not listen"
            exit 1
        fi
        nginx_conf_open ${port} ${appName} ${conf}
        if [ $(echo $?) -ne 0 ]; then
            echo "${name} open fail"
            exit 1
        fi
        reload_nginx
    fi
    echo "服务${name}开放访问成功"
}

nginx_conf_close() {
    port=$1
    appName=$2
    conf=$3
    addr="server 127.0.0.1:${port};"
    num=`grep -E -c "^[ ]*server 127.0.0.1:[0-9]{4};$" ${conf}`
    if [[ ${num} == '' ]]; then
        num=0
    fi
    if [ $num -le 1 ]; then
        echo "${appName} upstream close fail"
        exit 1
    fi
    num=`sed -n "/^[[:space:]]*#${addr}$/=" ${conf}`
    if [[ ${num} != '' ]]; then
        echo "${appName} ${port} is already closed"
        exit 1
    fi
    num=`sed -n "/^[[:space:]]*${addr}$/=" ${conf}`
    if [[ ${num} == '' ]]; then
        num=0
    fi
    if [ $num -eq 0 ]; then
        num=`sed -n "/^upstream ${appName}_upstream {$/=" ${conf}`
        if [ $num -eq 0 ]; then
            echo "${appName} upstream is not exist"
            exit 1
        fi
        #num=$[$num+1]
        sed -i "${num}a #${addr}" ${conf}
    elif [ ${num} -gt 0 ]; then
        sed -i "/^[[:space:]]*${addr}$/s/server/#server/g" ${conf}
    else
        echo "${appName} ${num} is error"
        exit 1
    fi
}

nginx_conf_close_remote() {
    host=$1
    port=$2
    appName=$3
    conf=$4
    addr="server 127.0.0.1:${port};"
    num=`remote_exec ${host} "grep -E -c '^[ ]*server 127.0.0.1:[0-9]{4};$' ${conf}"`
    if [[ ${num} == '' ]]; then
        num=0
    fi
    if [ $num -le 1 ]; then
        echo "${appName} upstream close fail"
        exit 1
    fi
    num=`remote_exec ${host} "sed -n '/^[[:space:]]*#${addr}$/=' ${conf}"`
    if [[ ${num} != '' ]]; then
        echo "${appName} ${port} is already closed"
        exit 1
    fi
    num=`remote_exec ${host} "sed -n '/^[[:space:]]*${addr}$/=' ${conf}"`
    if [[ ${num} == '' ]]; then
        num=0
    fi
    if [ $num -eq 0 ]; then
        num=`remote_exec ${host} "sed -n '/^upstream ${appName}_upstream {$/=' ${conf}"`
        if [ $num -eq 0 ]; then
            echo "${appName} upstream is not exist"
            exit 1
        fi
        #num=$[$num+1]
        remote_exec ${host} "sed -i '${num}a #${addr}' ${conf}"
    elif [ ${num} -gt 0 ]; then
        remote_exec ${host} "sed -i '/^[[:space:]]*${addr}$/s/server/#server/g' ${conf}"
    else
        echo "${appName} ${num} is error"
        exit 1
    fi
}

close() {
    host=$1
    path=$2
    name=$3
    port=$4
    check ${host} ${path} ${name}
    echo "关闭主机${host}上服务${name}访问"
    n=${#name}
    appName=${name:0:n-1}
    nginxName=${appName}.conf
    conf=${NGINXPATH}/conf/vhost/${nginxName}
    isRemote=`is_remote ${host}`
    if [[ ${isRemote} == "true" ]]; then
        check_nginx_conf_remote ${host} ${conf}
        check_nginx_remote ${host}
        nginx_conf_close_remote ${host} ${port} ${appName} ${conf}
        reload_nginx_remote ${host}
    else
        check_nginx_conf ${conf}
        check_nginx
        nginx_conf_close ${port} ${appName} ${conf}
        reload_nginx
    fi
    echo "服务${name}已经关闭访问"
}

case $1 in
    start)
        start $2 $3 $4 $5
        ;;
    stop)
        stop $2 $3 $4
        ;;
    open)
        open $2 $3 $4 $5
        ;;
    close)
        close $2 $3 $4 $5
        ;;
    create)
        create $2 $3 $4
        ;;
    update)
        update $2 $3 $4
        ;;
    *)
        usage
        ;;
esac
