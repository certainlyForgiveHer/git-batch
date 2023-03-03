### 下载
[gitbatch.exe](https://gitee.com/slagman001/git-batch/releases/tag/v1.0.0)

### 命令执行前检查清单
0. 确认仓库账户创建凭证, ssh -> public key, http/https -> api token
1. 确认设置仓库地址
1. 确认当前分支可以push且远程仓库有push权限
2. 确认设置仓库认证(ssh/token方式)
3. 推送代码

```
*********************************************** power by AIIS ********************************************************   
┌─┐       ┌─┐ + +                   
┌──┘ ┴───────┘ ┴──┐++                  
│                 │                    
│       ───       │++ + + +            
███████───███████ │+                   
│                 │+                   
│       ─┴─       │                    
│                 │                    
└───┐         ┌───┘                    
│         │                        
│         │   + +                  
│         │                        
│         └──────────────┐         
│                        │          
│                        ├─┐        
│                        ┌─┘        
│                        │          
└─┐  ┐  ┌───────┬──┐  ┌──┘  + + + +      
│ ─┤ ─┤       │ ─┤ ─┤               
└──┴──┘       └──┴──┘  + + + +         
神兽保佑                     
代码无BUG!   
*********************************************************************************************************************   
h  --> help   
swa--> set ciccwm api token   
sca--> set cicc api token   
sb --> show current branch   
scr--> set cicc repository   
swr--> set ciccwm repository   
sr --> show current repository   
pp --> show pre execute check list   
pu --> push cicc and ciccwm repo, execute git fetch first   
q  --> quit   
```