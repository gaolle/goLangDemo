##### Xshell

* 切换

```shell
su - username
su - gaolle
```

### git

切换Git仓库目录（最顶层文件目录）

* 初始化仓库(默认进入主分支（master）)

  ```shell
  git init
  ```

* 查看仓库目录

  ```shell
  git status
  ```

* 添加文件到仓库(临时缓冲区)

  ```shell
  git add filename
  ```

* 将文件提交到仓库(commit(提交) -m(提交信息))

  ```shell
  git commit -m "add filename"
  ```

* 打印仓库提交日志

  ```shell
  git log
  ```

* 查看仓库分支情况(*所在行表示当前所在分支)

  ```shell
  git branch
  ```

* 增加分支

  ```shell
  git branch branchname
  ```

* 切换当前分支(当前尾部有提示所在分支)

  ```shell
  git checkout branchname
  ```

* 创建并切换到新分支

  ```shell
  git checkout -b branchname
  ```

* 合并分支（主分支与分支无冲突）

  ```shell
  git merge branchname
  ```

* 删除分支(主分支不能删除)

  ```shell
  git branch -d branchname
  ```

* SSH密钥和公钥生产（/c/Users/username/.ssh）id_rsa（密钥）id_rsa.pub（公钥） 公钥添加到Github

  ```shell
  ssh-keygen.exe
  ```

* 添加公钥之后测试连接

  ```shell
  ssh -T git@github.com
  ```

* 关联远程仓库

  ```shell
  git remote add origin githubaddress
  ```

* 同步原创仓库和本地仓库

  ```shell
  git pull origin master
  ```

* 提交至远程仓库

  ```shell
  git push origin master
  ```

* 强制拉取并覆盖本地代码

  ```shell
  git fetch --all
  git reset --hard origin/master
  git pull
  ```

* 放弃commit

  ```shell
  git reset --soft HEAD^
  ```

  