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

* 删除分支

  ```shell
  git branch -D branchname
  ```

* 创建远程分支

  ```shell
  git checkout -b branchname origin/branchname
  ```

* 拉取远程

  ```shell
  git pull
  ```

  ```markdown
  合并到proto
  cd proto
  git pull
  git push
  git merge origin/master
  合并衝突
  vim 解決衝突
  git add 衝突文件
  git commit
  git push
  
  上傳之前先 pull
  git pull
  
  合并主分支代碼
  git merge origin/master
  
  error: Your local changes to the following files would be overwritten by merge:
          go.mod
          go.sum
  Please commit your changes or stash them before you merge.
  Aborting
  解決方法
  git stash
  
  在此合并
  git merge origin/master
  
  
  什么冲突之后 
  修改冲突文件然后合并
  git add
  git commit
  git push
  ```
  
  