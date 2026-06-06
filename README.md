Preview

https://github.com/user-attachments/assets/b27c7f92-6799-4ea2-a72f-9aae1465de48



Requirements:
- cloudflare installed on accessor's machine because I was poor to get proper tcp 
- add this to ~/.ssh/config:
```
Host safalski 
    HostName portfolio.safallama.com.np
    ProxyCommand cloudflared access ssh --hostname portfolio.safallama.com.np
    Port 1367
```
- run ssh safalski 
