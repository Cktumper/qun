.PHONY:service

# 正式机地址
production_host = peon.top

service:
	@$(call build,$(production_host))
template:
	@$(call template,$(production_host))

# 构建函数
define build
	# 编译Go
	$(call compile)

	# 部署上传至服务器
	$(call deply,$(1))

	# 重启服务
	$(call reload,$(1))
endef

# 编译函数
define compile
	# 编译阶段
	GOPROXY="https://goproxy.io,direct" GOOS=linux GOARCH=amd64 go build -o ./bin/release_linux .
endef

# 部署函数
define deply
	scp -P 22 ./bin/release_linux root@$(1):/srv/services/qun/services_tmp
endef

#	部署模板
define template
	scp -P 22 -r ./bin/resources/* root@$(1):/srv/services/qun/public/
endef

# 激活服务
define reload
	ssh -p 22 root@$(1) "systemctl stop roser; mv /srv/services/qun/services_tmp /srv/services/qun/qun; systemctl restart qund"
	echo reload...OK
endef